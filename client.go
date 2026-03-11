package okx

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	// DefaultBaseURL is the production base URL for OKX API
	DefaultBaseURL = "https://www.okx.com"
	// DemoBaseURL is the demo trading base URL
	DemoBaseURL = "https://www.okx.com"
)

// Client is the main OKX REST API client.
type Client struct {
	apiKey             string
	secretKey          string
	passphrase         string
	baseURL            string
	httpClient         *http.Client
	isDemo             bool
	logger             Logger
	rateLimiterEnabled bool
	mu                 sync.RWMutex
}

// NewClient creates a new OKX REST API client.
func NewClient(apiKey, secretKey, passphrase string, opts ...Option) *Client {
	c := &Client{
		apiKey:             apiKey,
		secretKey:          secretKey,
		passphrase:         passphrase,
		baseURL:            DefaultBaseURL,
		httpClient:         &http.Client{Timeout: 30 * time.Second},
		isDemo:             false,
		logger:             &noopLogger{},
		rateLimiterEnabled: false,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// sign generates the signature for OKX API authentication.
// Signature = Base64(HMAC-SHA256(timestamp + method + requestPath + body, secretKey))
func (c *Client) sign(timestamp, method, requestPath, body string) string {
	message := timestamp + method + requestPath + body
	h := hmac.New(sha256.New, []byte(c.secretKey))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// do executes an authenticated HTTP request and decodes the response.
func (c *Client) do(ctx context.Context, method, path string, params map[string]string, body interface{}, result interface{}) error {
	var bodyBytes []byte
	var err error

	if body != nil {
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	fullURL := c.baseURL + path
	if method == http.MethodGet && len(params) > 0 {
		queryParams := url.Values{}
		for k, v := range params {
			queryParams.Add(k, v)
		}
		fullURL += "?" + queryParams.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	requestPath := path
	if method == http.MethodGet && len(params) > 0 {
		queryParams := url.Values{}
		for k, v := range params {
			queryParams.Add(k, v)
		}
		requestPath += "?" + queryParams.Encode()
	}

	bodyStr := ""
	if len(bodyBytes) > 0 {
		bodyStr = string(bodyBytes)
	}

	signature := c.sign(timestamp, method, requestPath, bodyStr)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("OK-ACCESS-KEY", c.apiKey)
	req.Header.Set("OK-ACCESS-SIGN", signature)
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", c.passphrase)

	if c.isDemo {
		req.Header.Set("x-simulated-trading", "1")
	}

	c.logger.Debug("OKX API Request", "method", method, "path", requestPath)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	c.logger.Debug("OKX API Response", "status", resp.StatusCode, "body", string(respBody))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(respBody))
	}

	var envelope struct {
		Code string          `json:"code"`
		Msg  string          `json:"msg"`
		Data json.RawMessage `json:"data"`
	}

	if err := json.Unmarshal(respBody, &envelope); err != nil {
		return fmt.Errorf("failed to unmarshal response envelope: %w", err)
	}

	if envelope.Code != "0" {
		okxErr := &OKXError{
			Code:    envelope.Code,
			Message: envelope.Msg,
			Raw:     respBody,
		}
		if sentinelErr := MapErrorCode(envelope.Code); sentinelErr != nil {
			return fmt.Errorf("%w: %s", sentinelErr, okxErr.Error())
		}
		return okxErr
	}

	if result != nil {
		if err := json.Unmarshal(envelope.Data, result); err != nil {
			return fmt.Errorf("failed to unmarshal response data: %w", err)
		}
	}

	return nil
}

// doPublic executes an unauthenticated HTTP request (for public endpoints).
func (c *Client) doPublic(ctx context.Context, method, path string, params map[string]string, result interface{}) error {
	fullURL := c.baseURL + path
	if method == http.MethodGet && len(params) > 0 {
		queryParams := url.Values{}
		for k, v := range params {
			queryParams.Add(k, v)
		}
		fullURL += "?" + queryParams.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	c.logger.Debug("OKX Public API Request", "method", method, "path", path)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	c.logger.Debug("OKX Public API Response", "status", resp.StatusCode, "body", string(respBody))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(respBody))
	}

	var envelope struct {
		Code string          `json:"code"`
		Msg  string          `json:"msg"`
		Data json.RawMessage `json:"data"`
	}

	if err := json.Unmarshal(respBody, &envelope); err != nil {
		return fmt.Errorf("failed to unmarshal response envelope: %w", err)
	}

	if envelope.Code != "0" {
		okxErr := &OKXError{
			Code:    envelope.Code,
			Message: envelope.Msg,
			Raw:     respBody,
		}
		if sentinelErr := MapErrorCode(envelope.Code); sentinelErr != nil {
			return fmt.Errorf("%w: %s", sentinelErr, okxErr.Error())
		}
		return okxErr
	}

	if result != nil {
		if err := json.Unmarshal(envelope.Data, result); err != nil {
			return fmt.Errorf("failed to unmarshal response data: %w", err)
		}
	}

	return nil
}

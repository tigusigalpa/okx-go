package okx

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSign(t *testing.T) {
	client := NewClient("test-api-key", "test-secret-key", "test-passphrase")

	timestamp := "2020-12-08T09:08:57.715Z"
	method := "GET"
	requestPath := "/api/v5/account/balance"
	body := ""

	signature := client.sign(timestamp, method, requestPath, body)

	assert.NotEmpty(t, signature)
	assert.True(t, len(signature) > 0)
}

func TestNewClient(t *testing.T) {
	client := NewClient("api-key", "secret-key", "passphrase")

	assert.NotNil(t, client)
	assert.Equal(t, "api-key", client.apiKey)
	assert.Equal(t, "secret-key", client.secretKey)
	assert.Equal(t, "passphrase", client.passphrase)
	assert.Equal(t, DefaultBaseURL, client.baseURL)
	assert.False(t, client.isDemo)
}

func TestNewClientWithOptions(t *testing.T) {
	client := NewClient(
		"api-key",
		"secret-key",
		"passphrase",
		WithDemoTrading(),
		WithTimeout(60*time.Second),
		WithBaseURL("https://custom.url"),
	)

	assert.NotNil(t, client)
	assert.True(t, client.isDemo)
	assert.Equal(t, "https://custom.url", client.baseURL)
	assert.Equal(t, 60*time.Second, client.httpClient.Timeout)
}

func TestNewClientAppliesAllOptions(t *testing.T) {
	httpClient := &http.Client{}
	logger := &noopLogger{}
	client := NewClient("", "", "", WithHTTPClient(httpClient), WithRateLimiter(true), WithLogger(logger))
	require.Same(t, httpClient, client.httpClient)
	require.True(t, client.rateLimiterEnabled)
	require.Same(t, logger, client.logger)
}

func TestLoggersAndRestClient(t *testing.T) {
	noop := &noopLogger{}
	noop.Debug("debug")
	noop.Info("info")
	noop.Warn("warn")
	noop.Error("error")

	logger := NewSlogLogger(slog.New(slog.NewTextHandler(io.Discard, nil)))
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")

	restClient := NewRestClient("api-key", "secret", "passphrase")
	require.NotNil(t, restClient.Client)
	require.NotNil(t, restClient.Account)
	require.NotNil(t, restClient.Trade)
	require.NotNil(t, restClient.Market)
	require.NotNil(t, restClient.Public)
	require.NotNil(t, restClient.Asset)
	require.NotNil(t, restClient.System)
	require.NotNil(t, restClient.Support)
	require.NotNil(t, restClient.Users)
}

func TestMapErrorCode(t *testing.T) {
	tests := []struct {
		code     string
		expected error
	}{
		{"50100", ErrUnauthorized},
		{"50101", ErrUnauthorized},
		{"50011", ErrRateLimited},
		{"50000", ErrBadRequest},
		{"50014", ErrNotFound},
		{"50012", ErrInvalidParameter},
		{"50003", ErrServiceUnavail},
		{"1", ErrInternalServer},
		{"99999", nil},
	}

	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			err := MapErrorCode(tt.code)
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestError(t *testing.T) {
	err := &Error{
		Code:    "50000",
		Message: "Body cannot be empty",
		Raw:     []byte(`{"code":"50000","msg":"Body cannot be empty"}`),
	}

	assert.Contains(t, err.Error(), "50000")
	assert.Contains(t, err.Error(), "Body cannot be empty")
}

// TestDo_PreservesDetailedError verifies that when the OKX API returns a
// non-zero envelope code that maps to a sentinel error (e.g. code="1" ->
// ErrInternalServer), the underlying *Error (including the raw body with
// data[].sCode / data[].sMsg) is preserved in the error chain and can be
// retrieved via errors.As, while errors.Is still matches the sentinel.
func TestDo_PreservesDetailedError(t *testing.T) {
	const rawBody = `{"code":"1","msg":"All operations failed","data":[{"sCode":"51008","sMsg":"Order failed. Insufficient balance"}]}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(rawBody))
	}))
	defer server.Close()

	client := NewClient("api-key", "secret-key", "passphrase", WithBaseURL(server.URL))

	err := client.do(context.Background(), http.MethodGet, "/api/v5/trade/order", nil, nil, nil)
	require.Error(t, err)

	assert.True(t, errors.Is(err, ErrInternalServer))

	var apiErr *Error
	require.True(t, errors.As(err, &apiErr))
	assert.Equal(t, "1", apiErr.Code)
	assert.Equal(t, "All operations failed", apiErr.Message)
	assert.Contains(t, string(apiErr.Raw), "51008")
	assert.Contains(t, string(apiErr.Raw), "Order failed. Insufficient balance")
}

func TestDoBuildsAuthenticatedRequestAndDecodesResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/api/v5/trade/order", r.URL.Path)
		require.Equal(t, "api-key", r.Header.Get("OK-ACCESS-KEY"))
		require.Equal(t, "passphrase", r.Header.Get("OK-ACCESS-PASSPHRASE"))
		require.NotEmpty(t, r.Header.Get("OK-ACCESS-SIGN"))
		require.NotEmpty(t, r.Header.Get("OK-ACCESS-TIMESTAMP"))
		require.Equal(t, "1", r.Header.Get("x-simulated-trading"))

		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.JSONEq(t, `{"instId":"BTC-USDT"}`, string(body))
		_, _ = w.Write([]byte(`{"code":"0","msg":"","data":[{"ordId":"42"}]}`))
	}))
	defer server.Close()

	client := NewClient("api-key", "secret-key", "passphrase", WithBaseURL(server.URL), WithDemoTrading())
	var result []struct {
		OrdID string `json:"ordId"`
	}
	err := client.do(context.Background(), http.MethodPost, "/api/v5/trade/order", nil, map[string]string{"instId": "BTC-USDT"}, &result)
	require.NoError(t, err)
	require.Equal(t, "42", result[0].OrdID)
}

func TestDoPublicBuildsQueryAndDecodesResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v5/market/ticker", r.URL.Path)
		require.Equal(t, "BTC-USDT", r.URL.Query().Get("instId"))
		require.Empty(t, r.Header.Get("OK-ACCESS-KEY"))
		_, _ = w.Write([]byte(`{"code":"0","msg":"","data":[{"instId":"BTC-USDT"}]}`))
	}))
	defer server.Close()

	client := NewClient("", "", "", WithBaseURL(server.URL))
	var result []map[string]string
	err := client.doPublic(context.Background(), http.MethodGet, "/api/v5/market/ticker", map[string]string{"instId": "BTC-USDT"}, &result)
	require.NoError(t, err)
	require.Equal(t, "BTC-USDT", result[0]["instId"])
}

func TestDoReturnsResponseAndDecodeErrors(t *testing.T) {
	tests := []struct {
		name     string
		status   int
		response string
		errText  string
	}{
		{"unexpected status", http.StatusBadGateway, "upstream unavailable", "unexpected status code: 502"},
		{"invalid envelope", http.StatusOK, "not json", "failed to unmarshal response envelope"},
		{"invalid data", http.StatusOK, `{"code":"0","msg":"","data":{"value":true}}`, "failed to unmarshal response data"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.response))
			}))
			defer server.Close()

			client := NewClient("", "", "", WithBaseURL(server.URL))
			var result []string
			err := client.do(context.Background(), http.MethodGet, "/test", nil, nil, &result)
			require.ErrorContains(t, err, tt.errText)
		})
	}
}

func TestDoRejectsUnmarshalableBody(t *testing.T) {
	client := NewClient("", "", "")
	err := client.do(context.Background(), http.MethodPost, "/test", nil, make(chan int), nil)
	require.ErrorContains(t, err, "failed to marshal request body")
}

func TestDoPublicReturnsAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"code":"50011","msg":"too many requests","data":[]}`))
	}))
	defer server.Close()

	client := NewClient("", "", "", WithBaseURL(server.URL))
	err := client.doPublic(context.Background(), http.MethodGet, "/test", nil, nil)
	require.ErrorIs(t, err, ErrRateLimited)
	var apiErr *Error
	require.ErrorAs(t, err, &apiErr)
	require.Equal(t, "50011", apiErr.Code)
}

func TestDoUsesStableJSONBodyForSignature(t *testing.T) {
	client := NewClient("", "secret", "")
	body, err := json.Marshal(map[string]string{"instId": "BTC-USDT"})
	require.NoError(t, err)
	require.NotEmpty(t, client.sign("2020-01-01T00:00:00.000Z", http.MethodPost, "/test", string(body)))
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestDoReturnsTransportError(t *testing.T) {
	client := NewClient("", "", "", WithHTTPClient(&http.Client{Transport: roundTripperFunc(func(*http.Request) (*http.Response, error) {
		return nil, errors.New("network unavailable")
	})}))

	err := client.do(context.Background(), http.MethodGet, "/test", nil, nil, nil)
	require.ErrorContains(t, err, "failed to execute request")
}

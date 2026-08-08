package okx

import (
	"context"
	"errors"
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

func TestOKXError(t *testing.T) {
	err := &OKXError{
		Code:    "50000",
		Message: "Body cannot be empty",
		Raw:     []byte(`{"code":"50000","msg":"Body cannot be empty"}`),
	}

	assert.Contains(t, err.Error(), "50000")
	assert.Contains(t, err.Error(), "Body cannot be empty")
}

// TestDo_PreservesDetailedOKXError verifies that when the OKX API returns a
// non-zero envelope code that maps to a sentinel error (e.g. code="1" ->
// ErrInternalServer), the underlying *OKXError (including the raw body with
// data[].sCode / data[].sMsg) is preserved in the error chain and can be
// retrieved via errors.As, while errors.Is still matches the sentinel.
func TestDo_PreservesDetailedOKXError(t *testing.T) {
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

	var apiErr *OKXError
	require.True(t, errors.As(err, &apiErr))
	assert.Equal(t, "1", apiErr.Code)
	assert.Equal(t, "All operations failed", apiErr.Message)
	assert.Contains(t, string(apiErr.Raw), "51008")
	assert.Contains(t, string(apiErr.Raw), "Order failed. Insufficient balance")
}

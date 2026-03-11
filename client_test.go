package okx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

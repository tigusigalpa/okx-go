package okx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWSClient(t *testing.T) {
	ws := NewWSClient("api-key", "secret-key", "passphrase", WSPublicURL)

	assert.NotNil(t, ws)
	assert.Equal(t, "api-key", ws.apiKey)
	assert.Equal(t, "secret-key", ws.secretKey)
	assert.Equal(t, "passphrase", ws.passphrase)
	assert.Equal(t, WSPublicURL, ws.url)
	assert.False(t, ws.isDemo)
	assert.NotNil(t, ws.subscriptions)
}

func TestNewWSClientWithOptions(t *testing.T) {
	ws := NewWSClient(
		"api-key",
		"secret-key",
		"passphrase",
		WSPublicURL,
		WithWSDemo(),
	)

	assert.NotNil(t, ws)
	assert.True(t, ws.isDemo)
}

func TestMakeSubKey(t *testing.T) {
	ws := NewWSClient("", "", "", WSPublicURL)

	tests := []struct {
		name     string
		channel  string
		args     map[string]interface{}
		expected string
	}{
		{
			name:     "channel only",
			channel:  "tickers",
			args:     map[string]interface{}{},
			expected: "tickers",
		},
		{
			name:    "channel with instId",
			channel: "tickers",
			args: map[string]interface{}{
				"instId": "BTC-USDT",
			},
			expected: "tickers:BTC-USDT",
		},
		{
			name:    "channel with instType",
			channel: "tickers",
			args: map[string]interface{}{
				"instType": "SPOT",
			},
			expected: "tickers:SPOT",
		},
		{
			name:    "channel with ccy",
			channel: "account",
			args: map[string]interface{}{
				"ccy": "BTC",
			},
			expected: "account:BTC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := ws.makeSubKey(tt.channel, tt.args)
			assert.Equal(t, tt.expected, key)
		})
	}
}

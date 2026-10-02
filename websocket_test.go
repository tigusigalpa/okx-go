package okx

import (
	"context"
	"net/url"
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
	assert.Equal(t, WSDemoPublicURL, ws.url)
}

func TestWebSocketURL(t *testing.T) {
	tests := []struct {
		name    string
		region  WSRegion
		service WSService
		demo    bool
		want    string
	}{
		{"global public", WSRegionGlobal, WSServicePublic, false, WSPublicURL},
		{"global private", WSRegionGlobal, WSServicePrivate, false, WSPrivateURL},
		{"global business", WSRegionGlobal, WSServiceBusiness, false, WSBusinessURL},
		{"global demo public", WSRegionGlobal, WSServicePublic, true, WSDemoPublicURL},
		{"global demo private", WSRegionGlobal, WSServicePrivate, true, WSDemoPrivateURL},
		{"global demo business", WSRegionGlobal, WSServiceBusiness, true, WSDemoBusinessURL},
		{"eea public", WSRegionEEA, WSServicePublic, false, WSEEAPublicURL},
		{"eea private", WSRegionEEA, WSServicePrivate, false, WSEEAPrivateURL},
		{"eea business", WSRegionEEA, WSServiceBusiness, false, WSEEABusinessURL},
		{"eea demo public", WSRegionEEA, WSServicePublic, true, WSDemoEEAPublicURL},
		{"eea demo private", WSRegionEEA, WSServicePrivate, true, WSDemoEEAPrivateURL},
		{"eea demo business", WSRegionEEA, WSServiceBusiness, true, WSDemoEEABusinessURL},
		{"us public", WSRegionUS, WSServicePublic, false, WSUSPublicURL},
		{"us private", WSRegionUS, WSServicePrivate, false, WSUSPrivateURL},
		{"us business", WSRegionUS, WSServiceBusiness, false, WSUSBusinessURL},
		{"us demo public", WSRegionUS, WSServicePublic, true, WSDemoUSPublicURL},
		{"us demo private", WSRegionUS, WSServicePrivate, true, WSDemoUSPrivateURL},
		{"us demo business", WSRegionUS, WSServiceBusiness, true, WSDemoUSBusinessURL},
		{"turkiye uses global", WSRegionTR, WSServicePublic, false, WSPublicURL},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WebSocketURL(tt.region, tt.service, tt.demo)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
			parsed, err := url.Parse(got)
			assert.NoError(t, err)
			assert.Equal(t, "wss", parsed.Scheme)
			assert.Empty(t, parsed.Port())
			assert.Equal(t, "/ws/v5/"+string(tt.service), parsed.Path)
		})
	}

	_, err := WebSocketURL("moon", WSServicePublic, false)
	assert.Error(t, err)
	_, err = WebSocketURL(WSRegionGlobal, "unknown", false)
	assert.Error(t, err)
}

func TestWSClientCloseIsIdempotent(t *testing.T) {
	ws := NewWSClient("", "", "", WSPublicURL)

	assert.NoError(t, ws.Close())
	assert.NoError(t, ws.Close())
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
			name:    "channel with additional argument",
			channel: "funding-rate",
			args: map[string]interface{}{
				"instId": "BTC-USD-SWAP",
				"uly":    "BTC-USD",
			},
			expected: "funding-rate:BTC-USD-SWAP:{\"uly\":\"BTC-USD\"}",
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

func TestWSClientOptions(t *testing.T) {
	client := NewWSClient("", "", "", WSPublicURL, WithWSSubscriptionBuffer(7))
	assert.Equal(t, 7, client.subscriptionBuffer)
	assert.NoError(t, client.Close())

	invalid := NewWSClient("", "", "", WSPublicURL, WithWSSubscriptionBuffer(0))
	_, err := invalid.Subscribe(context.Background(), "tickers", map[string]interface{}{"instId": "BTC-USDT"})
	assert.ErrorIs(t, err, ErrWSInvalidSubscriptionBuffer)
	assert.NoError(t, invalid.Close())
}

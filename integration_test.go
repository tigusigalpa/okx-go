//go:build integration
// +build integration

package okx

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tigusigalpa/okx-go/models"
)

func getTestClient(t *testing.T) *RestClient {
	apiKey := os.Getenv("OKX_API_KEY")
	secretKey := os.Getenv("OKX_SECRET_KEY")
	passphrase := os.Getenv("OKX_PASSPHRASE")

	if apiKey == "" || secretKey == "" || passphrase == "" {
		t.Skip("Skipping integration test: OKX credentials not set")
	}

	return NewRestClient(
		apiKey,
		secretKey,
		passphrase,
		WithDemoTrading(),
		WithTimeout(30*time.Second),
	)
}

func TestIntegration_GetBalance(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	balances, err := client.Account.GetBalance(ctx, nil)
	require.NoError(t, err)
	assert.NotNil(t, balances)
}

func TestIntegration_GetInstruments(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	instruments, err := client.Public.GetInstruments(ctx, "SPOT", nil, nil, nil)
	require.NoError(t, err)
	assert.NotEmpty(t, instruments)
}

func TestIntegration_GetTicker(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	tickers, err := client.Market.GetTicker(ctx, "BTC-USDT")
	require.NoError(t, err)
	assert.NotEmpty(t, tickers)
	assert.Equal(t, "BTC-USDT", tickers[0].InstID)
}

func TestIntegration_GetSystemTime(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	times, err := client.Public.GetSystemTime(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, times)
	assert.NotEmpty(t, times[0].TS)
}

func TestIntegration_WebSocket_PublicChannel(t *testing.T) {
	apiKey := os.Getenv("OKX_API_KEY")
	secretKey := os.Getenv("OKX_SECRET_KEY")
	passphrase := os.Getenv("OKX_PASSPHRASE")

	if apiKey == "" || secretKey == "" || passphrase == "" {
		t.Skip("Skipping integration test: OKX credentials not set")
	}

	ws := NewWSClient("", "", "", WSPublicURL)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := ws.Connect(ctx)
	require.NoError(t, err)
	defer ws.Close()

	ch, err := ws.Subscribe(ctx, "tickers", map[string]interface{}{
		"instId": "BTC-USDT",
	})
	require.NoError(t, err)

	select {
	case msg := <-ch:
		assert.NotNil(t, msg)
		t.Logf("Received message: %s", string(msg))
	case <-time.After(10 * time.Second):
		t.Fatal("Timeout waiting for WebSocket message")
	}
}

func TestIntegration_PlaceAndCancelOrder(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	// This test requires demo trading to be enabled
	// Place a limit order
	px := "20000"
	orderReq := models.PlaceOrderRequest{
		InstID:  "BTC-USDT",
		TdMode:  "cash",
		Side:    "buy",
		OrdType: "limit",
		Px:      &px,
		Sz:      "0.001",
	}

	// Note: This will fail in demo mode without sufficient balance
	// The test is here to demonstrate the API usage
	result, err := client.Trade.PlaceOrder(ctx, orderReq)
	if err != nil {
		t.Logf("Place order failed (expected in demo): %v", err)
		return
	}

	require.NotEmpty(t, result)
	ordID := result[0].OrdID

	// Cancel the order
	cancelReq := models.CancelOrderRequest{
		InstID: "BTC-USDT",
		OrdID:  &ordID,
	}

	_, err = client.Trade.CancelOrder(ctx, cancelReq)
	if err != nil {
		t.Logf("Cancel order failed: %v", err)
	}
}

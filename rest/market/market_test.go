package market

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTickerBuildsRequest(t *testing.T) {
	client := NewClient(func(_ context.Context, method, path string, params map[string]string, result interface{}) error {
		require.Equal(t, http.MethodGet, method)
		require.Equal(t, "/api/v5/market/ticker", path)
		require.Equal(t, map[string]string{"instId": "BTC-USDT"}, params)
		return nil
	})

	_, err := client.GetTicker(context.Background(), "BTC-USDT")
	require.NoError(t, err)
}

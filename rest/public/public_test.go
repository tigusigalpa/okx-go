package public

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetInstrumentsBuildsRequest(t *testing.T) {
	instID := "BTC-USDT"
	client := NewClient(func(_ context.Context, method, path string, params map[string]string, result interface{}) error {
		require.Equal(t, http.MethodGet, method)
		require.Equal(t, "/api/v5/public/instruments", path)
		require.Equal(t, map[string]string{"instType": "SPOT", "instId": "BTC-USDT"}, params)
		return nil
	})

	_, err := client.GetInstruments(context.Background(), "SPOT", nil, nil, &instID)
	require.NoError(t, err)
}

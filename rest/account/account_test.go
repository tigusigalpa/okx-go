package account

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBalanceBuildsRequest(t *testing.T) {
	ccy := "BTC"
	client := NewClient(func(_ context.Context, method, path string, params map[string]string, body interface{}, result interface{}) error {
		require.Equal(t, http.MethodGet, method)
		require.Equal(t, "/api/v5/account/balance", path)
		require.Equal(t, map[string]string{"ccy": "BTC"}, params)
		require.Nil(t, body)
		return nil
	})

	_, err := client.GetBalance(context.Background(), &ccy)
	require.NoError(t, err)
}

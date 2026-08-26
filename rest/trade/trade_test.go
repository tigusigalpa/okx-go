package trade

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tigusigalpa/okx-go/models"
)

func TestPlaceOrderBuildsRequest(t *testing.T) {
	req := models.PlaceOrderRequest{InstID: "BTC-USDT", TdMode: "cash", Side: "buy", OrdType: "market", Sz: "0.001"}
	client := NewClient(func(_ context.Context, method, path string, params map[string]string, body interface{}, result interface{}) error {
		require.Equal(t, http.MethodPost, method)
		require.Equal(t, "/api/v5/trade/order", path)
		require.Nil(t, params)
		require.Equal(t, req, body)
		return nil
	})

	_, err := client.PlaceOrder(context.Background(), req)
	require.NoError(t, err)
}

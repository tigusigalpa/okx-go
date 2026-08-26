package asset

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tigusigalpa/okx-go/models"
)

func TestTransferBuildsRequest(t *testing.T) {
	req := models.TransferRequest{Ccy: "USDT", Amt: "1", From: "6", To: "18"}
	client := NewClient(func(_ context.Context, method, path string, params map[string]string, body interface{}, result interface{}) error {
		require.Equal(t, http.MethodPost, method)
		require.Equal(t, "/api/v5/asset/transfer", path)
		require.Nil(t, params)
		require.Equal(t, req, body)
		return nil
	})

	_, err := client.Transfer(context.Background(), req)
	require.NoError(t, err)
}

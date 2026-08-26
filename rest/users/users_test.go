package users

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tigusigalpa/okx-go/models"
)

func TestCreateSubAccountBuildsRequest(t *testing.T) {
	req := models.CreateSubAccountRequest{SubAcct: "strategy-a"}
	client := NewClient(func(_ context.Context, method, path string, params map[string]string, body interface{}, result interface{}) error {
		require.Equal(t, http.MethodPost, method)
		require.Equal(t, "/api/v5/users/subaccount/create-subaccount", path)
		require.Nil(t, params)
		require.Equal(t, req, body)
		return nil
	})

	_, err := client.CreateSubAccount(context.Background(), req)
	require.NoError(t, err)
}

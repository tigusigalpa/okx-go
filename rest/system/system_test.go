package system

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStatusBuildsRequest(t *testing.T) {
	state := "scheduled"
	client := NewClient(func(_ context.Context, method, path string, params map[string]string, result interface{}) error {
		require.Equal(t, http.MethodGet, method)
		require.Equal(t, "/api/v5/system/status", path)
		require.Equal(t, map[string]string{"state": "scheduled"}, params)
		return nil
	})

	_, err := client.GetStatus(context.Background(), &state)
	require.NoError(t, err)
}

package support

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAnnouncementTypesBuildsRequest(t *testing.T) {
	client := NewClient(func(_ context.Context, method, path string, params map[string]string, result interface{}) error {
		require.Equal(t, http.MethodGet, method)
		require.Equal(t, "/api/v5/support/announcement-types", path)
		require.Nil(t, params)
		return nil
	})

	_, err := client.GetAnnouncementTypes(context.Background())
	require.NoError(t, err)
}

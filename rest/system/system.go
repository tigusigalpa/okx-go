package system

import (
	"context"
	"net/http"

	"github.com/tigusigalpa/okx-go/models"
)

type Client struct {
	doPublicFunc func(ctx context.Context, method, path string, params map[string]string, result interface{}) error
}

func NewClient(doPublicFunc func(ctx context.Context, method, path string, params map[string]string, result interface{}) error) *Client {
	return &Client{doPublicFunc: doPublicFunc}
}

func (c *Client) GetStatus(ctx context.Context, state *string) ([]models.SystemStatus, error) {
	params := make(map[string]string)
	if state != nil {
		params["state"] = *state
	}

	var result []models.SystemStatus
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/system/status", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

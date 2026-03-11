package support

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

func (c *Client) GetAnnouncementTypes(ctx context.Context) ([]models.AnnouncementType, error) {
	var result []models.AnnouncementType
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/support/announcement-types", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetAnnouncements(ctx context.Context, annType *string, page *string, limit *string) ([]models.Announcement, error) {
	params := make(map[string]string)
	if annType != nil {
		params["annType"] = *annType
	}
	if page != nil {
		params["page"] = *page
	}
	if limit != nil {
		params["limit"] = *limit
	}

	var result []models.Announcement
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/support/announcements", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

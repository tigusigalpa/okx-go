package users

import (
	"context"
	"net/http"

	"github.com/tigusigalpa/okx-go/models"
)

type Client struct {
	doFunc func(ctx context.Context, method, path string, params map[string]string, body interface{}, result interface{}) error
}

func NewClient(doFunc func(ctx context.Context, method, path string, params map[string]string, body interface{}, result interface{}) error) *Client {
	return &Client{doFunc: doFunc}
}

func (c *Client) GetSubAccountList(ctx context.Context, enable *bool, subAcct *string, after *string, before *string, limit *string) ([]models.SubAccount, error) {
	params := make(map[string]string)
	if enable != nil {
		if *enable {
			params["enable"] = "true"
		} else {
			params["enable"] = "false"
		}
	}
	if subAcct != nil {
		params["subAcct"] = *subAcct
	}
	if after != nil {
		params["after"] = *after
	}
	if before != nil {
		params["before"] = *before
	}
	if limit != nil {
		params["limit"] = *limit
	}

	var result []models.SubAccount
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/users/subaccount/list", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) CreateSubAccount(ctx context.Context, req models.CreateSubAccountRequest) ([]models.CreateSubAccountResponse, error) {
	var result []models.CreateSubAccountResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/users/subaccount/create-subaccount", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetSubAccountAPIKey(ctx context.Context, subAcct string, apiKey *string) ([]models.SubAccountAPIKey, error) {
	params := map[string]string{
		"subAcct": subAcct,
	}
	if apiKey != nil {
		params["apiKey"] = *apiKey
	}

	var result []models.SubAccountAPIKey
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/users/subaccount/apikey", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) CreateSubAccountAPIKey(ctx context.Context, req models.CreateSubAccountAPIKeyRequest) ([]models.CreateSubAccountAPIKeyResponse, error) {
	var result []models.CreateSubAccountAPIKeyResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/users/subaccount/apikey", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) ModifySubAccountAPIKey(ctx context.Context, req models.ModifySubAccountAPIKeyRequest) ([]models.SubAccountAPIKey, error) {
	var result []models.SubAccountAPIKey
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/users/subaccount/modify-apikey", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) DeleteSubAccountAPIKey(ctx context.Context, req models.DeleteSubAccountAPIKeyRequest) error {
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/users/subaccount/delete-apikey", nil, req, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) SetSubAccountTransferOut(ctx context.Context, req models.SetSubAccountTransferOutRequest) error {
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/users/subaccount/set-transfer-out", nil, req, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetEntrustSubAccountList(ctx context.Context, subAcct *string) ([]models.EntrustSubAccountList, error) {
	params := make(map[string]string)
	if subAcct != nil {
		params["subAcct"] = *subAcct
	}

	var result []models.EntrustSubAccountList
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/users/entrust-subaccount-list", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

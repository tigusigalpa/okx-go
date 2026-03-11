package asset

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

func (c *Client) GetCurrencies(ctx context.Context, ccy *string) ([]models.Currency, error) {
	params := make(map[string]string)
	if ccy != nil {
		params["ccy"] = *ccy
	}

	var result []models.Currency
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/currencies", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetBalances(ctx context.Context, ccy *string) ([]models.AssetBalance, error) {
	params := make(map[string]string)
	if ccy != nil {
		params["ccy"] = *ccy
	}

	var result []models.AssetBalance
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/balances", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetAssetValuation(ctx context.Context, ccy *string) ([]models.AssetValuation, error) {
	params := make(map[string]string)
	if ccy != nil {
		params["ccy"] = *ccy
	}

	var result []models.AssetValuation
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/asset-valuation", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) Transfer(ctx context.Context, req models.TransferRequest) ([]models.TransferResponse, error) {
	var result []models.TransferResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/asset/transfer", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetTransferState(ctx context.Context, transID *string, clientID *string, type_ *string) ([]models.TransferState, error) {
	params := make(map[string]string)
	if transID != nil {
		params["transId"] = *transID
	}
	if clientID != nil {
		params["clientId"] = *clientID
	}
	if type_ != nil {
		params["type"] = *type_
	}

	var result []models.TransferState
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/transfer-state", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) Withdrawal(ctx context.Context, req models.WithdrawalRequest) ([]models.WithdrawalResponse, error) {
	var result []models.WithdrawalResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/asset/withdrawal", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) CancelWithdrawal(ctx context.Context, wdID string) error {
	body := map[string]string{
		"wdId": wdID,
	}
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/asset/cancel-withdrawal", nil, body, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetWithdrawalHistory(ctx context.Context, ccy *string, wdID *string, clientID *string, txID *string, type_ *string, state *string, after *string, before *string, limit *string) ([]models.WithdrawalHistory, error) {
	params := make(map[string]string)
	if ccy != nil {
		params["ccy"] = *ccy
	}
	if wdID != nil {
		params["wdId"] = *wdID
	}
	if clientID != nil {
		params["clientId"] = *clientID
	}
	if txID != nil {
		params["txId"] = *txID
	}
	if type_ != nil {
		params["type"] = *type_
	}
	if state != nil {
		params["state"] = *state
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

	var result []models.WithdrawalHistory
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/withdrawal-history", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetDepositAddress(ctx context.Context, ccy string) ([]models.DepositAddress, error) {
	params := map[string]string{
		"ccy": ccy,
	}

	var result []models.DepositAddress
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/deposit-address", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetDepositHistory(ctx context.Context, ccy *string, depID *string, fromWdID *string, txID *string, type_ *string, state *string, after *string, before *string, limit *string) ([]models.DepositHistory, error) {
	params := make(map[string]string)
	if ccy != nil {
		params["ccy"] = *ccy
	}
	if depID != nil {
		params["depId"] = *depID
	}
	if fromWdID != nil {
		params["fromWdId"] = *fromWdID
	}
	if txID != nil {
		params["txId"] = *txID
	}
	if type_ != nil {
		params["type"] = *type_
	}
	if state != nil {
		params["state"] = *state
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

	var result []models.DepositHistory
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/deposit-history", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetBills(ctx context.Context, ccy *string, type_ *string, clientID *string, after *string, before *string, limit *string) ([]models.AssetBill, error) {
	params := make(map[string]string)
	if ccy != nil {
		params["ccy"] = *ccy
	}
	if type_ != nil {
		params["type"] = *type_
	}
	if clientID != nil {
		params["clientId"] = *clientID
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

	var result []models.AssetBill
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/bills", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetSavingBalance(ctx context.Context, ccy *string) ([]models.SavingBalance, error) {
	params := make(map[string]string)
	if ccy != nil {
		params["ccy"] = *ccy
	}

	var result []models.SavingBalance
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/saving-balance", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) PurchaseRedempt(ctx context.Context, req models.PurchaseRedemptRequest) ([]models.PurchaseRedemptResponse, error) {
	var result []models.PurchaseRedemptResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/asset/purchase-redempt", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) SetLendingRate(ctx context.Context, ccy string, rate string) ([]models.LendingRate, error) {
	body := map[string]string{
		"ccy":  ccy,
		"rate": rate,
	}

	var result []models.LendingRate
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/asset/set-lending-rate", nil, body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetLendingRateSummary(ctx context.Context, ccy *string) ([]models.LendingRate, error) {
	params := make(map[string]string)
	if ccy != nil {
		params["ccy"] = *ccy
	}

	var result []models.LendingRate
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/lending-rate-summary", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetLendingRateHistory(ctx context.Context, ccy *string, after *string, before *string, limit *string) ([]models.LendingRate, error) {
	params := make(map[string]string)
	if ccy != nil {
		params["ccy"] = *ccy
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

	var result []models.LendingRate
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/lending-rate-history", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetConvertCurrencies(ctx context.Context) ([]models.ConvertCurrencyPair, error) {
	var result []models.ConvertCurrencyPair
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/convert/currencies", nil, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetConvertCurrencyPair(ctx context.Context, fromCcy string, toCcy string) ([]models.ConvertCurrencyPair, error) {
	params := map[string]string{
		"fromCcy": fromCcy,
		"toCcy":   toCcy,
	}

	var result []models.ConvertCurrencyPair
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/convert/currency-pair", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) EstimateConvertQuote(ctx context.Context, req models.ConvertEstimateQuoteRequest) ([]models.ConvertEstimateQuoteResponse, error) {
	var result []models.ConvertEstimateQuoteResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/asset/convert/estimate-quote", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) ConvertTrade(ctx context.Context, req models.ConvertTradeRequest) ([]models.ConvertTradeResponse, error) {
	var result []models.ConvertTradeResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/asset/convert/trade", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetConvertHistory(ctx context.Context, after *string, before *string, limit *string, tag *string) ([]models.ConvertTradeResponse, error) {
	params := make(map[string]string)
	if after != nil {
		params["after"] = *after
	}
	if before != nil {
		params["before"] = *before
	}
	if limit != nil {
		params["limit"] = *limit
	}
	if tag != nil {
		params["tag"] = *tag
	}

	var result []models.ConvertTradeResponse
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/convert/history", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetMonthlyStatement(ctx context.Context, month *string) ([]models.MonthlyStatement, error) {
	params := make(map[string]string)
	if month != nil {
		params["month"] = *month
	}

	var result []models.MonthlyStatement
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/asset/monthly-statement", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

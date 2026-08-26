package account

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tigusigalpa/okx-go/models"
)

// Client provides access to the OKX API endpoints in this package.
type Client struct {
	doFunc func(ctx context.Context, method, path string, params map[string]string, body interface{}, result interface{}) error
}

// NewClient creates a client for the API endpoints in this package.
func NewClient(doFunc func(ctx context.Context, method, path string, params map[string]string, body interface{}, result interface{}) error) *Client {
	return &Client{doFunc: doFunc}
}

// GetBalance invokes the corresponding OKX API operation.
func (c *Client) GetBalance(ctx context.Context, ccy *string) ([]models.Balance, error) {
	params := make(map[string]string)
	if ccy != nil {
		params["ccy"] = *ccy
	}

	var result []models.Balance
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/balance", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetPositions invokes the corresponding OKX API operation.
func (c *Client) GetPositions(ctx context.Context, instType *string, instID *string) ([]models.Position, error) {
	params := make(map[string]string)
	if instType != nil {
		params["instType"] = *instType
	}
	if instID != nil {
		params["instId"] = *instID
	}

	var result []models.Position
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/positions", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetPositionsHistory invokes the corresponding OKX API operation.
func (c *Client) GetPositionsHistory(ctx context.Context, instType *string, instID *string, mgnMode *string, positionType *string, posID *string, after *string, before *string, limit *string) ([]models.PositionHistory, error) {
	params := make(map[string]string)
	if instType != nil {
		params["instType"] = *instType
	}
	if instID != nil {
		params["instId"] = *instID
	}
	if mgnMode != nil {
		params["mgnMode"] = *mgnMode
	}
	if positionType != nil {
		params["type"] = *positionType
	}
	if posID != nil {
		params["posId"] = *posID
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

	var result []models.PositionHistory
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/positions-history", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAccountConfig invokes the corresponding OKX API operation.
func (c *Client) GetAccountConfig(ctx context.Context) ([]models.AccountConfig, error) {
	var result []models.AccountConfig
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/config", nil, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetPositionMode invokes the corresponding OKX API operation.
func (c *Client) SetPositionMode(ctx context.Context, posMode string) error {
	body := models.SetPositionModeRequest{
		PosMode: posMode,
	}
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/account/set-position-mode", nil, body, nil); err != nil {
		return err
	}
	return nil
}

// SetLeverage invokes the corresponding OKX API operation.
func (c *Client) SetLeverage(ctx context.Context, req models.SetLeverageRequest) ([]models.LeverageInfo, error) {
	var result []models.LeverageInfo
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/account/set-leverage", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMaxSize invokes the corresponding OKX API operation.
func (c *Client) GetMaxSize(ctx context.Context, instID string, tdMode string, ccy *string, px *string) ([]models.MaxSize, error) {
	params := map[string]string{
		"instId": instID,
		"tdMode": tdMode,
	}
	if ccy != nil {
		params["ccy"] = *ccy
	}
	if px != nil {
		params["px"] = *px
	}

	var result []models.MaxSize
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/max-size", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMaxAvailSize invokes the corresponding OKX API operation.
func (c *Client) GetMaxAvailSize(ctx context.Context, instID string, tdMode string, ccy *string, reduceOnly *bool, unSpotOffset *bool, quickMgnType *string) ([]models.MaxAvailSize, error) {
	params := map[string]string{
		"instId": instID,
		"tdMode": tdMode,
	}
	if ccy != nil {
		params["ccy"] = *ccy
	}
	if reduceOnly != nil {
		params["reduceOnly"] = fmt.Sprintf("%t", *reduceOnly)
	}
	if unSpotOffset != nil {
		params["unSpotOffset"] = fmt.Sprintf("%t", *unSpotOffset)
	}
	if quickMgnType != nil {
		params["quickMgnType"] = *quickMgnType
	}

	var result []models.MaxAvailSize
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/max-avail-size", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMaxLoan invokes the corresponding OKX API operation.
func (c *Client) GetMaxLoan(ctx context.Context, instID string, mgnMode string, mgnCcy string) ([]models.MaxLoan, error) {
	params := map[string]string{
		"instId":  instID,
		"mgnMode": mgnMode,
		"mgnCcy":  mgnCcy,
	}

	var result []models.MaxLoan
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/max-loan", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetTradeFee invokes the corresponding OKX API operation.
func (c *Client) GetTradeFee(ctx context.Context, instType string, instID *string, uly *string, category *string, instFamily *string) ([]models.TradeFee, error) {
	params := map[string]string{
		"instType": instType,
	}
	if instID != nil {
		params["instId"] = *instID
	}
	if uly != nil {
		params["uly"] = *uly
	}
	if category != nil {
		params["category"] = *category
	}
	if instFamily != nil {
		params["instFamily"] = *instFamily
	}

	var result []models.TradeFee
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/trade-fee", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetInterestAccrued invokes the corresponding OKX API operation.
func (c *Client) GetInterestAccrued(ctx context.Context, instID *string, ccy *string, mgnMode *string, after *string, before *string, limit *string) ([]models.InterestAccrued, error) {
	params := make(map[string]string)
	if instID != nil {
		params["instId"] = *instID
	}
	if ccy != nil {
		params["ccy"] = *ccy
	}
	if mgnMode != nil {
		params["mgnMode"] = *mgnMode
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

	var result []models.InterestAccrued
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/interest-accrued", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetInterestRate invokes the corresponding OKX API operation.
func (c *Client) GetInterestRate(ctx context.Context, ccy *string) ([]models.InterestRate, error) {
	params := make(map[string]string)
	if ccy != nil {
		params["ccy"] = *ccy
	}

	var result []models.InterestRate
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/interest-rate", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetGreeks invokes the corresponding OKX API operation.
func (c *Client) SetGreeks(ctx context.Context, greeksType string) error {
	body := models.SetGreeksRequest{
		GreeksType: greeksType,
	}
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/account/set-greeks", nil, body, nil); err != nil {
		return err
	}
	return nil
}

// GetGreeks invokes the corresponding OKX API operation.
func (c *Client) GetGreeks(ctx context.Context, ccy *string) ([]models.Greeks, error) {
	params := make(map[string]string)
	if ccy != nil {
		params["ccy"] = *ccy
	}

	var result []models.Greeks
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/greeks", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMaxWithdrawal invokes the corresponding OKX API operation.
func (c *Client) GetMaxWithdrawal(ctx context.Context, ccy *string) ([]models.MaxWithdrawal, error) {
	params := make(map[string]string)
	if ccy != nil {
		params["ccy"] = *ccy
	}

	var result []models.MaxWithdrawal
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/max-withdrawal", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetRiskState invokes the corresponding OKX API operation.
func (c *Client) GetRiskState(ctx context.Context) ([]models.RiskState, error) {
	var result []models.RiskState
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/risk-state", nil, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// BorrowRepay invokes the corresponding OKX API operation.
func (c *Client) BorrowRepay(ctx context.Context, req models.BorrowRepayRequest) error {
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/account/spot-manual-borrow-repay", nil, req, nil); err != nil {
		return err
	}
	return nil
}

// GetBorrowRepayHistory invokes the corresponding OKX API operation.
func (c *Client) GetBorrowRepayHistory(ctx context.Context, ccy *string, after *string, before *string, limit *string) ([]models.BorrowRepayHistory, error) {
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

	var result []models.BorrowRepayHistory
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/spot-manual-borrow-repay-history", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBills invokes the corresponding OKX API operation.
func (c *Client) GetBills(ctx context.Context, instType *string, ccy *string, mgnMode *string, ctType *string, billType *string, subType *string, after *string, before *string, begin *string, end *string, limit *string) ([]models.Bill, error) {
	params := make(map[string]string)
	if instType != nil {
		params["instType"] = *instType
	}
	if ccy != nil {
		params["ccy"] = *ccy
	}
	if mgnMode != nil {
		params["mgnMode"] = *mgnMode
	}
	if ctType != nil {
		params["ctType"] = *ctType
	}
	if billType != nil {
		params["type"] = *billType
	}
	if subType != nil {
		params["subType"] = *subType
	}
	if after != nil {
		params["after"] = *after
	}
	if before != nil {
		params["before"] = *before
	}
	if begin != nil {
		params["begin"] = *begin
	}
	if end != nil {
		params["end"] = *end
	}
	if limit != nil {
		params["limit"] = *limit
	}

	var result []models.Bill
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/bills", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBillsHistory invokes the corresponding OKX API operation.
func (c *Client) GetBillsHistory(ctx context.Context, instType *string, ccy *string, mgnMode *string, ctType *string, billType *string, subType *string, after *string, before *string, begin *string, end *string, limit *string) ([]models.Bill, error) {
	params := make(map[string]string)
	if instType != nil {
		params["instType"] = *instType
	}
	if ccy != nil {
		params["ccy"] = *ccy
	}
	if mgnMode != nil {
		params["mgnMode"] = *mgnMode
	}
	if ctType != nil {
		params["ctType"] = *ctType
	}
	if billType != nil {
		params["type"] = *billType
	}
	if subType != nil {
		params["subType"] = *subType
	}
	if after != nil {
		params["after"] = *after
	}
	if before != nil {
		params["before"] = *before
	}
	if begin != nil {
		params["begin"] = *begin
	}
	if end != nil {
		params["end"] = *end
	}
	if limit != nil {
		params["limit"] = *limit
	}

	var result []models.Bill
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/bills-history", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAccountLevel invokes the corresponding OKX API operation.
func (c *Client) GetAccountLevel(ctx context.Context) ([]models.AccountLevel, error) {
	var result []models.AccountLevel
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/account-level", nil, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetPositionMarginBalance invokes the corresponding OKX API operation.
func (c *Client) SetPositionMarginBalance(ctx context.Context, req models.PositionMarginBalanceRequest) error {
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/account/position/margin-balance", nil, req, nil); err != nil {
		return err
	}
	return nil
}

// GetLeverageInfo invokes the corresponding OKX API operation.
func (c *Client) GetLeverageInfo(ctx context.Context, instID string, mgnMode string) ([]models.LeverageInfo, error) {
	params := map[string]string{
		"instId":  instID,
		"mgnMode": mgnMode,
	}

	var result []models.LeverageInfo
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/leverage-info", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetInterestLimits invokes the corresponding OKX API operation.
func (c *Client) GetInterestLimits(ctx context.Context, interestType *string, ccy *string) ([]models.InterestLimits, error) {
	params := make(map[string]string)
	if interestType != nil {
		params["type"] = *interestType
	}
	if ccy != nil {
		params["ccy"] = *ccy
	}

	var result []models.InterestLimits
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/interest-limits", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMMPConfig invokes the corresponding OKX API operation.
func (c *Client) GetMMPConfig(ctx context.Context, instFamily string) ([]models.MMPConfig, error) {
	params := map[string]string{
		"instFamily": instFamily,
	}

	var result []models.MMPConfig
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/mmp-config", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetMMPConfig invokes the corresponding OKX API operation.
func (c *Client) SetMMPConfig(ctx context.Context, instFamily, timeInterval, frozenInterval, qtyLimit string) error {
	body := models.MMPConfig{
		InstFamily:     instFamily,
		TimeInterval:   timeInterval,
		FrozenInterval: frozenInterval,
		QtyLimit:       qtyLimit,
	}
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/account/mmp-config", nil, body, nil); err != nil {
		return err
	}
	return nil
}

// ResetMMPConfig invokes the corresponding OKX API operation.
func (c *Client) ResetMMPConfig(ctx context.Context, instFamily string) error {
	body := map[string]string{
		"instFamily": instFamily,
	}
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/account/mmp-reset", nil, body, nil); err != nil {
		return err
	}
	return nil
}

// GetMMPState invokes the corresponding OKX API operation.
func (c *Client) GetMMPState(ctx context.Context, instFamily string) ([]models.MMPState, error) {
	params := map[string]string{
		"instFamily": instFamily,
	}

	var result []models.MMPState
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/account/mmp-state", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

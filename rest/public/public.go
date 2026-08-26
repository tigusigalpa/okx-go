package public

import (
	"context"
	"net/http"

	"github.com/tigusigalpa/okx-go/models"
)

// Client provides access to the OKX API endpoints in this package.
type Client struct {
	doPublicFunc func(ctx context.Context, method, path string, params map[string]string, result interface{}) error
}

// NewClient creates a client for the API endpoints in this package.
func NewClient(doPublicFunc func(ctx context.Context, method, path string, params map[string]string, result interface{}) error) *Client {
	return &Client{doPublicFunc: doPublicFunc}
}

// GetInstruments invokes the corresponding OKX API operation.
func (c *Client) GetInstruments(ctx context.Context, instType string, uly *string, instFamily *string, instID *string) ([]models.Instrument, error) {
	params := map[string]string{
		"instType": instType,
	}
	if uly != nil {
		params["uly"] = *uly
	}
	if instFamily != nil {
		params["instFamily"] = *instFamily
	}
	if instID != nil {
		params["instId"] = *instID
	}

	var result []models.Instrument
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/instruments", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDeliveryExerciseHistory invokes the corresponding OKX API operation.
func (c *Client) GetDeliveryExerciseHistory(ctx context.Context, instType string, uly *string, after *string, before *string, limit *string) ([]models.DeliveryExerciseHistory, error) {
	params := map[string]string{
		"instType": instType,
	}
	if uly != nil {
		params["uly"] = *uly
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

	var result []models.DeliveryExerciseHistory
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/delivery-exercise-history", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetOpenInterest invokes the corresponding OKX API operation.
func (c *Client) GetOpenInterest(ctx context.Context, instType string, uly *string, instFamily *string, instID *string) ([]models.OpenInterest, error) {
	params := map[string]string{
		"instType": instType,
	}
	if uly != nil {
		params["uly"] = *uly
	}
	if instFamily != nil {
		params["instFamily"] = *instFamily
	}
	if instID != nil {
		params["instId"] = *instID
	}

	var result []models.OpenInterest
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/open-interest", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetFundingRate invokes the corresponding OKX API operation.
func (c *Client) GetFundingRate(ctx context.Context, instID string) ([]models.FundingRate, error) {
	params := map[string]string{
		"instId": instID,
	}

	var result []models.FundingRate
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/funding-rate", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetFundingRateHistory invokes the corresponding OKX API operation.
func (c *Client) GetFundingRateHistory(ctx context.Context, instID string, after *string, before *string, limit *string) ([]models.FundingRateHistory, error) {
	params := map[string]string{
		"instId": instID,
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

	var result []models.FundingRateHistory
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/funding-rate-history", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetPriceLimit invokes the corresponding OKX API operation.
func (c *Client) GetPriceLimit(ctx context.Context, instID string) ([]models.PriceLimit, error) {
	params := map[string]string{
		"instId": instID,
	}

	var result []models.PriceLimit
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/price-limit", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetOptionSummary invokes the corresponding OKX API operation.
func (c *Client) GetOptionSummary(ctx context.Context, uly string, expTime *string) ([]models.OptionSummary, error) {
	params := map[string]string{
		"uly": uly,
	}
	if expTime != nil {
		params["expTime"] = *expTime
	}

	var result []models.OptionSummary
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/opt-summary", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetEstimatedPrice invokes the corresponding OKX API operation.
func (c *Client) GetEstimatedPrice(ctx context.Context, instID string) ([]models.EstimatedPrice, error) {
	params := map[string]string{
		"instId": instID,
	}

	var result []models.EstimatedPrice
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/estimated-delivery-exercise-price", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDiscountRateInterestFreeQuota invokes the corresponding OKX API operation.
func (c *Client) GetDiscountRateInterestFreeQuota(ctx context.Context, ccy *string) ([]models.DiscountRateInterestFreeQuota, error) {
	params := make(map[string]string)
	if ccy != nil {
		params["ccy"] = *ccy
	}

	var result []models.DiscountRateInterestFreeQuota
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/discount-rate-interest-free-quota", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetSystemTime invokes the corresponding OKX API operation.
func (c *Client) GetSystemTime(ctx context.Context) ([]models.SystemTime, error) {
	var result []models.SystemTime
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/time", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetLiquidationOrders invokes the corresponding OKX API operation.
func (c *Client) GetLiquidationOrders(ctx context.Context, instType string, mgnMode *string, instID *string, ccy *string, uly *string, alias *string, state *string, before *string, after *string, limit *string) ([]models.LiquidationOrder, error) {
	params := map[string]string{
		"instType": instType,
	}
	if mgnMode != nil {
		params["mgnMode"] = *mgnMode
	}
	if instID != nil {
		params["instId"] = *instID
	}
	if ccy != nil {
		params["ccy"] = *ccy
	}
	if uly != nil {
		params["uly"] = *uly
	}
	if alias != nil {
		params["alias"] = *alias
	}
	if state != nil {
		params["state"] = *state
	}
	if before != nil {
		params["before"] = *before
	}
	if after != nil {
		params["after"] = *after
	}
	if limit != nil {
		params["limit"] = *limit
	}

	var result []models.LiquidationOrder
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/liquidation-orders", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMarkPrice invokes the corresponding OKX API operation.
func (c *Client) GetMarkPrice(ctx context.Context, instType string, uly *string, instFamily *string, instID *string) ([]models.MarkPrice, error) {
	params := map[string]string{
		"instType": instType,
	}
	if uly != nil {
		params["uly"] = *uly
	}
	if instFamily != nil {
		params["instFamily"] = *instFamily
	}
	if instID != nil {
		params["instId"] = *instID
	}

	var result []models.MarkPrice
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/mark-price", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetPositionTiers invokes the corresponding OKX API operation.
func (c *Client) GetPositionTiers(ctx context.Context, instType string, tdMode string, uly *string, instFamily *string, instID *string, ccy *string, tier *string) ([]models.PositionTier, error) {
	params := map[string]string{
		"instType": instType,
		"tdMode":   tdMode,
	}
	if uly != nil {
		params["uly"] = *uly
	}
	if instFamily != nil {
		params["instFamily"] = *instFamily
	}
	if instID != nil {
		params["instId"] = *instID
	}
	if ccy != nil {
		params["ccy"] = *ccy
	}
	if tier != nil {
		params["tier"] = *tier
	}

	var result []models.PositionTier
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/position-tiers", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetInterestRateLoanQuota invokes the corresponding OKX API operation.
func (c *Client) GetInterestRateLoanQuota(ctx context.Context) ([]models.InterestRateLoanQuota, error) {
	var result []models.InterestRateLoanQuota
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/interest-rate-loan-quota", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetVIPInterestRateLoanQuota invokes the corresponding OKX API operation.
func (c *Client) GetVIPInterestRateLoanQuota(ctx context.Context) ([]models.VIPInterestRateLoanQuota, error) {
	var result []models.VIPInterestRateLoanQuota
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/vip-interest-rate-loan-quota", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetUnderlying invokes the corresponding OKX API operation.
func (c *Client) GetUnderlying(ctx context.Context, instType string) ([]models.Underlying, error) {
	params := map[string]string{
		"instType": instType,
	}

	var result []models.Underlying
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/underlying", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetInsuranceFund invokes the corresponding OKX API operation.
func (c *Client) GetInsuranceFund(ctx context.Context, instType string, fundType *string, uly *string, ccy *string, before *string, after *string, limit *string) ([]models.InsuranceFund, error) {
	params := map[string]string{
		"instType": instType,
	}
	if fundType != nil {
		params["type"] = *fundType
	}
	if uly != nil {
		params["uly"] = *uly
	}
	if ccy != nil {
		params["ccy"] = *ccy
	}
	if before != nil {
		params["before"] = *before
	}
	if after != nil {
		params["after"] = *after
	}
	if limit != nil {
		params["limit"] = *limit
	}

	var result []models.InsuranceFund
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/insurance-fund", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ConvertContractCoin invokes the corresponding OKX API operation.
func (c *Client) ConvertContractCoin(ctx context.Context, instID string, sz string, px *string, conversionType *string, unit *string) ([]models.UnitConvert, error) {
	params := map[string]string{
		"instId": instID,
		"sz":     sz,
	}
	if px != nil {
		params["px"] = *px
	}
	if conversionType != nil {
		params["type"] = *conversionType
	}
	if unit != nil {
		params["unit"] = *unit
	}

	var result []models.UnitConvert
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/convert-contract-coin", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetEconomicCalendar invokes the corresponding OKX API operation.
func (c *Client) GetEconomicCalendar(ctx context.Context, region *string, importance *string, before *string, after *string, limit *string) ([]models.EconomicCalendar, error) {
	params := make(map[string]string)
	if region != nil {
		params["region"] = *region
	}
	if importance != nil {
		params["importance"] = *importance
	}
	if before != nil {
		params["before"] = *before
	}
	if after != nil {
		params["after"] = *after
	}
	if limit != nil {
		params["limit"] = *limit
	}

	var result []models.EconomicCalendar
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/public/economic-calendar", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

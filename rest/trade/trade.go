package trade

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

func (c *Client) PlaceOrder(ctx context.Context, req models.PlaceOrderRequest) ([]models.PlaceOrderResponse, error) {
	var result []models.PlaceOrderResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/trade/order", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) PlaceBatchOrders(ctx context.Context, reqs []models.PlaceOrderRequest) ([]models.PlaceOrderResponse, error) {
	var result []models.PlaceOrderResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/trade/batch-orders", nil, reqs, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) CancelOrder(ctx context.Context, req models.CancelOrderRequest) ([]models.CancelOrderResponse, error) {
	var result []models.CancelOrderResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/trade/cancel-order", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) CancelBatchOrders(ctx context.Context, reqs []models.CancelOrderRequest) ([]models.CancelOrderResponse, error) {
	var result []models.CancelOrderResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/trade/cancel-batch-orders", nil, reqs, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) AmendOrder(ctx context.Context, req models.AmendOrderRequest) ([]models.AmendOrderResponse, error) {
	var result []models.AmendOrderResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/trade/amend-order", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) AmendBatchOrders(ctx context.Context, reqs []models.AmendOrderRequest) ([]models.AmendOrderResponse, error) {
	var result []models.AmendOrderResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/trade/amend-batch-orders", nil, reqs, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) ClosePosition(ctx context.Context, req models.ClosePositionRequest) ([]models.PlaceOrderResponse, error) {
	var result []models.PlaceOrderResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/trade/close-position", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetOrder(ctx context.Context, instID string, ordID *string, clOrdID *string) ([]models.Order, error) {
	params := map[string]string{
		"instId": instID,
	}
	if ordID != nil {
		params["ordId"] = *ordID
	}
	if clOrdID != nil {
		params["clOrdId"] = *clOrdID
	}

	var result []models.Order
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/trade/order", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetOrdersPending(ctx context.Context, instType *string, uly *string, instFamily *string, instID *string, ordType *string, state *string, after *string, before *string, limit *string) ([]models.Order, error) {
	params := make(map[string]string)
	if instType != nil {
		params["instType"] = *instType
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
	if ordType != nil {
		params["ordType"] = *ordType
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

	var result []models.Order
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/trade/orders-pending", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetOrdersHistory(ctx context.Context, instType string, uly *string, instFamily *string, instID *string, ordType *string, state *string, category *string, after *string, before *string, begin *string, end *string, limit *string) ([]models.Order, error) {
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
	if ordType != nil {
		params["ordType"] = *ordType
	}
	if state != nil {
		params["state"] = *state
	}
	if category != nil {
		params["category"] = *category
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

	var result []models.Order
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/trade/orders-history", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetOrdersHistoryArchive(ctx context.Context, instType string, uly *string, instFamily *string, instID *string, ordType *string, state *string, category *string, after *string, before *string, begin *string, end *string, limit *string) ([]models.Order, error) {
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
	if ordType != nil {
		params["ordType"] = *ordType
	}
	if state != nil {
		params["state"] = *state
	}
	if category != nil {
		params["category"] = *category
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

	var result []models.Order
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/trade/orders-history-archive", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetFills(ctx context.Context, instType *string, uly *string, instFamily *string, instID *string, ordID *string, after *string, before *string, begin *string, end *string, limit *string) ([]models.Fill, error) {
	params := make(map[string]string)
	if instType != nil {
		params["instType"] = *instType
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
	if ordID != nil {
		params["ordId"] = *ordID
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

	var result []models.Fill
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/trade/fills", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetFillsHistory(ctx context.Context, instType string, uly *string, instFamily *string, instID *string, ordID *string, after *string, before *string, begin *string, end *string, limit *string) ([]models.Fill, error) {
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
	if ordID != nil {
		params["ordId"] = *ordID
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

	var result []models.Fill
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/trade/fills-history", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) PlaceAlgoOrder(ctx context.Context, req models.PlaceAlgoOrderRequest) ([]models.PlaceAlgoOrderResponse, error) {
	var result []models.PlaceAlgoOrderResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/trade/order-algo", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) CancelAlgoOrder(ctx context.Context, reqs []models.CancelAlgoOrderRequest) ([]models.CancelAlgoOrderResponse, error) {
	var result []models.CancelAlgoOrderResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/trade/cancel-algos", nil, reqs, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetAlgoOrdersPending(ctx context.Context, ordType string, algoID *string, instType *string, instID *string, after *string, before *string, limit *string) ([]models.AlgoOrder, error) {
	params := map[string]string{
		"ordType": ordType,
	}
	if algoID != nil {
		params["algoId"] = *algoID
	}
	if instType != nil {
		params["instType"] = *instType
	}
	if instID != nil {
		params["instId"] = *instID
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

	var result []models.AlgoOrder
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/trade/orders-algo-pending", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetAlgoOrdersHistory(ctx context.Context, ordType string, state *string, algoID *string, instType *string, instID *string, after *string, before *string, limit *string) ([]models.AlgoOrder, error) {
	params := map[string]string{
		"ordType": ordType,
	}
	if state != nil {
		params["state"] = *state
	}
	if algoID != nil {
		params["algoId"] = *algoID
	}
	if instType != nil {
		params["instType"] = *instType
	}
	if instID != nil {
		params["instId"] = *instID
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

	var result []models.AlgoOrder
	if err := c.doFunc(ctx, http.MethodGet, "/api/v5/trade/orders-algo-history", params, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) MassCancel(ctx context.Context, req models.MassCancelRequest) ([]models.MassCancelResponse, error) {
	var result []models.MassCancelResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/trade/mass-cancel", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) CancelAllAfter(ctx context.Context, req models.CancelAllAfterRequest) ([]models.CancelAllAfterResponse, error) {
	var result []models.CancelAllAfterResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/trade/cancel-all-after", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) EasyConvert(ctx context.Context, req models.EasyConvertRequest) ([]models.EasyConvertResponse, error) {
	var result []models.EasyConvertResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/trade/easy-convert", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) OneClickRepay(ctx context.Context, req models.OneClickRepayRequest) ([]models.OneClickRepayResponse, error) {
	var result []models.OneClickRepayResponse
	if err := c.doFunc(ctx, http.MethodPost, "/api/v5/trade/one-click-repay", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

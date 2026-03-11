package market

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

func (c *Client) GetTickers(ctx context.Context, instType string, uly *string, instFamily *string) ([]models.Ticker, error) {
	params := map[string]string{
		"instType": instType,
	}
	if uly != nil {
		params["uly"] = *uly
	}
	if instFamily != nil {
		params["instFamily"] = *instFamily
	}

	var result []models.Ticker
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/tickers", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetTicker(ctx context.Context, instID string) ([]models.Ticker, error) {
	params := map[string]string{
		"instId": instID,
	}

	var result []models.Ticker
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/ticker", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetIndexTickers(ctx context.Context, quoteCcy *string, instID *string) ([]models.IndexTicker, error) {
	params := make(map[string]string)
	if quoteCcy != nil {
		params["quoteCcy"] = *quoteCcy
	}
	if instID != nil {
		params["instId"] = *instID
	}

	var result []models.IndexTicker
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/index-tickers", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetOrderBook(ctx context.Context, instID string, sz *string) ([]models.OrderBook, error) {
	params := map[string]string{
		"instId": instID,
	}
	if sz != nil {
		params["sz"] = *sz
	}

	var result []models.OrderBook
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/books", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetOrderBookFull(ctx context.Context, instID string, sz *string) ([]models.OrderBook, error) {
	params := map[string]string{
		"instId": instID,
	}
	if sz != nil {
		params["sz"] = *sz
	}

	var result []models.OrderBook
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/books-full", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetOrderBookLite(ctx context.Context, instID string) ([]models.OrderBook, error) {
	params := map[string]string{
		"instId": instID,
	}

	var result []models.OrderBook
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/books-lite", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetCandles(ctx context.Context, instID string, bar *string, after *string, before *string, limit *string) ([]models.Candle, error) {
	params := map[string]string{
		"instId": instID,
	}
	if bar != nil {
		params["bar"] = *bar
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

	var result []models.Candle
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/candles", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetHistoryCandles(ctx context.Context, instID string, bar *string, after *string, before *string, limit *string) ([]models.Candle, error) {
	params := map[string]string{
		"instId": instID,
	}
	if bar != nil {
		params["bar"] = *bar
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

	var result []models.Candle
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/history-candles", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetIndexCandles(ctx context.Context, instID string, bar *string, after *string, before *string, limit *string) ([]models.Candle, error) {
	params := map[string]string{
		"instId": instID,
	}
	if bar != nil {
		params["bar"] = *bar
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

	var result []models.Candle
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/index-candles", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetHistoryIndexCandles(ctx context.Context, instID string, bar *string, after *string, before *string, limit *string) ([]models.Candle, error) {
	params := map[string]string{
		"instId": instID,
	}
	if bar != nil {
		params["bar"] = *bar
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

	var result []models.Candle
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/history-index-candles", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetMarkPriceCandles(ctx context.Context, instID string, bar *string, after *string, before *string, limit *string) ([]models.Candle, error) {
	params := map[string]string{
		"instId": instID,
	}
	if bar != nil {
		params["bar"] = *bar
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

	var result []models.Candle
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/mark-price-candles", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetHistoryMarkPriceCandles(ctx context.Context, instID string, bar *string, after *string, before *string, limit *string) ([]models.Candle, error) {
	params := map[string]string{
		"instId": instID,
	}
	if bar != nil {
		params["bar"] = *bar
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

	var result []models.Candle
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/history-mark-price-candles", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetTrades(ctx context.Context, instID string, limit *string) ([]models.Trade, error) {
	params := map[string]string{
		"instId": instID,
	}
	if limit != nil {
		params["limit"] = *limit
	}

	var result []models.Trade
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/trades", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetHistoryTrades(ctx context.Context, instID string, type_ *string, after *string, before *string, limit *string) ([]models.Trade, error) {
	params := map[string]string{
		"instId": instID,
	}
	if type_ != nil {
		params["type"] = *type_
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

	var result []models.Trade
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/history-trades", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) Get24hVolume(ctx context.Context) ([]models.Platform24Volume, error) {
	var result []models.Platform24Volume
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/platform-24-volume", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetOpenOracle(ctx context.Context) ([]models.OpenOracle, error) {
	var result []models.OpenOracle
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/open-oracle", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetExchangeRate(ctx context.Context) ([]models.ExchangeRate, error) {
	var result []models.ExchangeRate
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/exchange-rate", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetIndexComponents(ctx context.Context, index string) ([]models.IndexComponents, error) {
	params := map[string]string{
		"index": index,
	}

	var result []models.IndexComponents
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/index-components", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetBlockTicker(ctx context.Context, instID string) ([]models.BlockTicker, error) {
	params := map[string]string{
		"instId": instID,
	}

	var result []models.BlockTicker
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/block-ticker", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetBlockTrades(ctx context.Context, instID string) ([]models.BlockTrade, error) {
	params := map[string]string{
		"instId": instID,
	}

	var result []models.BlockTrade
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/block-trades", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetUnderlying(ctx context.Context, instType string) ([]models.Underlying, error) {
	params := map[string]string{
		"instType": instType,
	}

	var result []models.Underlying
	if err := c.doPublicFunc(ctx, http.MethodGet, "/api/v5/market/underlying", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

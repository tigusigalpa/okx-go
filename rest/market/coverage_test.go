package market

import (
	"context"
	"testing"
)

func TestAllEndpointsAcceptPopulatedOptionalArguments(t *testing.T) {
	text := "value"
	enabled := true
	_ = enabled
	client := NewClient(func(_ context.Context, _ string, _ string, _ map[string]string, _ interface{}) error { return nil })
	if _, err := client.GetTickers(context.Background(), "value", &text, &text); err != nil {
		t.Fatalf("GetTickers: %v", err)
	}
	if _, err := client.GetTicker(context.Background(), "value"); err != nil {
		t.Fatalf("GetTicker: %v", err)
	}
	if _, err := client.GetIndexTickers(context.Background(), &text, &text); err != nil {
		t.Fatalf("GetIndexTickers: %v", err)
	}
	if _, err := client.GetOrderBook(context.Background(), "value", &text); err != nil {
		t.Fatalf("GetOrderBook: %v", err)
	}
	if _, err := client.GetOrderBookFull(context.Background(), "value", &text); err != nil {
		t.Fatalf("GetOrderBookFull: %v", err)
	}
	if _, err := client.GetOrderBookLite(context.Background(), "value"); err != nil {
		t.Fatalf("GetOrderBookLite: %v", err)
	}
	if _, err := client.GetCandles(context.Background(), "value", &text, &text, &text, &text); err != nil {
		t.Fatalf("GetCandles: %v", err)
	}
	if _, err := client.GetHistoryCandles(context.Background(), "value", &text, &text, &text, &text); err != nil {
		t.Fatalf("GetHistoryCandles: %v", err)
	}
	if _, err := client.GetIndexCandles(context.Background(), "value", &text, &text, &text, &text); err != nil {
		t.Fatalf("GetIndexCandles: %v", err)
	}
	if _, err := client.GetHistoryIndexCandles(context.Background(), "value", &text, &text, &text, &text); err != nil {
		t.Fatalf("GetHistoryIndexCandles: %v", err)
	}
	if _, err := client.GetMarkPriceCandles(context.Background(), "value", &text, &text, &text, &text); err != nil {
		t.Fatalf("GetMarkPriceCandles: %v", err)
	}
	if _, err := client.GetHistoryMarkPriceCandles(context.Background(), "value", &text, &text, &text, &text); err != nil {
		t.Fatalf("GetHistoryMarkPriceCandles: %v", err)
	}
	if _, err := client.GetTrades(context.Background(), "value", &text); err != nil {
		t.Fatalf("GetTrades: %v", err)
	}
	if _, err := client.GetHistoryTrades(context.Background(), "value", &text, &text, &text, &text); err != nil {
		t.Fatalf("GetHistoryTrades: %v", err)
	}
	if _, err := client.Get24hVolume(context.Background()); err != nil {
		t.Fatalf("Get24hVolume: %v", err)
	}
	if _, err := client.GetOpenOracle(context.Background()); err != nil {
		t.Fatalf("GetOpenOracle: %v", err)
	}
	if _, err := client.GetExchangeRate(context.Background()); err != nil {
		t.Fatalf("GetExchangeRate: %v", err)
	}
	if _, err := client.GetIndexComponents(context.Background(), "value"); err != nil {
		t.Fatalf("GetIndexComponents: %v", err)
	}
	if _, err := client.GetBlockTicker(context.Background(), "value"); err != nil {
		t.Fatalf("GetBlockTicker: %v", err)
	}
	if _, err := client.GetBlockTrades(context.Background(), "value"); err != nil {
		t.Fatalf("GetBlockTrades: %v", err)
	}
	if _, err := client.GetUnderlying(context.Background(), "value"); err != nil {
		t.Fatalf("GetUnderlying: %v", err)
	}
}

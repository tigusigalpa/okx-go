package public

import (
	"context"
	"testing"
)

func TestAllEndpointsAcceptPopulatedOptionalArguments(t *testing.T) {
	text := "value"
	enabled := true
	_ = enabled
	client := NewClient(func(_ context.Context, _ string, _ string, _ map[string]string, _ interface{}) error { return nil })
	if _, err := client.GetInstruments(context.Background(), "value", &text, &text, &text); err != nil {
		t.Fatalf("GetInstruments: %v", err)
	}
	if _, err := client.GetDeliveryExerciseHistory(context.Background(), "value", &text, &text, &text, &text); err != nil {
		t.Fatalf("GetDeliveryExerciseHistory: %v", err)
	}
	if _, err := client.GetOpenInterest(context.Background(), "value", &text, &text, &text); err != nil {
		t.Fatalf("GetOpenInterest: %v", err)
	}
	if _, err := client.GetFundingRate(context.Background(), "value"); err != nil {
		t.Fatalf("GetFundingRate: %v", err)
	}
	if _, err := client.GetFundingRateHistory(context.Background(), "value", &text, &text, &text); err != nil {
		t.Fatalf("GetFundingRateHistory: %v", err)
	}
	if _, err := client.GetPriceLimit(context.Background(), "value"); err != nil {
		t.Fatalf("GetPriceLimit: %v", err)
	}
	if _, err := client.GetOptionSummary(context.Background(), "value", &text); err != nil {
		t.Fatalf("GetOptionSummary: %v", err)
	}
	if _, err := client.GetEstimatedPrice(context.Background(), "value"); err != nil {
		t.Fatalf("GetEstimatedPrice: %v", err)
	}
	if _, err := client.GetDiscountRateInterestFreeQuota(context.Background(), &text); err != nil {
		t.Fatalf("GetDiscountRateInterestFreeQuota: %v", err)
	}
	if _, err := client.GetSystemTime(context.Background()); err != nil {
		t.Fatalf("GetSystemTime: %v", err)
	}
	if _, err := client.GetLiquidationOrders(context.Background(), "value", &text, &text, &text, &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetLiquidationOrders: %v", err)
	}
	if _, err := client.GetMarkPrice(context.Background(), "value", &text, &text, &text); err != nil {
		t.Fatalf("GetMarkPrice: %v", err)
	}
	if _, err := client.GetPositionTiers(context.Background(), "value", "value", &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetPositionTiers: %v", err)
	}
	if _, err := client.GetInterestRateLoanQuota(context.Background()); err != nil {
		t.Fatalf("GetInterestRateLoanQuota: %v", err)
	}
	if _, err := client.GetVIPInterestRateLoanQuota(context.Background()); err != nil {
		t.Fatalf("GetVIPInterestRateLoanQuota: %v", err)
	}
	if _, err := client.GetUnderlying(context.Background(), "value"); err != nil {
		t.Fatalf("GetUnderlying: %v", err)
	}
	if _, err := client.GetInsuranceFund(context.Background(), "value", &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetInsuranceFund: %v", err)
	}
	if _, err := client.ConvertContractCoin(context.Background(), "value", "value", &text, &text, &text); err != nil {
		t.Fatalf("ConvertContractCoin: %v", err)
	}
	if _, err := client.GetEconomicCalendar(context.Background(), &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetEconomicCalendar: %v", err)
	}
}

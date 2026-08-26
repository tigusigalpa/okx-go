package asset

import (
	"context"
	"testing"

	"github.com/tigusigalpa/okx-go/models"
)

func TestAllEndpointsAcceptPopulatedOptionalArguments(t *testing.T) {
	text := "value"
	enabled := true
	_ = enabled
	client := NewClient(func(_ context.Context, _ string, _ string, _ map[string]string, _ interface{}, _ interface{}) error {
		return nil
	})
	if _, err := client.GetCurrencies(context.Background(), &text); err != nil {
		t.Fatalf("GetCurrencies: %v", err)
	}
	if _, err := client.GetBalances(context.Background(), &text); err != nil {
		t.Fatalf("GetBalances: %v", err)
	}
	if _, err := client.GetAssetValuation(context.Background(), &text); err != nil {
		t.Fatalf("GetAssetValuation: %v", err)
	}
	if _, err := client.Transfer(context.Background(), models.TransferRequest{}); err != nil {
		t.Fatalf("Transfer: %v", err)
	}
	if _, err := client.GetTransferState(context.Background(), &text, &text, &text); err != nil {
		t.Fatalf("GetTransferState: %v", err)
	}
	if _, err := client.Withdrawal(context.Background(), models.WithdrawalRequest{}); err != nil {
		t.Fatalf("Withdrawal: %v", err)
	}
	if err := client.CancelWithdrawal(context.Background(), "value"); err != nil {
		t.Fatalf("CancelWithdrawal: %v", err)
	}
	if _, err := client.GetWithdrawalHistory(context.Background(), &text, &text, &text, &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetWithdrawalHistory: %v", err)
	}
	if _, err := client.GetDepositAddress(context.Background(), "value"); err != nil {
		t.Fatalf("GetDepositAddress: %v", err)
	}
	if _, err := client.GetDepositHistory(context.Background(), &text, &text, &text, &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetDepositHistory: %v", err)
	}
	if _, err := client.GetBills(context.Background(), &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetBills: %v", err)
	}
	if _, err := client.GetSavingBalance(context.Background(), &text); err != nil {
		t.Fatalf("GetSavingBalance: %v", err)
	}
	if _, err := client.PurchaseRedempt(context.Background(), models.PurchaseRedemptRequest{}); err != nil {
		t.Fatalf("PurchaseRedempt: %v", err)
	}
	if _, err := client.SetLendingRate(context.Background(), "value", "value"); err != nil {
		t.Fatalf("SetLendingRate: %v", err)
	}
	if _, err := client.GetLendingRateSummary(context.Background(), &text); err != nil {
		t.Fatalf("GetLendingRateSummary: %v", err)
	}
	if _, err := client.GetLendingRateHistory(context.Background(), &text, &text, &text, &text); err != nil {
		t.Fatalf("GetLendingRateHistory: %v", err)
	}
	if _, err := client.GetConvertCurrencies(context.Background()); err != nil {
		t.Fatalf("GetConvertCurrencies: %v", err)
	}
	if _, err := client.GetConvertCurrencyPair(context.Background(), "value", "value"); err != nil {
		t.Fatalf("GetConvertCurrencyPair: %v", err)
	}
	if _, err := client.EstimateConvertQuote(context.Background(), models.ConvertEstimateQuoteRequest{}); err != nil {
		t.Fatalf("EstimateConvertQuote: %v", err)
	}
	if _, err := client.ConvertTrade(context.Background(), models.ConvertTradeRequest{}); err != nil {
		t.Fatalf("ConvertTrade: %v", err)
	}
	if _, err := client.GetConvertHistory(context.Background(), &text, &text, &text, &text); err != nil {
		t.Fatalf("GetConvertHistory: %v", err)
	}
	if _, err := client.GetMonthlyStatement(context.Background(), &text); err != nil {
		t.Fatalf("GetMonthlyStatement: %v", err)
	}
}

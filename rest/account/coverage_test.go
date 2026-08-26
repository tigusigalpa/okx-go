package account

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
	if _, err := client.GetBalance(context.Background(), &text); err != nil {
		t.Fatalf("GetBalance: %v", err)
	}
	if _, err := client.GetPositions(context.Background(), &text, &text); err != nil {
		t.Fatalf("GetPositions: %v", err)
	}
	if _, err := client.GetPositionsHistory(context.Background(), &text, &text, &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetPositionsHistory: %v", err)
	}
	if _, err := client.GetAccountConfig(context.Background()); err != nil {
		t.Fatalf("GetAccountConfig: %v", err)
	}
	if err := client.SetPositionMode(context.Background(), "value"); err != nil {
		t.Fatalf("SetPositionMode: %v", err)
	}
	if _, err := client.SetLeverage(context.Background(), models.SetLeverageRequest{}); err != nil {
		t.Fatalf("SetLeverage: %v", err)
	}
	if _, err := client.GetMaxSize(context.Background(), "value", "value", &text, &text); err != nil {
		t.Fatalf("GetMaxSize: %v", err)
	}
	if _, err := client.GetMaxAvailSize(context.Background(), "value", "value", &text, &enabled, &enabled, &text); err != nil {
		t.Fatalf("GetMaxAvailSize: %v", err)
	}
	if _, err := client.GetMaxLoan(context.Background(), "value", "value", "value"); err != nil {
		t.Fatalf("GetMaxLoan: %v", err)
	}
	if _, err := client.GetTradeFee(context.Background(), "value", &text, &text, &text, &text); err != nil {
		t.Fatalf("GetTradeFee: %v", err)
	}
	if _, err := client.GetInterestAccrued(context.Background(), &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetInterestAccrued: %v", err)
	}
	if _, err := client.GetInterestRate(context.Background(), &text); err != nil {
		t.Fatalf("GetInterestRate: %v", err)
	}
	if err := client.SetGreeks(context.Background(), "value"); err != nil {
		t.Fatalf("SetGreeks: %v", err)
	}
	if _, err := client.GetGreeks(context.Background(), &text); err != nil {
		t.Fatalf("GetGreeks: %v", err)
	}
	if _, err := client.GetMaxWithdrawal(context.Background(), &text); err != nil {
		t.Fatalf("GetMaxWithdrawal: %v", err)
	}
	if _, err := client.GetRiskState(context.Background()); err != nil {
		t.Fatalf("GetRiskState: %v", err)
	}
	if err := client.BorrowRepay(context.Background(), models.BorrowRepayRequest{}); err != nil {
		t.Fatalf("BorrowRepay: %v", err)
	}
	if _, err := client.GetBorrowRepayHistory(context.Background(), &text, &text, &text, &text); err != nil {
		t.Fatalf("GetBorrowRepayHistory: %v", err)
	}
	if _, err := client.GetBills(context.Background(), &text, &text, &text, &text, &text, &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetBills: %v", err)
	}
	if _, err := client.GetBillsHistory(context.Background(), &text, &text, &text, &text, &text, &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetBillsHistory: %v", err)
	}
	if _, err := client.GetAccountLevel(context.Background()); err != nil {
		t.Fatalf("GetAccountLevel: %v", err)
	}
	if err := client.SetPositionMarginBalance(context.Background(), models.PositionMarginBalanceRequest{}); err != nil {
		t.Fatalf("SetPositionMarginBalance: %v", err)
	}
	if _, err := client.GetLeverageInfo(context.Background(), "value", "value"); err != nil {
		t.Fatalf("GetLeverageInfo: %v", err)
	}
	if _, err := client.GetInterestLimits(context.Background(), &text, &text); err != nil {
		t.Fatalf("GetInterestLimits: %v", err)
	}
	if _, err := client.GetMMPConfig(context.Background(), "value"); err != nil {
		t.Fatalf("GetMMPConfig: %v", err)
	}
	if err := client.SetMMPConfig(context.Background(), "value", "value", "value", "value"); err != nil {
		t.Fatalf("SetMMPConfig: %v", err)
	}
	if err := client.ResetMMPConfig(context.Background(), "value"); err != nil {
		t.Fatalf("ResetMMPConfig: %v", err)
	}
	if _, err := client.GetMMPState(context.Background(), "value"); err != nil {
		t.Fatalf("GetMMPState: %v", err)
	}
}

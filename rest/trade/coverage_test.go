package trade

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
	if _, err := client.PlaceOrder(context.Background(), models.PlaceOrderRequest{}); err != nil {
		t.Fatalf("PlaceOrder: %v", err)
	}
	if _, err := client.PlaceBatchOrders(context.Background(), []models.PlaceOrderRequest{}); err != nil {
		t.Fatalf("PlaceBatchOrders: %v", err)
	}
	if _, err := client.CancelOrder(context.Background(), models.CancelOrderRequest{}); err != nil {
		t.Fatalf("CancelOrder: %v", err)
	}
	if _, err := client.CancelBatchOrders(context.Background(), []models.CancelOrderRequest{}); err != nil {
		t.Fatalf("CancelBatchOrders: %v", err)
	}
	if _, err := client.AmendOrder(context.Background(), models.AmendOrderRequest{}); err != nil {
		t.Fatalf("AmendOrder: %v", err)
	}
	if _, err := client.AmendBatchOrders(context.Background(), []models.AmendOrderRequest{}); err != nil {
		t.Fatalf("AmendBatchOrders: %v", err)
	}
	if _, err := client.ClosePosition(context.Background(), models.ClosePositionRequest{}); err != nil {
		t.Fatalf("ClosePosition: %v", err)
	}
	if _, err := client.GetOrder(context.Background(), "value", &text, &text); err != nil {
		t.Fatalf("GetOrder: %v", err)
	}
	if _, err := client.GetOrdersPending(context.Background(), &text, &text, &text, &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetOrdersPending: %v", err)
	}
	if _, err := client.GetOrdersHistory(context.Background(), "value", &text, &text, &text, &text, &text, &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetOrdersHistory: %v", err)
	}
	if _, err := client.GetOrdersHistoryArchive(context.Background(), "value", &text, &text, &text, &text, &text, &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetOrdersHistoryArchive: %v", err)
	}
	if _, err := client.GetFills(context.Background(), &text, &text, &text, &text, &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetFills: %v", err)
	}
	if _, err := client.GetFillsHistory(context.Background(), "value", &text, &text, &text, &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetFillsHistory: %v", err)
	}
	if _, err := client.PlaceAlgoOrder(context.Background(), models.PlaceAlgoOrderRequest{}); err != nil {
		t.Fatalf("PlaceAlgoOrder: %v", err)
	}
	if _, err := client.CancelAlgoOrder(context.Background(), []models.CancelAlgoOrderRequest{}); err != nil {
		t.Fatalf("CancelAlgoOrder: %v", err)
	}
	if _, err := client.GetAlgoOrdersPending(context.Background(), "value", &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetAlgoOrdersPending: %v", err)
	}
	if _, err := client.GetAlgoOrdersHistory(context.Background(), "value", &text, &text, &text, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetAlgoOrdersHistory: %v", err)
	}
	if _, err := client.MassCancel(context.Background(), models.MassCancelRequest{}); err != nil {
		t.Fatalf("MassCancel: %v", err)
	}
	if _, err := client.CancelAllAfter(context.Background(), models.CancelAllAfterRequest{}); err != nil {
		t.Fatalf("CancelAllAfter: %v", err)
	}
	if _, err := client.EasyConvert(context.Background(), models.EasyConvertRequest{}); err != nil {
		t.Fatalf("EasyConvert: %v", err)
	}
	if _, err := client.OneClickRepay(context.Background(), models.OneClickRepayRequest{}); err != nil {
		t.Fatalf("OneClickRepay: %v", err)
	}
}

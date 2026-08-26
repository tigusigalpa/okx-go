package users

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
	if _, err := client.GetSubAccountList(context.Background(), &enabled, &text, &text, &text, &text); err != nil {
		t.Fatalf("GetSubAccountList: %v", err)
	}
	if _, err := client.CreateSubAccount(context.Background(), models.CreateSubAccountRequest{}); err != nil {
		t.Fatalf("CreateSubAccount: %v", err)
	}
	if _, err := client.GetSubAccountAPIKey(context.Background(), "value", &text); err != nil {
		t.Fatalf("GetSubAccountAPIKey: %v", err)
	}
	if _, err := client.CreateSubAccountAPIKey(context.Background(), models.CreateSubAccountAPIKeyRequest{}); err != nil {
		t.Fatalf("CreateSubAccountAPIKey: %v", err)
	}
	if _, err := client.ModifySubAccountAPIKey(context.Background(), models.ModifySubAccountAPIKeyRequest{}); err != nil {
		t.Fatalf("ModifySubAccountAPIKey: %v", err)
	}
	if err := client.DeleteSubAccountAPIKey(context.Background(), models.DeleteSubAccountAPIKeyRequest{}); err != nil {
		t.Fatalf("DeleteSubAccountAPIKey: %v", err)
	}
	if err := client.SetSubAccountTransferOut(context.Background(), models.SetSubAccountTransferOutRequest{}); err != nil {
		t.Fatalf("SetSubAccountTransferOut: %v", err)
	}
	if _, err := client.GetEntrustSubAccountList(context.Background(), &text); err != nil {
		t.Fatalf("GetEntrustSubAccountList: %v", err)
	}
}

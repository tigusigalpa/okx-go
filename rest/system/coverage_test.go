package system

import (
	"context"
	"testing"
)

func TestAllEndpointsAcceptPopulatedOptionalArguments(t *testing.T) {
	text := "value"
	enabled := true
	_ = enabled
	client := NewClient(func(_ context.Context, _ string, _ string, _ map[string]string, _ interface{}) error { return nil })
	if _, err := client.GetStatus(context.Background(), &text); err != nil {
		t.Fatalf("GetStatus: %v", err)
	}
}

package support

import (
	"context"
	"testing"
)

func TestAllEndpointsAcceptPopulatedOptionalArguments(t *testing.T) {
	text := "value"
	enabled := true
	_ = enabled
	client := NewClient(func(_ context.Context, _ string, _ string, _ map[string]string, _ interface{}) error { return nil })
	if _, err := client.GetAnnouncementTypes(context.Background()); err != nil {
		t.Fatalf("GetAnnouncementTypes: %v", err)
	}
	if _, err := client.GetAnnouncements(context.Background(), &text, &text, &text); err != nil {
		t.Fatalf("GetAnnouncements: %v", err)
	}
}

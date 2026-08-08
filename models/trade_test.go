package models

import (
	"encoding/json"
	"testing"
)

func strPtr(s string) *string {
	return &s
}

// TestPlaceOrderRequest_NoAttachAlgoOrds_OmitsField verifies that when
// AttachAlgoOrds is left nil/empty, the "attachAlgoOrds" key is omitted
// entirely, and the rest of the JSON shape is unaffected.
func TestPlaceOrderRequest_NoAttachAlgoOrds_OmitsField(t *testing.T) {
	req := PlaceOrderRequest{
		InstID:  "BTC-USDT",
		TdMode:  "cash",
		Side:    "buy",
		OrdType: "limit",
		Sz:      "0.001",
		Px:      strPtr("20000"),
	}

	b, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, ok := m["attachAlgoOrds"]; ok {
		t.Errorf("attachAlgoOrds should be omitted, got: %s", b)
	}

	want := map[string]interface{}{
		"instId":  "BTC-USDT",
		"tdMode":  "cash",
		"side":    "buy",
		"ordType": "limit",
		"sz":      "0.001",
		"px":      "20000",
	}
	for k, v := range want {
		if m[k] != v {
			t.Errorf("field %q = %v, want %v", k, m[k], v)
		}
	}
	if len(m) != len(want) {
		t.Errorf("unexpected extra fields in %s, got %d fields, want %d", b, len(m), len(want))
	}
}

// TestPlaceOrderRequest_AttachAlgoOrds_TPOnly verifies that a TP-only
// attachment serializes only TP fields (no SL keys present).
func TestPlaceOrderRequest_AttachAlgoOrds_TPOnly(t *testing.T) {
	req := PlaceOrderRequest{
		InstID:  "BTC-USDT-SWAP",
		TdMode:  "isolated",
		Side:    "sell",
		OrdType: "market",
		Sz:      "1",
		PosSide: strPtr("short"),
		AttachAlgoOrds: []AttachAlgoOrderRequest{
			{
				TpTriggerPx:     strPtr("64000"),
				TpOrdPx:         strPtr("-1"),
				TpTriggerPxType: strPtr("last"),
			},
		},
	}

	b, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	algos, ok := m["attachAlgoOrds"].([]interface{})
	if !ok || len(algos) != 1 {
		t.Fatalf("attachAlgoOrds missing or wrong length in %s", b)
	}
	item, ok := algos[0].(map[string]interface{})
	if !ok {
		t.Fatalf("attachAlgoOrds[0] is not an object: %s", b)
	}

	want := map[string]interface{}{
		"tpTriggerPx":     "64000",
		"tpOrdPx":         "-1",
		"tpTriggerPxType": "last",
	}
	for k, v := range want {
		if item[k] != v {
			t.Errorf("field %q = %v, want %v", k, item[k], v)
		}
	}
	if len(item) != len(want) {
		t.Errorf("unexpected extra/missing fields in %v, got %d fields, want %d", item, len(item), len(want))
	}
	for _, slKey := range []string{"slTriggerPx", "slOrdPx", "slTriggerPxType"} {
		if _, present := item[slKey]; present {
			t.Errorf("SL key %q should not be present for TP-only attachment", slKey)
		}
	}
}

// TestPlaceOrderRequest_AttachAlgoOrds_SLOnly verifies that an SL-only
// attachment serializes only SL fields (no TP keys present).
func TestPlaceOrderRequest_AttachAlgoOrds_SLOnly(t *testing.T) {
	req := PlaceOrderRequest{
		InstID:  "BTC-USDT-SWAP",
		TdMode:  "isolated",
		Side:    "sell",
		OrdType: "market",
		Sz:      "1",
		PosSide: strPtr("short"),
		AttachAlgoOrds: []AttachAlgoOrderRequest{
			{
				SlTriggerPx:     strPtr("66000"),
				SlOrdPx:         strPtr("-1"),
				SlTriggerPxType: strPtr("last"),
			},
		},
	}

	b, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	algos, ok := m["attachAlgoOrds"].([]interface{})
	if !ok || len(algos) != 1 {
		t.Fatalf("attachAlgoOrds missing or wrong length in %s", b)
	}
	item, ok := algos[0].(map[string]interface{})
	if !ok {
		t.Fatalf("attachAlgoOrds[0] is not an object: %s", b)
	}

	want := map[string]interface{}{
		"slTriggerPx":     "66000",
		"slOrdPx":         "-1",
		"slTriggerPxType": "last",
	}
	for k, v := range want {
		if item[k] != v {
			t.Errorf("field %q = %v, want %v", k, item[k], v)
		}
	}
	if len(item) != len(want) {
		t.Errorf("unexpected extra/missing fields in %v, got %d fields, want %d", item, len(item), len(want))
	}
	for _, tpKey := range []string{"tpTriggerPx", "tpOrdPx", "tpTriggerPxType"} {
		if _, present := item[tpKey]; present {
			t.Errorf("TP key %q should not be present for SL-only attachment", tpKey)
		}
	}
}

// TestPlaceOrderRequest_AttachAlgoOrds_CombinedTPAndSL verifies that a
// combined TP + SL attachment serializes as a single attachAlgoOrds array
// item containing both TP and SL fields, matching the OKX example payload.
func TestPlaceOrderRequest_AttachAlgoOrds_CombinedTPAndSL(t *testing.T) {
	req := PlaceOrderRequest{
		InstID:  "BTC-USDT-SWAP",
		TdMode:  "isolated",
		Side:    "sell",
		OrdType: "market",
		Sz:      "1",
		PosSide: strPtr("short"),
		AttachAlgoOrds: []AttachAlgoOrderRequest{
			{
				TpTriggerPx:     strPtr("64000"),
				TpOrdPx:         strPtr("-1"),
				TpTriggerPxType: strPtr("last"),
				SlTriggerPx:     strPtr("66000"),
				SlOrdPx:         strPtr("-1"),
				SlTriggerPxType: strPtr("last"),
			},
		},
	}

	b, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	algos, ok := m["attachAlgoOrds"].([]interface{})
	if !ok || len(algos) != 1 {
		t.Fatalf("attachAlgoOrds missing or wrong length in %s", b)
	}
	item, ok := algos[0].(map[string]interface{})
	if !ok {
		t.Fatalf("attachAlgoOrds[0] is not an object: %s", b)
	}

	want := map[string]interface{}{
		"tpTriggerPx":     "64000",
		"tpOrdPx":         "-1",
		"tpTriggerPxType": "last",
		"slTriggerPx":     "66000",
		"slOrdPx":         "-1",
		"slTriggerPxType": "last",
	}
	for k, v := range want {
		if item[k] != v {
			t.Errorf("field %q = %v, want %v", k, item[k], v)
		}
	}
	if len(item) != len(want) {
		t.Errorf("unexpected extra/missing fields in %v, got %d fields, want %d", item, len(item), len(want))
	}

	// Top-level fields must match the documented example exactly.
	topWant := map[string]interface{}{
		"instId":  "BTC-USDT-SWAP",
		"tdMode":  "isolated",
		"side":    "sell",
		"ordType": "market",
		"sz":      "1",
		"posSide": "short",
	}
	for k, v := range topWant {
		if m[k] != v {
			t.Errorf("top-level field %q = %v, want %v", k, m[k], v)
		}
	}
}

// TestAttachAlgoOrderRequest_MarketExecutionSentinel verifies that "-1" for
// TpOrdPx/SlOrdPx (market execution on trigger) is retained as a JSON string,
// not converted to a number or omitted.
func TestAttachAlgoOrderRequest_MarketExecutionSentinel(t *testing.T) {
	algo := AttachAlgoOrderRequest{
		TpOrdPx: strPtr("-1"),
		SlOrdPx: strPtr("-1"),
	}

	b, err := json.Marshal(algo)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if v, ok := m["tpOrdPx"].(string); !ok || v != "-1" {
		t.Errorf("tpOrdPx = %v (%T), want string \"-1\"", m["tpOrdPx"], m["tpOrdPx"])
	}
	if v, ok := m["slOrdPx"].(string); !ok || v != "-1" {
		t.Errorf("slOrdPx = %v (%T), want string \"-1\"", m["slOrdPx"], m["slOrdPx"])
	}
}

// TestAttachAlgoOrderRequest_NilFieldsOmitted verifies that unset optional
// pointer fields are absent from the JSON output rather than serialized as
// empty strings.
func TestAttachAlgoOrderRequest_NilFieldsOmitted(t *testing.T) {
	algo := AttachAlgoOrderRequest{
		TpTriggerPx: strPtr("64000"),
	}

	b, err := json.Marshal(algo)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(m) != 1 {
		t.Errorf("expected only tpTriggerPx to be present, got %s", b)
	}
	if _, present := m["tpTriggerPx"]; !present {
		t.Errorf("expected tpTriggerPx to be present in %s", b)
	}
	for _, absentKey := range []string{
		"attachAlgoClOrdId", "tpOrdPx", "tpTriggerPxType",
		"slOrdPx", "slTriggerPx", "slTriggerPxType",
		"sz", "amendPxOnTriggerType",
	} {
		if _, present := m[absentKey]; present {
			t.Errorf("expected %q to be omitted (nil), but it was present in %s", absentKey, b)
		}
	}
}

// TestPlaceOrderRequest_LegacyShapePreserved verifies that a normal order
// without TP/SL (and without AttachAlgoOrds) preserves its exact existing
// JSON shape, i.e. this change does not alter serialization for existing
// callers.
func TestPlaceOrderRequest_LegacyShapePreserved(t *testing.T) {
	req := PlaceOrderRequest{
		InstID:      "BTC-USDT",
		TdMode:      "cash",
		Side:        "buy",
		OrdType:     "limit",
		Sz:          "0.001",
		Px:          strPtr("20000"),
		TpTriggerPx: strPtr("21000"),
		TpOrdPx:     strPtr("21000"),
		SlTriggerPx: strPtr("19000"),
		SlOrdPx:     strPtr("19000"),
	}

	b, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, present := m["attachAlgoOrds"]; present {
		t.Errorf("attachAlgoOrds should be omitted when unset, got %s", b)
	}

	want := map[string]interface{}{
		"instId":      "BTC-USDT",
		"tdMode":      "cash",
		"side":        "buy",
		"ordType":     "limit",
		"sz":          "0.001",
		"px":          "20000",
		"tpTriggerPx": "21000",
		"tpOrdPx":     "21000",
		"slTriggerPx": "19000",
		"slOrdPx":     "19000",
	}
	for k, v := range want {
		if m[k] != v {
			t.Errorf("field %q = %v, want %v", k, m[k], v)
		}
	}
	if len(m) != len(want) {
		t.Errorf("unexpected extra/missing fields in %s, got %d fields, want %d", b, len(m), len(want))
	}
}

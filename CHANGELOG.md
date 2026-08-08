# Changelog

All notable changes to this project are documented in this file.

## [v1.0.3] - 2026-08-08

### Added
- **Attached TP/SL via `attachAlgoOrds`**: added `models.AttachAlgoOrderRequest`, a request-only DTO (all fields pointers) for `models.PlaceOrderRequest.AttachAlgoOrds`. OKX now rejects legacy inline `tpTriggerPx`/`slTriggerPx` fields on `POST /api/v5/trade/order` for some scenarios with `sCode=54070` ("use the attachAlgoOrds array to place orders via Open API"); `AttachAlgoOrds` is the current Open API mechanism for attaching TP/SL to the main order rather than creating a separate, unlinked algo order.
- `models.PlaceOrderRequest` gained `AttachAlgoOrds []AttachAlgoOrderRequest` (`json:"attachAlgoOrds,omitempty"`). Legacy inline TP/SL fields on `PlaceOrderRequest` are unchanged and kept for backward compatibility.
- `AttachAlgoOrderRequest` is intentionally distinct from the existing response DTO `models.AttachAlgoOrder` (non-optional string fields, used in `Order.AttachAlgoOrds`), avoiding unwanted empty-string JSON fields on requests.
- No changes to the request/signing pipeline: `Trade.PlaceOrder` continues to use the SDK's existing authenticated `do` pipeline unchanged.
- README documents the `attachAlgoOrds` usage, the `-1` market-execution sentinel for `tpOrdPx`/`slOrdPx`, and the `last`/`index`/`mark` trigger price types.

### Tests
- Added `models/trade_test.go` covering: `attachAlgoOrds` omitted when empty, TP-only serialization, SL-only serialization, combined TP+SL as a single array item, `"-1"` retained as a string sentinel, nil optional fields omitted (not empty strings), and unchanged JSON shape for orders without TP/SL.

## [v1.0.2] - 2026-08-08

### Fixed
- **Error handling**: preserved detailed OKX API errors when the response envelope has a non-zero `code`. Previously, error codes mapped to a sentinel error (e.g. `code="1"` → `ErrInternalServer`) were wrapped via `fmt.Errorf("%w: %s", sentinelErr, okxErr.Error())`, which stringified the underlying `*OKXError` and made it unreachable via `errors.As()`. Callers could not access `OKXError.Raw` to inspect `data[].sCode` / `data[].sMsg` for the real failure reason (e.g. `51008: Order failed. Insufficient balance` inside a top-level `code="1": All operations failed` response).
- Fixed in both `Client.do` (authenticated requests) and `Client.doPublic` (public requests) in `client.go` by wrapping with `fmt.Errorf("%w: %w", sentinelErr, okxErr)`, so `errors.Is(err, ErrInternalServer)` and `errors.As(err, &apiErr)` (`*OKXError`) both work on the same error chain.
- No changes to public API, sentinel errors, or successful-response handling. Unknown error codes continue to return the raw `*OKXError` unchanged.

### Tests
- Added `TestDo_PreservesDetailedOKXError` in `client_test.go`, using `httptest.NewServer` to emulate an OKX `code="1"` response with `data[].sCode="51008"` / `sMsg="Order failed. Insufficient balance"`. Verifies `errors.Is`, `errors.As`, `OKXError.Code`, `OKXError.Message`, and `OKXError.Raw`.

## [v1.0.1] - 2026-06-02

### Fixed
- `fix(market): parse OKX candles as positional arrays` — OKX returns each candle as a positional JSON array (e.g. `["ts","o","h","l","c","vol",...]`), but `models.Candle` was decoded as a keyed object, causing `cannot unmarshal array into Go value of type models.Candle`. Added a custom `UnmarshalJSON` on `Candle` that maps by index and handles both the 9-field market candles and the 6-field index/mark-price candles. Fixes all market candle endpoints.

## [v1.0.0] - 2026-06-02

- Initial release.

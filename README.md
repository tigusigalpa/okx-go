# OKX Golang SDK

![OKX Golang client](https://i.postimg.cc/rpvZ9816/okx-golang-github-hero.jpg)

[![CI](https://github.com/tigusigalpa/okx-go/actions/workflows/ci.yml/badge.svg)](https://github.com/tigusigalpa/okx-go/actions/workflows/ci.yml)
[![Tests](https://img.shields.io/badge/tests-go%20test%20--race-brightgreen)](https://github.com/tigusigalpa/okx-go/actions/workflows/ci.yml)
[![Go vet](https://img.shields.io/badge/code%20analysis-go%20vet-brightgreen)](https://github.com/tigusigalpa/okx-go/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![CodeQL](https://github.com/tigusigalpa/okx-go/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/tigusigalpa/okx-go/actions/workflows/codeql.yml)
[![codecov](https://codecov.io/gh/tigusigalpa/okx-go/graph/badge.svg)](https://codecov.io/gh/tigusigalpa/okx-go)

Go client for the [OKX v5 API](https://www.okx.com/docs-v5/en/). Covers 335 REST endpoints and 53 WebSocket channels.

**Package:** [pkg.go.dev/github.com/tigusigalpa/okx-go](https://pkg.go.dev/github.com/tigusigalpa/okx-go)

> 📖 **[Full documentation available on Wiki](https://github.com/tigusigalpa/okx-go/wiki)**

## Install

```bash
go get github.com/tigusigalpa/okx-go
```

> [!IMPORTANT]
> **Breaking change in v1.1.0:** `OKXError` has been renamed to `Error` to follow Go naming conventions. Update type assertions and `errors.As` targets from `*okx.OKXError` to `*okx.Error` when upgrading.

## What's inside

- 335 REST endpoints across 16 categories
- 53 WebSocket channels (public, private, business)
- Demo trading mode (`x-simulated-trading: 1`)
- `context.Context` everywhere
- Typed request/response structs
- WebSocket reconnect with exponential backoff
- Goroutine-safe
- Dependencies: stdlib + `gorilla/websocket`

## Quick start

Runnable examples:

```bash
go run ./examples/rest
go run ./examples/websocket/public
go run ./examples/websocket/business
go run ./examples/websocket/private
```

### REST

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/tigusigalpa/okx-go"
)

func main() {
    client := okx.NewRestClient(
        "your-api-key",
        "your-secret-key",
        "your-passphrase",
        okx.WithDemoTrading(),
    )

    ctx := context.Background()

    balances, err := client.Account.GetBalance(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    for _, b := range balances {
        fmt.Printf("Total Equity: %s\n", *b.TotalEq)
    }

    // Place a limit order
    px := "30000"
    order := models.PlaceOrderRequest{
        InstID:  "BTC-USDT",
        TdMode:  "cash",
        Side:    "buy",
        OrdType: "limit",
        Px:      &px,
        Sz:      "0.01",
    }

    result, err := client.Trade.PlaceOrder(ctx, order)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Order ID: %s\n", result[0].OrdID)
}
```

### WebSocket (public, typed)

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

ws := okx.NewWSClient("", "", "", okx.WSPublicURL, okx.WithWSSubscriptionBuffer(256))
defer ws.Close()

if err := ws.Connect(ctx); err != nil {
    log.Fatal(err)
}

stream, err := ws.Subscribe(ctx, "tickers", map[string]interface{}{
    "instId": "BTC-USDT",
})
if err != nil {
    log.Fatal(err)
}

for raw := range stream {
    message, err := models.DecodeWSMessage[models.Ticker](raw)
    if err != nil {
        log.Printf("decode ticker: %v", err)
        continue
    }
    for _, ticker := range message.Data {
        fmt.Printf("%s last=%s\n", ticker.InstID, ticker.Last)
    }
}
```

### WebSocket (private)

```go
ws := okx.NewWSClient(
    "your-api-key",
    "your-secret-key",
    "your-passphrase",
    okx.WSPrivateURL,
)

ctx := context.Background()
if err := ws.Connect(ctx); err != nil {
    log.Fatal(err)
}
defer ws.Close()

if err := ws.Login(ctx); err != nil {
    log.Fatal(err)
}

ch, err := ws.Subscribe(ctx, "account", map[string]interface{}{
    "ccy": "BTC",
})
if err != nil {
    log.Fatal(err)
}

for msg := range ch {
    fmt.Printf("%s\n", msg)
}
```

`Login` only returns successfully after OKX acknowledges the request. Use
`ws.IsAuthenticated()` when an application needs to inspect the current connection state.

### WebSocket regions, demo, and protocol support

OKX will discontinue the legacy WebSocket port 8443 on **2026-10-31**. This SDK uses
the default secure WebSocket port (443) for every endpoint below.

| Region | REST configuration | Production JSON services | Demo JSON services |
|--------|--------------------|--------------------------|--------------------|
| Global | `DefaultBaseURL` | `wss://ws.okx.com/ws/v5/{public,private,business}` | `wss://wspap.okx.com/ws/v5/{public,private,business}` |
| EEA | `EEABaseURL` | `wss://wseea.okx.com/ws/v5/{public,private,business}` | `wss://wseeapap.okx.com/ws/v5/{public,private,business}` |
| United States | `USBaseURL` | `wss://wsus.okx.com/ws/v5/{public,private,business}` | `wss://wsuspap.okx.com/ws/v5/{public,private,business}` |
| Türkiye | `TRBaseURL` | Global endpoints | Global demo endpoints |

Use `WSPublicURL`, `WSPrivateURL`, and `WSBusinessURL` for Global, or select an endpoint
without hard-coding hosts:

```go
endpoint, err := okx.WebSocketURL(okx.WSRegionEEA, okx.WSServiceBusiness, true)
if err != nil {
    log.Fatal(err)
}
ws := okx.NewWSClient(apiKey, secret, passphrase, endpoint)
```

The exported EEA and US constants follow the same pattern: `WSEEAPublicURL`,
`WSEEAPrivateURL`, `WSEEABusinessURL`, `WSUSPublicURL`, `WSUSPrivateURL`, and
`WSUSBusinessURL`; their `WSDemo...` counterparts select demo hosts. `WithWSDemo()` also
maps a supported production endpoint to its matching demo endpoint. Türkiye currently uses
the Global WebSocket deployment; no distinct Türkiye WebSocket hostname is invented by this SDK.

`WSPublicSBEURL` and `WSDemoPublicSBEURL` remain available for source compatibility, but this
client implements JSON WebSocket messages only and does not decode SBE binary market data.

### WebSocket lifecycle, errors, and backpressure

`Connect` starts one activity-aware heartbeat. It sends a text `ping` only while the connection
is idle and reconnects when no activity arrives before the pong timeout. A disconnect triggers
one reconnect loop with bounded exponential backoff; successful reconnections restore a confirmed
login and active subscriptions. `Close` is idempotent and immediately interrupts a pending backoff.

`Subscribe` preserves the original raw-message API. A successful call means the request was
written; subscribe acknowledgements, server errors, reconnects, and buffer overflows are exposed
through `Events()` (the channel closes with `Close`):

```go
for event := range ws.Events() {
    if event.Err != nil {
        log.Printf("WebSocket %s: %v", event.Type, event.Err)
    }
}
```

Each subscription has a bounded message channel (100 messages by default). Make overload explicit
with `WithWSSubscriptionBuffer`, and receive every local overflow synchronously with
`WithWSDropHandler`:

```go
ws := okx.NewWSClient("", "", "", okx.WSPublicURL,
    okx.WithWSSubscriptionBuffer(512),
    okx.WithWSDropHandler(func(drop okx.WSDrop) {
        log.Printf("dropped %d messages on %s", drop.Count, drop.Channel)
    }),
)
```

### WebSocket candles and order books

Decode push payloads with `models.DecodeWSMessage[T]`. It keeps financial values as strings and
supports `Ticker`, `Trade`, `Candle`, `OpenInterest`, `FundingRate`, `OrderBook`,
`LiquidationOrder`, and `MarkPrice`. Market candles use the 9-field OKX array; index and mark-price
candles use the 6-field form. `Candle.IsConfirmed()` distinguishes a closed candle from an update.

Candle channels are served through the Business endpoint:

```go
ws := okx.NewWSClient("", "", "", okx.WSBusinessURL)
// Connect, then subscribe to "candle1m" with {"instId": "BTC-USDT"}.
```

For incremental `books` data, use `seqId` and `prevSeqId` rather than `checksum`. OKX has
deprecated checksum validation for incremental books streams; a zero checksum is expected. The
helper accepts snapshots (`prevSeqId == -1`) and maintenance resets while detecting a real gap:

```go
if err := okx.ValidateOrderBookSequence(previous, next); err != nil {
    // resubscribe or request a fresh snapshot
}
```

`books5` is a snapshot-style channel and should not be treated as an incremental checksum feed.

## Attached TP/SL (attachAlgoOrds)

OKX rejects the legacy inline `tpTriggerPx`/`slTriggerPx` fields on `POST /api/v5/trade/order`
for some instruments/scenarios with `sCode=54070` ("use the attachAlgoOrds array to place
orders via Open API"). Use `models.PlaceOrderRequest.AttachAlgoOrds` — the current Open API
mechanism for attaching TP/SL to the main order (as opposed to a separate, unlinked algo order
via `PlaceAlgoOrder`):

```go
tpTriggerPx := "64000"
tpOrdPx := "-1" // "-1" = execute TP as a market order once triggered
slTriggerPx := "66000"
slOrdPx := "-1" // "-1" = execute SL as a market order once triggered
triggerPxType := "last" // "last", "index", or "mark"

order := models.PlaceOrderRequest{
    InstID:  "BTC-USDT-SWAP",
    TdMode:  "isolated",
    Side:    "sell",
    OrdType: "market",
    Sz:      "1",
    PosSide: strPtr("short"),
    AttachAlgoOrds: []models.AttachAlgoOrderRequest{
        {
            TpTriggerPx:     &tpTriggerPx,
            TpOrdPx:         &tpOrdPx,
            TpTriggerPxType: &triggerPxType,
            SlTriggerPx:     &slTriggerPx,
            SlOrdPx:         &slOrdPx,
            SlTriggerPxType: &triggerPxType,
        },
    },
}

result, err := client.Trade.PlaceOrder(ctx, order)
```

Notes:

- `AttachAlgoOrderRequest` is a request-only DTO (all fields are pointers, so unset fields are
  omitted from the JSON payload). It is distinct from `models.AttachAlgoOrder`, which is the
  response shape returned inside `Order.AttachAlgoOrds` and has non-optional string fields.
- Legacy inline `TpTriggerPx`/`TpOrdPx`/`SlTriggerPx`/`SlOrdPx` fields on `PlaceOrderRequest`
  are kept for backward compatibility, but prefer `AttachAlgoOrds` going forward.

## Options

| Option                  | Description           | Default                      |
|-------------------------|-----------------------|------------------------------|
| `WithHTTPClient(c)`     | Custom `*http.Client` | `&http.Client{Timeout: 30s}` |
| `WithBaseURL(url)`      | Override base URL     | `https://openapi.okx.com`    |
| `WithDemoTrading()`     | Demo mode             | off                          |
| `WithTimeout(d)`        | Request timeout       | `30s`                        |
| `WithRateLimiter(true)` | Rate limiter          | off                          |
| `WithLogger(l)`         | Custom `Logger`       | no-op                        |

### Regional REST endpoints

`DefaultBaseURL` uses `https://openapi.okx.com`, the recommended endpoint for
OKX Global accounts. Regional accounts must use the matching REST domain:

```go
// OKX United States
client := okx.NewRestClient(key, secret, passphrase, okx.WithBaseURL(okx.USBaseURL))

// OKX European Economic Area
client := okx.NewRestClient(key, secret, passphrase, okx.WithBaseURL(okx.EEABaseURL))

// OKX Türkiye
client := okx.NewRestClient(key, secret, passphrase, okx.WithBaseURL(okx.TRBaseURL))
```

WebSocket hosts are region-specific. See [WebSocket regions, demo, and protocol support](#websocket-regions-demo-and-protocol-support)
for the exact Global, EEA, US, and Türkiye behavior.

## REST endpoints

| Category           | Count   | Docs                                                                            |
|--------------------|---------|---------------------------------------------------------------------------------|
| Account            | 53      | [link](https://www.okx.com/docs-v5/en/#trading-account-rest-api)                |
| Trade              | 32      | [link](https://www.okx.com/docs-v5/en/#order-book-trading-trade-rest-api)       |
| Market Data        | 24      | [link](https://www.okx.com/docs-v5/en/#order-book-trading-market-data-rest-api) |
| Public Data        | 24      | [link](https://www.okx.com/docs-v5/en/#public-data-rest-api)                    |
| Asset              | 26      | [link](https://www.okx.com/docs-v5/en/#funding-account-rest-api)                |
| Sub-account        | 8       | [link](https://www.okx.com/docs-v5/en/#sub-account-rest-api)                    |
| Trading Bot        | 44      | [link](https://www.okx.com/docs-v5/en/#trading-bot-grid-trading-rest-api)       |
| Copy Trading       | 26      | [link](https://www.okx.com/docs-v5/en/#copy-trading-rest-api)                   |
| Block Trading      | 20      | [link](https://www.okx.com/docs-v5/en/#block-trading-rest-api)                  |
| Spread Trading     | 13      | [link](https://www.okx.com/docs-v5/en/#spread-trading-rest-api)                 |
| Financial Products | 33      | [link](https://www.okx.com/docs-v5/en/#financial-product-rest-api)              |
| Fiat               | 13      | [link](https://www.okx.com/docs-v5/en/#fiat-rest-api)                           |
| Trading Statistics | 15      | [link](https://www.okx.com/docs-v5/en/#trading-statistics-rest-api)             |
| System             | 1       | [link](https://www.okx.com/docs-v5/en/#status-rest-api)                         |
| Announcement       | 2       | [link](https://www.okx.com/docs-v5/en/#announcement-rest-api)                   |
| Affiliate          | 1       | [link](https://www.okx.com/docs-v5/en/#affiliate-rest-api)                      |
| **Total**          | **335** |                                                                                 |

## WebSocket channels

**Public (31):**
`tickers`, `candle1D`, `candle1H`, `candle30m`, `trades`, `books`, `books5`, `bbo-tbt`, `opt-summary`,
`estimated-price`, `mark-price`, `mark-price-candle1D`, `price-limit`, `open-interest`, `funding-rate`,
`index-candle30m`, `index-tickers`, `status`, `public-struc-block-trades`, `block-tickers`, `block-trades`,
`liquidation-orders`, `sprd-tickers`, `sprd-books5`, `sprd-books-l2-tbt`, `sprd-public-trades`, `sprd-candle1D`,
`economic-calendar`, `call-auction-details`, `instruments`, `trades-all`

**Private (22):**
`account`, `positions`, `balance_and_position`, `orders`, `orders-algo`, `algo-advance`, `liquidation-warning`,
`account-greeks`, `rfqs`, `quotes`, `sprd-orders`, `sprd-trades`, `adl-warning`, `fills`, `deposit-info`,
`withdrawal-info`, `grid-orders-spot`, `grid-orders-contract`, `grid-positions`, `grid-sub-orders`,
`algo-recurring-buy`, `copytrading-lead-notification`

## Demo trading

REST — pass `okx.WithDemoTrading()`:

```go
client := okx.NewRestClient(apiKey, secret, passphrase, okx.WithDemoTrading())
```

WebSocket — use demo URLs or `WithWSDemo()`:

- `okx.WSDemoPublicURL`
- `okx.WSDemoPrivateURL`
- `okx.WSDemoBusinessURL`
- `okx.WSDemoEEAPublicURL`, `okx.WSDemoEEAPrivateURL`, `okx.WSDemoEEABusinessURL`
- `okx.WSDemoUSPublicURL`, `okx.WSDemoUSPrivateURL`, `okx.WSDemoUSBusinessURL`

## Errors

```go
balances, err := client.Account.GetBalance(ctx, nil)
if err != nil {
    if errors.Is(err, okx.ErrUnauthorized) {
        // bad credentials
    } else if errors.Is(err, okx.ErrRateLimited) {
        // slow down
    } else if okxErr, ok := err.(*okx.Error); ok {
        fmt.Printf("code=%s msg=%s\n", okxErr.Code, okxErr.Message)
    }
}
```

## Pagination

OKX uses cursor-based pagination (`before`/`after`). There's a generic `Paginator[T]` helper:

```go
paginator := models.NewPaginator(func(after string) ([]models.Order, string, error) {
    orders, err := client.Trade.GetOrdersHistory(ctx, "SPOT", nil, nil, nil, nil, nil, nil, &after, nil, nil, nil, nil)
    if err != nil {
        return nil, "", err
    }
    var next string
    if len(orders) > 0 {
        next = orders[len(orders)-1].OrdID
    }
    return orders, next, nil
})

allOrders, err := paginator.All()
```

## Tests

```bash
# unit
go test ./...

# integration (demo env)
OKX_API_KEY=... OKX_SECRET_KEY=... OKX_PASSPHRASE=... go test -tags=integration ./...

# public WebSocket integration (no API credentials required)
go test -tags=integration -run TestIntegration_WebSocket_PublicChannel ./...
```

## Contributing

Fork, branch, PR. Make sure `go test ./...` passes and new code has tests. See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

MIT. See [LICENSE](LICENSE).

## Author

Igor Sazonov — [@tigusigalpa](https://github.com/tigusigalpa) — sovletig@gmail.com

## Links

- [OKX API docs](https://www.okx.com/docs-v5/en/)
- [Issues](https://github.com/tigusigalpa/okx-go/issues)

Not affiliated with OKX. Test on demo before going live. Golang library.

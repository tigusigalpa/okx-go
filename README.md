# okx-go

![OKX Golang client](https://i.postimg.cc/Rh3kyD66/561657912-aa6baa0a-566e-4ac3-b1df-a495edaf1328.jpg)

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Go client for the [OKX v5 API](https://www.okx.com/docs-v5/en/). Covers 335 REST endpoints and 53 WebSocket channels.

**Package:** [pkg.go.dev/github.com/tigusigalpa/okx-go](https://pkg.go.dev/github.com/tigusigalpa/okx-go)

> 📖 **[Full documentation available on Wiki](https://github.com/tigusigalpa/okx-go/wiki)**

## Install

```bash
go get github.com/tigusigalpa/okx-go
```

## What's inside

- 335 REST endpoints across 16 categories
- 53 WebSocket channels (public, private, business)
- Demo trading mode (`x-simulated-trading: 1`)
- `context.Context` everywhere
- Typed request/response structs with generics
- Rate limiter (token bucket, configurable per category)
- WebSocket reconnect with exponential backoff
- Goroutine-safe
- Dependencies: stdlib + `gorilla/websocket`

## Quick start

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

### WebSocket (public)

```go
ws := okx.NewWSClient("", "", "", okx.WSPublicURL)

ctx := context.Background()
if err := ws.Connect(ctx); err != nil {
    log.Fatal(err)
}
defer ws.Close()

ch, err := ws.Subscribe(ctx, "tickers", map[string]interface{}{
    "instId": "BTC-USDT",
})
if err != nil {
    log.Fatal(err)
}

for msg := range ch {
    fmt.Printf("%s\n", msg)
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

## Options

| Option                  | Description           | Default                      |
|-------------------------|-----------------------|------------------------------|
| `WithHTTPClient(c)`     | Custom `*http.Client` | `&http.Client{Timeout: 30s}` |
| `WithBaseURL(url)`      | Override base URL     | `https://www.okx.com`        |
| `WithDemoTrading()`     | Demo mode             | off                          |
| `WithTimeout(d)`        | Request timeout       | `30s`                        |
| `WithRateLimiter(true)` | Rate limiter          | off                          |
| `WithLogger(l)`         | Custom `Logger`       | no-op                        |

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

WebSocket — use demo URLs:

- `okx.WSDemoPublicURL`
- `okx.WSDemoPrivateURL`
- `okx.WSDemoBusinessURL`

## Errors

```go
balances, err := client.Account.GetBalance(ctx, nil)
if err != nil {
    if errors.Is(err, okx.ErrUnauthorized) {
        // bad credentials
    } else if errors.Is(err, okx.ErrRateLimited) {
        // slow down
    } else if okxErr, ok := err.(*okx.OKXError); ok {
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

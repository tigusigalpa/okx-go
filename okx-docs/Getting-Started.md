# Getting Started

## Create API credentials

1. Log in to your OKX account
2. Go to API management
3. Create a new API key
4. Save your API key, secret key, and passphrase securely
5. Set appropriate permissions (read, trade, withdraw)

For testing, create demo trading credentials at https://www.okx.com/demo-trading

## Basic REST example

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
    )

    ctx := context.Background()

    // Get account balance
    balances, err := client.Account.GetBalance(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    for _, b := range balances {
        fmt.Printf("Total Equity: %s\n", *b.TotalEq)
    }
}
```

## Basic WebSocket example

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/tigusigalpa/okx-go"
)

func main() {
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
        fmt.Printf("Ticker: %s\n", msg)
    }
}
```

## Demo trading mode

For testing without real funds:

```go
client := okx.NewRestClient(
    "demo-api-key",
    "demo-secret-key",
    "demo-passphrase",
    okx.WithDemoTrading(),
)
```

WebSocket demo:

```go
ws := okx.NewWSClient(
    "demo-api-key",
    "demo-secret-key",
    "demo-passphrase",
    okx.WSDemoPrivateURL,
)
```

## Common patterns

### Using context for timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

balances, err := client.Account.GetBalance(ctx, nil)
```

### Error handling

```go
balances, err := client.Account.GetBalance(ctx, nil)
if err != nil {
    if errors.Is(err, okx.ErrUnauthorized) {
        log.Fatal("Invalid credentials")
    }
    log.Fatal(err)
}
```

### Optional parameters

Many endpoints have optional parameters. Use pointers:

```go
ccy := "BTC"
balances, err := client.Account.GetBalance(ctx, &ccy)
```

## Next Steps

- [REST API Guide](REST-API-Guide)
- [WebSocket Guide](WebSocket-Guide)
- [Examples](Examples)

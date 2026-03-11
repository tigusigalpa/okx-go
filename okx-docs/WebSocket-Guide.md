# WebSocket Guide

## Overview

The WebSocket client provides real-time data streams for 53 channels.

## Connection types

### Public channels

No authentication required.

```go
ws := okx.NewWSClient("", "", "", okx.WSPublicURL)
```

### Private channels

Requires authentication.

```go
ws := okx.NewWSClient(
    "api-key",
    "secret-key",
    "passphrase",
    okx.WSPrivateURL,
)
```

### Business channels

For business-specific data.

```go
ws := okx.NewWSClient(
    "api-key",
    "secret-key",
    "passphrase",
    okx.WSBusinessURL,
)
```

## Basic usage

### Connect

```go
ctx := context.Background()
if err := ws.Connect(ctx); err != nil {
    log.Fatal(err)
}
defer ws.Close()
```

### Authenticate (private channels only)

```go
if err := ws.Login(ctx); err != nil {
    log.Fatal(err)
}
```

### Subscribe

```go
ch, err := ws.Subscribe(ctx, "tickers", map[string]interface{}{
    "instId": "BTC-USDT",
})
if err != nil {
    log.Fatal(err)
}

for msg := range ch {
    fmt.Printf("Message: %s\n", msg)
}
```

### Unsubscribe

```go
err := ws.Unsubscribe("tickers", map[string]interface{}{
    "instId": "BTC-USDT",
})
```

## Public channels (31)

### Tickers

```go
ch, _ := ws.Subscribe(ctx, "tickers", map[string]interface{}{
    "instId": "BTC-USDT",
})
```

### Order book

```go
ch, _ := ws.Subscribe(ctx, "books", map[string]interface{}{
    "instId": "BTC-USDT",
})
```

### Trades

```go
ch, _ := ws.Subscribe(ctx, "trades", map[string]interface{}{
    "instId": "BTC-USDT",
})
```

### Candles

```go
ch, _ := ws.Subscribe(ctx, "candle1m", map[string]interface{}{
    "instId": "BTC-USDT",
})
```

### Mark price

```go
ch, _ := ws.Subscribe(ctx, "mark-price", map[string]interface{}{
    "instId": "BTC-USDT-SWAP",
})
```

### Funding rate

```go
ch, _ := ws.Subscribe(ctx, "funding-rate", map[string]interface{}{
    "instId": "BTC-USDT-SWAP",
})
```

### Index tickers

```go
ch, _ := ws.Subscribe(ctx, "index-tickers", map[string]interface{}{
    "instId": "BTC-USDT",
})
```

### Liquidation orders

```go
ch, _ := ws.Subscribe(ctx, "liquidation-orders", map[string]interface{}{
    "instType": "SWAP",
})
```

## Private channels (22)

### Account

```go
ch, _ := ws.Subscribe(ctx, "account", map[string]interface{}{})
```

### Positions

```go
ch, _ := ws.Subscribe(ctx, "positions", map[string]interface{}{
    "instType": "SWAP",
})
```

### Orders

```go
ch, _ := ws.Subscribe(ctx, "orders", map[string]interface{}{
    "instType": "SPOT",
})
```

### Algo orders

```go
ch, _ := ws.Subscribe(ctx, "orders-algo", map[string]interface{}{
    "instType": "SWAP",
})
```

### Fills

```go
ch, _ := ws.Subscribe(ctx, "fills", map[string]interface{}{})
```

### Balance and position

```go
ch, _ := ws.Subscribe(ctx, "balance_and_position", map[string]interface{}{})
```

### Deposit info

```go
ch, _ := ws.Subscribe(ctx, "deposit-info", map[string]interface{}{})
```

### Withdrawal info

```go
ch, _ := ws.Subscribe(ctx, "withdrawal-info", map[string]interface{}{})
```

## Advanced usage

### Multiple subscriptions

```go
instruments := []string{"BTC-USDT", "ETH-USDT", "SOL-USDT"}
channels := make(map[string]<-chan []byte)

for _, inst := range instruments {
    ch, err := ws.Subscribe(ctx, "tickers", map[string]interface{}{
        "instId": inst,
    })
    if err != nil {
        log.Fatal(err)
    }
    channels[inst] = ch
}

for inst, ch := range channels {
    go func(instrument string, channel <-chan []byte) {
        for msg := range channel {
            fmt.Printf("[%s] %s\n", instrument, msg)
        }
    }(inst, ch)
}
```

### Parse messages

```go
type TickerData struct {
    InstID  string `json:"instId"`
    Last    string `json:"last"`
    BidPx   string `json:"bidPx"`
    AskPx   string `json:"askPx"`
    Vol24h  string `json:"vol24h"`
}

type TickerMessage struct {
    Arg struct {
        Channel string `json:"channel"`
        InstID  string `json:"instId"`
    } `json:"arg"`
    Data []TickerData `json:"data"`
}

ch, _ := ws.Subscribe(ctx, "tickers", map[string]interface{}{
    "instId": "BTC-USDT",
})

for msg := range ch {
    var ticker TickerMessage
    if err := json.Unmarshal(msg, &ticker); err != nil {
        log.Printf("Parse error: %v", err)
        continue
    }
    
    for _, data := range ticker.Data {
        fmt.Printf("Price: %s, Volume: %s\n", data.Last, data.Vol24h)
    }
}
```

### Graceful shutdown

```go
sigCh := make(chan os.Signal, 1)
signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

go func() {
    <-sigCh
    fmt.Println("Shutting down...")
    ws.Close()
}()

ch, _ := ws.Subscribe(ctx, "tickers", map[string]interface{}{
    "instId": "BTC-USDT",
})

for msg := range ch {
    fmt.Printf("%s\n", msg)
}
```

### Reconnection

The client automatically reconnects with exponential backoff. You don't need to handle reconnection manually.

### Heartbeat

The client automatically sends ping/pong messages every 25 seconds. No manual intervention needed.

## Demo trading

```go
ws := okx.NewWSClient(
    "demo-key",
    "demo-secret",
    "demo-pass",
    okx.WSDemoPrivateURL,
)
```

## Troubleshooting

### Connection fails

- Check network connectivity
- Verify URL is correct
- Check firewall settings

### Authentication fails

- Verify credentials are correct
- Check timestamp is UTC
- Ensure signature is valid

### No messages received

- Verify subscription was successful
- Check channel name is correct
- Ensure instrument exists

### Channel buffer full

The default buffer is 100 messages. If you're not reading fast enough:

```go
// Process messages in goroutine
go func() {
    for msg := range ch {
        // Process asynchronously
        go handleMessage(msg)
    }
}()
```

## Next Steps

- [REST API Guide](REST-API-Guide)
- [Error Handling](Error-Handling)
- [Examples](Examples)

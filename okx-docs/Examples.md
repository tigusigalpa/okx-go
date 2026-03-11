# Examples

## Basic examples

### Get account balance

```go
client := okx.NewRestClient(apiKey, secret, pass)
ctx := context.Background()

balances, err := client.Account.GetBalance(ctx, nil)
if err != nil {
    log.Fatal(err)
}

for _, b := range balances {
    fmt.Printf("Total: %s\n", *b.TotalEq)
    for _, d := range b.Details {
        fmt.Printf("  %s: %s\n", d.Ccy, *d.Eq)
    }
}
```

### Place market order

```go
order := models.PlaceOrderRequest{
    InstID:  "BTC-USDT",
    TdMode:  "cash",
    Side:    "buy",
    OrdType: "market",
    Sz:      "0.001",
}

result, err := client.Trade.PlaceOrder(ctx, order)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Order ID: %s\n", result[0].OrdID)
```

### Place limit order

```go
px := "30000"
order := models.PlaceOrderRequest{
    InstID:  "BTC-USDT",
    TdMode:  "cash",
    Side:    "buy",
    OrdType: "limit",
    Px:      &px,
    Sz:      "0.001",
}

result, err := client.Trade.PlaceOrder(ctx, order)
```

### Cancel order

```go
cancel := models.CancelOrderRequest{
    InstID: "BTC-USDT",
    OrdID:  &orderID,
}

_, err := client.Trade.CancelOrder(ctx, cancel)
```

### Get ticker

```go
ticker, err := client.Market.GetTicker(ctx, "BTC-USDT")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Last: %s\n", ticker[0].Last)
fmt.Printf("Bid: %s\n", ticker[0].BidPx)
fmt.Printf("Ask: %s\n", ticker[0].AskPx)
```

### Get candles

```go
bar := "1H"
limit := "100"

candles, err := client.Market.GetCandles(ctx, "BTC-USDT", &bar, nil, nil, &limit)
if err != nil {
    log.Fatal(err)
}

for _, c := range candles {
    fmt.Printf("Time: %s, O: %s, H: %s, L: %s, C: %s\n",
        c.TS, c.O, c.H, c.L, c.C)
}
```

## WebSocket examples

### Subscribe to ticker

```go
ws := okx.NewWSClient("", "", "", okx.WSPublicURL)

ctx := context.Background()
ws.Connect(ctx)
defer ws.Close()

ch, _ := ws.Subscribe(ctx, "tickers", map[string]interface{}{
    "instId": "BTC-USDT",
})

for msg := range ch {
    var data map[string]interface{}
    json.Unmarshal(msg, &data)
    fmt.Printf("%+v\n", data)
}
```

### Subscribe to order book

```go
ch, _ := ws.Subscribe(ctx, "books", map[string]interface{}{
    "instId": "BTC-USDT",
})

for msg := range ch {
    fmt.Printf("Order book: %s\n", msg)
}
```

### Subscribe to trades

```go
ch, _ := ws.Subscribe(ctx, "trades", map[string]interface{}{
    "instId": "BTC-USDT",
})

for msg := range ch {
    fmt.Printf("Trade: %s\n", msg)
}
```

### Private channel - orders

```go
ws := okx.NewWSClient(apiKey, secret, pass, okx.WSPrivateURL)

ctx := context.Background()
ws.Connect(ctx)
ws.Login(ctx)
defer ws.Close()

ch, _ := ws.Subscribe(ctx, "orders", map[string]interface{}{
    "instType": "SPOT",
})

for msg := range ch {
    fmt.Printf("Order update: %s\n", msg)
}
```

## Advanced examples

### Place order with stop loss and take profit

```go
px := "30000"
tpPx := "35000"
slPx := "28000"

order := models.PlaceOrderRequest{
    InstID:       "BTC-USDT",
    TdMode:       "cash",
    Side:         "buy",
    OrdType:      "limit",
    Px:           &px,
    Sz:           "0.001",
    TpTriggerPx:  &tpPx,
    TpOrdPx:      &tpPx,
    SlTriggerPx:  &slPx,
    SlOrdPx:      &slPx,
}

result, err := client.Trade.PlaceOrder(ctx, order)
```

### Batch order placement

```go
orders := []models.PlaceOrderRequest{
    {
        InstID:  "BTC-USDT",
        TdMode:  "cash",
        Side:    "buy",
        OrdType: "market",
        Sz:      "0.001",
    },
    {
        InstID:  "ETH-USDT",
        TdMode:  "cash",
        Side:    "buy",
        OrdType: "market",
        Sz:      "0.01",
    },
}

results, err := client.Trade.PlaceBatchOrders(ctx, orders)
for _, r := range results {
    fmt.Printf("Order %s: %s\n", r.OrdID, r.SMsg)
}
```

### Transfer between accounts

```go
transfer := models.TransferRequest{
    Ccy:  "USDT",
    Amt:  "100",
    From: "6",  // funding
    To:   "18", // trading
}

result, err := client.Asset.Transfer(ctx, transfer)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Transfer ID: %s\n", result[0].TransID)
```

### Get order history with pagination

```go
var allOrders []models.Order
after := ""
limit := "100"

for {
    orders, err := client.Trade.GetOrdersHistory(
        ctx, "SPOT", nil, nil, nil, nil, nil, nil,
        &after, nil, nil, nil, &limit,
    )
    if err != nil {
        break
    }
    if len(orders) == 0 {
        break
    }
    
    allOrders = append(allOrders, orders...)
    after = orders[len(orders)-1].OrdID
}

fmt.Printf("Total orders: %d\n", len(allOrders))
```

### Monitor multiple instruments

```go
ws := okx.NewWSClient("", "", "", okx.WSPublicURL)
ws.Connect(ctx)
defer ws.Close()

instruments := []string{"BTC-USDT", "ETH-USDT", "SOL-USDT"}

for _, inst := range instruments {
    ch, _ := ws.Subscribe(ctx, "tickers", map[string]interface{}{
        "instId": inst,
    })
    
    go func(instrument string, channel <-chan []byte) {
        for msg := range channel {
            fmt.Printf("[%s] %s\n", instrument, msg)
        }
    }(inst, ch)
}

select {} // Keep running
```

### Set leverage for futures

```go
req := models.SetLeverageRequest{
    InstID:  "BTC-USDT-SWAP",
    Lever:   "10",
    MgnMode: "cross",
}

result, err := client.Account.SetLeverage(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Leverage set to %s\n", result[0].Lever)
```

### Get funding rate

```go
rate, err := client.Public.GetFundingRate(ctx, "BTC-USDT-SWAP")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Current: %s\n", rate[0].FundingRate)
fmt.Printf("Next: %s\n", rate[0].NextFundingRate)
fmt.Printf("Time: %s\n", rate[0].NextFundingTime)
```

### Close all positions

```go
positions, _ := client.Account.GetPositions(ctx, nil, nil)

for _, pos := range positions {
    req := models.ClosePositionRequest{
        InstID:  pos.InstID,
        MgnMode: pos.MgnMode,
    }
    
    _, err := client.Trade.ClosePosition(ctx, req)
    if err != nil {
        log.Printf("Failed to close %s: %v", pos.InstID, err)
    }
}
```

## Error handling examples

### Retry on rate limit

```go
var balances []models.Balance
var err error

for i := 0; i < 3; i++ {
    balances, err = client.Account.GetBalance(ctx, nil)
    if err == nil {
        break
    }
    
    if errors.Is(err, okx.ErrRateLimited) {
        time.Sleep(time.Second * time.Duration(i+1))
        continue
    }
    
    break
}
```

### Handle specific error codes

```go
_, err := client.Trade.PlaceOrder(ctx, order)
if err != nil {
    if okxErr, ok := err.(*okx.OKXError); ok {
        switch okxErr.Code {
        case "51008":
            log.Println("Insufficient balance")
        case "51009":
            log.Println("Order amount too small")
        case "51010":
            log.Println("Order price out of range")
        default:
            log.Printf("Error: %s", okxErr.Message)
        }
    }
}
```

## Next Steps

- [API Reference](API-Reference)
- [FAQ](FAQ)

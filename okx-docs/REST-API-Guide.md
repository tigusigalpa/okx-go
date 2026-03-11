# REST API Guide

## Overview

The REST client provides access to 335 endpoints across 16 categories.

## Client initialization

```go
client := okx.NewRestClient(
    "api-key",
    "secret-key",
    "passphrase",
    okx.WithTimeout(30*time.Second),
)
```

## API categories

### Account (53 endpoints)

Account and position management.

```go
// Get balance
balances, err := client.Account.GetBalance(ctx, nil)

// Get positions
positions, err := client.Account.GetPositions(ctx, nil, nil)

// Set leverage
req := models.SetLeverageRequest{
    InstID:  "BTC-USDT-SWAP",
    Lever:   "10",
    MgnMode: "cross",
}
result, err := client.Account.SetLeverage(ctx, req)

// Get account config
config, err := client.Account.GetAccountConfig(ctx)
```

### Trade (32 endpoints)

Order placement and management.

```go
// Place order
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

// Cancel order
cancel := models.CancelOrderRequest{
    InstID: "BTC-USDT",
    OrdID:  &orderID,
}
_, err = client.Trade.CancelOrder(ctx, cancel)

// Get pending orders
orders, err := client.Trade.GetOrdersPending(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil)

// Get order history
instType := "SPOT"
history, err := client.Trade.GetOrdersHistory(ctx, instType, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)

// Get fills
fills, err := client.Trade.GetFills(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
```

### Market Data (24 endpoints)

Real-time and historical market data.

```go
// Get ticker
ticker, err := client.Market.GetTicker(ctx, "BTC-USDT")

// Get all tickers
tickers, err := client.Market.GetTickers(ctx, "SPOT", nil, nil)

// Get order book
book, err := client.Market.GetOrderBook(ctx, "BTC-USDT", nil)

// Get candles
bar := "1H"
limit := "100"
candles, err := client.Market.GetCandles(ctx, "BTC-USDT", &bar, nil, nil, &limit)

// Get trades
trades, err := client.Market.GetTrades(ctx, "BTC-USDT", nil)

// Get 24h volume
volume, err := client.Market.Get24hVolume(ctx)
```

### Public Data (24 endpoints)

Public information about instruments and system.

```go
// Get instruments
instruments, err := client.Public.GetInstruments(ctx, "SPOT", nil, nil, nil)

// Get funding rate
rate, err := client.Public.GetFundingRate(ctx, "BTC-USDT-SWAP")

// Get mark price
mark, err := client.Public.GetMarkPrice(ctx, "SWAP", nil, nil, nil)

// Get system time
time, err := client.Public.GetSystemTime(ctx)

// Get position tiers
tiers, err := client.Public.GetPositionTiers(ctx, "SWAP", "cross", nil, nil, nil, nil, nil)
```

### Asset (26 endpoints)

Funding account operations.

```go
// Get currencies
currencies, err := client.Asset.GetCurrencies(ctx, nil)

// Get balances
balances, err := client.Asset.GetBalances(ctx, nil)

// Transfer
transfer := models.TransferRequest{
    Ccy:  "USDT",
    Amt:  "100",
    From: "6",  // funding
    To:   "18", // trading
}
result, err := client.Asset.Transfer(ctx, transfer)

// Get deposit address
addr, err := client.Asset.GetDepositAddress(ctx, "USDT")

// Get deposit history
deposits, err := client.Asset.GetDepositHistory(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil)

// Withdraw
withdraw := models.WithdrawalRequest{
    Ccy:    "USDT",
    Amt:    "100",
    Dest:   "4",
    ToAddr: "0x...",
    Fee:    "1",
}
_, err = client.Asset.Withdrawal(ctx, withdraw)
```

### System (1 endpoint)

System status.

```go
status, err := client.System.GetStatus(ctx, nil)
```

### Support (2 endpoints)

Announcements.

```go
// Get announcement types
types, err := client.Support.GetAnnouncementTypes(ctx)

// Get announcements
announcements, err := client.Support.GetAnnouncements(ctx, nil, nil, nil)
```

### Users (8 endpoints)

Sub-account management.

```go
// Get sub-accounts
subs, err := client.Users.GetSubAccountList(ctx, nil, nil, nil, nil, nil)

// Create sub-account
req := models.CreateSubAccountRequest{
    SubAcct: "sub1",
}
result, err := client.Users.CreateSubAccount(ctx, req)
```

## Common patterns

### Optional parameters

Use pointers for optional params:

```go
ccy := "BTC"
balances, err := client.Account.GetBalance(ctx, &ccy)
```

### Pagination

```go
limit := "100"
after := ""

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
    
    // Process orders
    
    after = orders[len(orders)-1].OrdID
}
```

Or use the Paginator helper:

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

### Context timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

balances, err := client.Account.GetBalance(ctx, nil)
```

### Batch operations

```go
orders := []models.PlaceOrderRequest{
    {InstID: "BTC-USDT", TdMode: "cash", Side: "buy", OrdType: "market", Sz: "0.01"},
    {InstID: "ETH-USDT", TdMode: "cash", Side: "buy", OrdType: "market", Sz: "0.1"},
}

results, err := client.Trade.PlaceBatchOrders(ctx, orders)
```

## Rate limits

OKX enforces rate limits per endpoint. The library includes a configurable rate limiter:

```go
client := okx.NewRestClient(
    apiKey, secret, pass,
    okx.WithRateLimiter(true),
)
```

See [OKX rate limits documentation](https://www.okx.com/docs-v5/en/#overview-rate-limits).

## Next Steps

- [WebSocket Guide](WebSocket-Guide)
- [Error Handling](Error-Handling)
- [Examples](Examples)

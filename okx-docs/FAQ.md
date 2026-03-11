# FAQ

## General

### What is okx-go?

A Go client library for the OKX v5 API covering 335 REST endpoints and 53 WebSocket channels.

### Is it production-ready?

Yes. The library is tested and follows Go best practices.

### What Go version is required?

Go 1.21 or higher (for generics support).

### Is it thread-safe?

Yes. All operations are goroutine-safe.

## Installation

### How do I install it?

```bash
go get github.com/tigusigalpa/okx-go
```

### What are the dependencies?

- `github.com/gorilla/websocket` - WebSocket support
- `github.com/stretchr/testify` - Testing only

## Authentication

### Where do I get API credentials?

Log in to OKX → API Management → Create API Key

### Can I test without real funds?

Yes. Use demo trading mode with demo credentials from https://www.okx.com/demo-trading

### How do I enable demo mode?

```go
client := okx.NewRestClient(key, secret, pass, okx.WithDemoTrading())
```

### Why am I getting "Invalid signature"?

- Check your secret key is correct
- Verify system time is accurate (use NTP)
- Ensure passphrase matches

## Usage

### How do I handle optional parameters?

Use pointers:

```go
ccy := "BTC"
balances, err := client.Account.GetBalance(ctx, &ccy)
```

### How do I paginate results?

Use the `after` parameter or the `Paginator` helper. See [Pagination example](Examples#pagination).

### How do I set a timeout?

Use context:

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
```

### How do I retry on rate limit?

Check for `okx.ErrRateLimited` and implement backoff. See [Error Handling](Error-Handling).

## WebSocket

### How do I connect to WebSocket?

```go
ws := okx.NewWSClient("", "", "", okx.WSPublicURL)
ws.Connect(ctx)
```

### Do I need authentication for public channels?

No. Leave credentials empty for public channels.

### How do I subscribe to multiple instruments?

Call `Subscribe` multiple times or use goroutines. See [WebSocket Guide](WebSocket-Guide).

### Does it auto-reconnect?

Yes. The client automatically reconnects with exponential backoff.

### How do I parse WebSocket messages?

Unmarshal JSON:

```go
var data map[string]interface{}
json.Unmarshal(msg, &data)
```

## Errors

### How do I check for specific errors?

Use `errors.Is` or type assertion:

```go
if errors.Is(err, okx.ErrRateLimited) {
    // handle rate limit
}
```

### What does error code 50011 mean?

Rate limit exceeded. Slow down your requests.

### What does error code 50100 mean?

Invalid API key. Check your credentials.

### Where can I find all error codes?

https://www.okx.com/docs-v5/en/#error-code

## Trading

### How do I place a market order?

```go
order := models.PlaceOrderRequest{
    InstID:  "BTC-USDT",
    TdMode:  "cash",
    Side:    "buy",
    OrdType: "market",
    Sz:      "0.001",
}
```

### How do I place a limit order?

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
```

### How do I cancel an order?

```go
cancel := models.CancelOrderRequest{
    InstID: "BTC-USDT",
    OrdID:  &orderID,
}
client.Trade.CancelOrder(ctx, cancel)
```

### How do I set stop loss / take profit?

Use `TpTriggerPx`, `TpOrdPx`, `SlTriggerPx`, `SlOrdPx` in `PlaceOrderRequest`.

## Performance

### Is it fast?

Yes. Uses connection pooling and minimal allocations.

### Can I use it for high-frequency trading?

Yes, but respect OKX rate limits.

### Should I enable the rate limiter?

Yes for production to avoid hitting limits.

## Troubleshooting

### Connection timeout

- Check network connectivity
- Increase timeout: `okx.WithTimeout(60*time.Second)`
- Check firewall settings

### WebSocket disconnects

The client auto-reconnects. Check logs for errors.

### "Insufficient balance"

You don't have enough funds. Check balance or use demo mode.

### "Order amount too small"

Check instrument's `minSz` via `GetInstruments`.

## Contributing

### How can I contribute?

Fork, branch, PR. See [CONTRIBUTING.md](https://github.com/tigusigalpa/okx-go/blob/main/CONTRIBUTING.md).

### Where do I report bugs?

https://github.com/tigusigalpa/okx-go/issues

## Support

### Where can I get help?

- [GitHub Issues](https://github.com/tigusigalpa/okx-go/issues)
- [OKX API Docs](https://www.okx.com/docs-v5/en/)
- Email: sovletig@gmail.com

### Is there a Telegram/Discord?

Not yet. Use GitHub issues for now.

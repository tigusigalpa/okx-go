# API Reference

## REST Client

### Account

- `GetBalance(ctx, ccy)` - Get account balance
- `GetPositions(ctx, instType, instID)` - Get positions
- `GetAccountConfig(ctx)` - Get account configuration
- `SetLeverage(ctx, req)` - Set leverage
- `SetPositionMode(ctx, posMode)` - Set position mode
- `GetMaxSize(ctx, instID, tdMode, ccy, px)` - Get max order size
- `GetTradeFee(ctx, instType, instID, uly, category, instFamily)` - Get trade fee
- `GetBills(ctx, ...)` - Get account bills
- `GetInterestAccrued(ctx, ...)` - Get interest accrued

[Full account API docs](https://www.okx.com/docs-v5/en/#trading-account-rest-api)

### Trade

- `PlaceOrder(ctx, req)` - Place order
- `PlaceBatchOrders(ctx, reqs)` - Place batch orders
- `CancelOrder(ctx, req)` - Cancel order
- `CancelBatchOrders(ctx, reqs)` - Cancel batch orders
- `AmendOrder(ctx, req)` - Amend order
- `GetOrder(ctx, instID, ordID, clOrdID)` - Get order details
- `GetOrdersPending(ctx, ...)` - Get pending orders
- `GetOrdersHistory(ctx, ...)` - Get order history
- `GetFills(ctx, ...)` - Get fills
- `PlaceAlgoOrder(ctx, req)` - Place algo order
- `CancelAlgoOrder(ctx, reqs)` - Cancel algo order

[Full trade API docs](https://www.okx.com/docs-v5/en/#order-book-trading-trade-rest-api)

### Market

- `GetTicker(ctx, instID)` - Get ticker
- `GetTickers(ctx, instType, uly, instFamily)` - Get all tickers
- `GetOrderBook(ctx, instID, sz)` - Get order book
- `GetCandles(ctx, instID, bar, after, before, limit)` - Get candles
- `GetTrades(ctx, instID, limit)` - Get trades
- `Get24hVolume(ctx)` - Get 24h volume

[Full market API docs](https://www.okx.com/docs-v5/en/#order-book-trading-market-data-rest-api)

### Public

- `GetInstruments(ctx, instType, uly, instFamily, instID)` - Get instruments
- `GetFundingRate(ctx, instID)` - Get funding rate
- `GetMarkPrice(ctx, instType, uly, instFamily, instID)` - Get mark price
- `GetSystemTime(ctx)` - Get system time
- `GetPositionTiers(ctx, ...)` - Get position tiers

[Full public API docs](https://www.okx.com/docs-v5/en/#public-data-rest-api)

### Asset

- `GetCurrencies(ctx, ccy)` - Get currencies
- `GetBalances(ctx, ccy)` - Get balances
- `Transfer(ctx, req)` - Transfer funds
- `Withdrawal(ctx, req)` - Withdraw
- `GetDepositAddress(ctx, ccy)` - Get deposit address
- `GetDepositHistory(ctx, ...)` - Get deposit history
- `GetWithdrawalHistory(ctx, ...)` - Get withdrawal history

[Full asset API docs](https://www.okx.com/docs-v5/en/#funding-account-rest-api)

## WebSocket Client

### Methods

- `Connect(ctx)` - Connect to WebSocket
- `Login(ctx)` - Authenticate (private channels)
- `Subscribe(ctx, channel, args)` - Subscribe to channel
- `Unsubscribe(channel, args)` - Unsubscribe from channel
- `Close()` - Close connection

### Public Channels

`tickers`, `books`, `trades`, `candle1m`, `mark-price`, `funding-rate`, `index-tickers`, `liquidation-orders`

### Private Channels

`account`, `positions`, `orders`, `orders-algo`, `fills`, `balance_and_position`, `deposit-info`, `withdrawal-info`

[Full WebSocket docs](https://www.okx.com/docs-v5/en/#overview-websocket)

## Models

All request/response types are in `github.com/tigusigalpa/okx-go/models`

Key types:
- `PlaceOrderRequest`
- `Balance`
- `Position`
- `Order`
- `Fill`
- `Ticker`
- `Candle`

See [GoDoc](https://pkg.go.dev/github.com/tigusigalpa/okx-go) for complete reference.

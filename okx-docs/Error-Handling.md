# Error Handling

## Error types

### OKXError

Custom error type for OKX API errors.

```go
type OKXError struct {
    Code    string // OKX error code
    Message string // Error message
    Raw     []byte // Raw response
}
```

### Sentinel errors

Pre-defined errors for common scenarios:

- `ErrUnauthorized` - Invalid credentials
- `ErrRateLimited` - Rate limit exceeded
- `ErrInvalidParameter` - Invalid request parameter
- `ErrNotFound` - Resource not found
- `ErrInternalServer` - Server error
- `ErrBadRequest` - Malformed request
- `ErrForbidden` - Forbidden
- `ErrServiceUnavail` - Service unavailable

## Checking errors

### Using errors.Is

```go
balances, err := client.Account.GetBalance(ctx, nil)
if err != nil {
    if errors.Is(err, okx.ErrUnauthorized) {
        log.Fatal("Invalid API credentials")
    }
    if errors.Is(err, okx.ErrRateLimited) {
        time.Sleep(time.Second)
        // Retry
    }
}
```

### Type assertion

```go
balances, err := client.Account.GetBalance(ctx, nil)
if err != nil {
    if okxErr, ok := err.(*okx.OKXError); ok {
        fmt.Printf("OKX error code: %s\n", okxErr.Code)
        fmt.Printf("Message: %s\n", okxErr.Message)
        
        // Handle specific codes
        switch okxErr.Code {
        case "50000":
            // Body cannot be empty
        case "50001":
            // Service temporarily unavailable
        case "50011":
            // Rate limit exceeded
        }
    }
}
```

## Common error codes

| Code | Message | Meaning |
|------|---------|---------|
| 50000 | Body cannot be empty | Missing request body |
| 50001 | Service temporarily unavailable | Try again later |
| 50002 | JSON syntax error | Invalid JSON |
| 50004 | Endpoint request timeout | Request took too long |
| 50005 | API endpoint does not exist | Wrong URL |
| 50006 | Invalid Content-Type | Must be application/json |
| 50011 | Rate limit exceeded | Slow down |
| 50013 | System busy | Try again |
| 50014 | Parameter error | Invalid parameter |
| 50100 | API key invalid | Check API key |
| 50101 | Timestamp invalid | Check system time |
| 50102 | Signature invalid | Check secret key |
| 50103 | Passphrase incorrect | Check passphrase |
| 50104 | API key expired | Renew API key |
| 50105 | IP not whitelisted | Add IP to whitelist |
| 50111 | Invalid sign | Signature error |

Full list: https://www.okx.com/docs-v5/en/#error-code

## Retry logic

### Simple retry

```go
func getBalanceWithRetry(client *okx.RestClient, ctx context.Context, maxRetries int) ([]models.Balance, error) {
    var balances []models.Balance
    var err error
    
    for i := 0; i < maxRetries; i++ {
        balances, err = client.Account.GetBalance(ctx, nil)
        if err == nil {
            return balances, nil
        }
        
        if errors.Is(err, okx.ErrRateLimited) {
            time.Sleep(time.Second * time.Duration(i+1))
            continue
        }
        
        return nil, err
    }
    
    return nil, err
}
```

### Exponential backoff

```go
func retryWithBackoff(fn func() error, maxRetries int) error {
    backoff := time.Second
    
    for i := 0; i < maxRetries; i++ {
        err := fn()
        if err == nil {
            return nil
        }
        
        if !errors.Is(err, okx.ErrRateLimited) {
            return err
        }
        
        time.Sleep(backoff)
        backoff *= 2
        if backoff > 30*time.Second {
            backoff = 30 * time.Second
        }
    }
    
    return fmt.Errorf("max retries exceeded")
}

// Usage
err := retryWithBackoff(func() error {
    _, err := client.Account.GetBalance(ctx, nil)
    return err
}, 5)
```

## Context errors

### Timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

balances, err := client.Account.GetBalance(ctx, nil)
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        log.Println("Request timed out")
    }
}
```

### Cancellation

```go
ctx, cancel := context.WithCancel(context.Background())

go func() {
    time.Sleep(2 * time.Second)
    cancel()
}()

balances, err := client.Account.GetBalance(ctx, nil)
if err != nil {
    if errors.Is(err, context.Canceled) {
        log.Println("Request was cancelled")
    }
}
```

## WebSocket errors

### Connection errors

```go
ws := okx.NewWSClient("", "", "", okx.WSPublicURL)

ctx := context.Background()
if err := ws.Connect(ctx); err != nil {
    log.Printf("Connection failed: %v", err)
    // Handle connection error
}
```

### Authentication errors

```go
if err := ws.Login(ctx); err != nil {
    log.Printf("Authentication failed: %v", err)
    // Check credentials
}
```

### Subscription errors

```go
ch, err := ws.Subscribe(ctx, "tickers", map[string]interface{}{
    "instId": "INVALID",
})
if err != nil {
    log.Printf("Subscription failed: %v", err)
    // Check channel name and parameters
}
```

## Best practices

### Always check errors

```go
// BAD
balances, _ := client.Account.GetBalance(ctx, nil)

// GOOD
balances, err := client.Account.GetBalance(ctx, nil)
if err != nil {
    return fmt.Errorf("get balance: %w", err)
}
```

### Wrap errors

```go
balances, err := client.Account.GetBalance(ctx, nil)
if err != nil {
    return fmt.Errorf("failed to get account balance: %w", err)
}
```

### Log errors

```go
balances, err := client.Account.GetBalance(ctx, nil)
if err != nil {
    log.Printf("Error getting balance: %v", err)
    return err
}
```

### Handle specific errors

```go
balances, err := client.Account.GetBalance(ctx, nil)
if err != nil {
    if errors.Is(err, okx.ErrUnauthorized) {
        // Refresh credentials
    } else if errors.Is(err, okx.ErrRateLimited) {
        // Implement backoff
    } else {
        // Generic error handling
    }
}
```

## Next Steps

- [Configuration](Configuration)
- [Examples](Examples)
- [FAQ](FAQ)

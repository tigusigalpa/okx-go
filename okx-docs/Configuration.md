# Configuration

## Client options

Configure the REST client using functional options.

## Available options

### WithHTTPClient

Use a custom HTTP client.

```go
httpClient := &http.Client{
    Timeout: 60 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
}

client := okx.NewRestClient(
    apiKey, secret, pass,
    okx.WithHTTPClient(httpClient),
)
```

### WithBaseURL

Override the base URL (for testing or proxies).

```go
client := okx.NewRestClient(
    apiKey, secret, pass,
    okx.WithBaseURL("https://custom.okx.com"),
)
```

### WithDemoTrading

Enable demo trading mode.

```go
client := okx.NewRestClient(
    apiKey, secret, pass,
    okx.WithDemoTrading(),
)
```

Adds header: `x-simulated-trading: 1`

### WithTimeout

Set request timeout.

```go
client := okx.NewRestClient(
    apiKey, secret, pass,
    okx.WithTimeout(30*time.Second),
)
```

### WithRateLimiter

Enable built-in rate limiter.

```go
client := okx.NewRestClient(
    apiKey, secret, pass,
    okx.WithRateLimiter(true),
)
```

### WithLogger

Use custom logger.

```go
import "log/slog"

logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

client := okx.NewRestClient(
    apiKey, secret, pass,
    okx.WithLogger(okx.NewSlogLogger(logger)),
)
```

## Combining options

```go
client := okx.NewRestClient(
    apiKey,
    secret,
    pass,
    okx.WithDemoTrading(),
    okx.WithTimeout(30*time.Second),
    okx.WithRateLimiter(true),
    okx.WithLogger(logger),
)
```

## Custom logger

Implement the `Logger` interface:

```go
type Logger interface {
    Debug(msg string, args ...any)
    Info(msg string, args ...any)
    Warn(msg string, args ...any)
    Error(msg string, args ...any)
}
```

Example:

```go
type MyLogger struct {
    logger *log.Logger
}

func (l *MyLogger) Debug(msg string, args ...any) {
    l.logger.Printf("[DEBUG] %s %v", msg, args)
}

func (l *MyLogger) Info(msg string, args ...any) {
    l.logger.Printf("[INFO] %s %v", msg, args)
}

func (l *MyLogger) Warn(msg string, args ...any) {
    l.logger.Printf("[WARN] %s %v", msg, args)
}

func (l *MyLogger) Error(msg string, args ...any) {
    l.logger.Printf("[ERROR] %s %v", msg, args)
}

// Usage
myLogger := &MyLogger{logger: log.New(os.Stdout, "", log.LstdFlags)}
client := okx.NewRestClient(
    apiKey, secret, pass,
    okx.WithLogger(myLogger),
)
```

## WebSocket options

### WithWSDemo

Enable demo mode for WebSocket.

```go
ws := okx.NewWSClient(
    apiKey, secret, pass,
    okx.WSPrivateURL,
    okx.WithWSDemo(),
)
```

### WithWSLogger

Custom logger for WebSocket.

```go
ws := okx.NewWSClient(
    apiKey, secret, pass,
    okx.WSPrivateURL,
    okx.WithWSLogger(logger),
)
```

## Environment-based configuration

```go
func NewClientFromEnv() *okx.RestClient {
    apiKey := os.Getenv("OKX_API_KEY")
    secret := os.Getenv("OKX_SECRET_KEY")
    pass := os.Getenv("OKX_PASSPHRASE")
    
    opts := []okx.Option{}
    
    if os.Getenv("OKX_DEMO") == "true" {
        opts = append(opts, okx.WithDemoTrading())
    }
    
    if timeout := os.Getenv("OKX_TIMEOUT"); timeout != "" {
        d, _ := time.ParseDuration(timeout)
        opts = append(opts, okx.WithTimeout(d))
    }
    
    return okx.NewRestClient(apiKey, secret, pass, opts...)
}
```

## Production configuration

```go
client := okx.NewRestClient(
    os.Getenv("OKX_API_KEY"),
    os.Getenv("OKX_SECRET_KEY"),
    os.Getenv("OKX_PASSPHRASE"),
    okx.WithTimeout(30*time.Second),
    okx.WithRateLimiter(true),
    okx.WithLogger(productionLogger),
)
```

## Testing configuration

```go
client := okx.NewRestClient(
    "test-key",
    "test-secret",
    "test-pass",
    okx.WithDemoTrading(),
    okx.WithTimeout(10*time.Second),
    okx.WithLogger(testLogger),
)
```

## Next Steps

- [Authentication](Authentication)
- [Error Handling](Error-Handling)
- [Examples](Examples)

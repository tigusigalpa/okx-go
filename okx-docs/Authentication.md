# Authentication

## API credentials

OKX uses API key authentication with three components:

- **API Key**: Public identifier
- **Secret Key**: Used for signing requests
- **Passphrase**: Set when creating the API key

## Creating credentials

### Production

1. Log in to OKX
2. Navigate to API management
3. Click "Create API"
4. Set permissions (read, trade, withdraw)
5. Set IP whitelist (recommended)
6. Save credentials securely

### Demo trading

1. Go to https://www.okx.com/demo-trading
2. Create demo account
3. Generate demo API credentials
4. Use with `okx.WithDemoTrading()` option

## How it works

### REST authentication

Each authenticated request includes headers:

- `OK-ACCESS-KEY`: Your API key
- `OK-ACCESS-SIGN`: HMAC-SHA256 signature
- `OK-ACCESS-TIMESTAMP`: UTC timestamp (ISO 8601)
- `OK-ACCESS-PASSPHRASE`: Your passphrase

The signature is computed as:

```
signature = Base64(HMAC-SHA256(timestamp + method + requestPath + body, secretKey))
```

The library handles this automatically.

### WebSocket authentication

For private channels, send a login message after connecting:

```json
{
  "op": "login",
  "args": [{
    "apiKey": "...",
    "passphrase": "...",
    "timestamp": "...",
    "sign": "..."
  }]
}
```

The library handles this via `ws.Login(ctx)`.

## Security best practices

### Store credentials securely

Never hardcode credentials:

```go
// BAD
client := okx.NewRestClient("abc123", "secret", "pass")

// GOOD
client := okx.NewRestClient(
    os.Getenv("OKX_API_KEY"),
    os.Getenv("OKX_SECRET_KEY"),
    os.Getenv("OKX_PASSPHRASE"),
)
```

### Use environment variables

```bash
export OKX_API_KEY="your-api-key"
export OKX_SECRET_KEY="your-secret-key"
export OKX_PASSPHRASE="your-passphrase"
```

### Set minimal permissions

Only grant permissions you need:

- Read-only for monitoring
- Trade for automated trading
- Withdraw only if necessary

### Use IP whitelist

Restrict API access to specific IPs in OKX settings.

### Rotate keys regularly

Change API credentials periodically.

### Test with demo first

Always test new code with demo credentials before using production.

## Demo vs Production

### Demo mode

```go
client := okx.NewRestClient(
    demoKey,
    demoSecret,
    demoPass,
    okx.WithDemoTrading(),
)
```

Adds header: `x-simulated-trading: 1`

### Production mode

```go
client := okx.NewRestClient(
    prodKey,
    prodSecret,
    prodPass,
)
```

No additional headers.

## Troubleshooting

### "Invalid signature"

- Check timestamp is UTC
- Verify secret key is correct
- Ensure request path matches exactly
- Check body is properly formatted

### "Invalid API key"

- Verify API key is active
- Check key hasn't been deleted
- Ensure key has required permissions

### "IP not whitelisted"

- Add your IP to whitelist in OKX settings
- Or remove IP restriction

### "Passphrase incorrect"

- Passphrase is case-sensitive
- Cannot be recovered if lost
- Must create new API key if forgotten

## Next Steps

- [Configuration](Configuration)
- [Error Handling](Error-Handling)

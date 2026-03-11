# Contributing

Thanks for your interest in contributing to okx-go!

## Quick start

1. Fork the repo
2. Clone your fork
3. Create a branch
4. Make changes
5. Run tests
6. Submit PR

## Development setup

```bash
git clone https://github.com/YOUR_USERNAME/okx-go.git
cd okx-go
go mod download
go test ./...
```

## Making changes

### Code style

- Follow standard Go conventions
- Run `go fmt`
- Run `go vet`
- Use `golangci-lint` if available

### Adding endpoints

1. Add models in `models/[category].go`
2. Implement in `rest/[category]/[category].go`
3. Add tests
4. Update docs

Example:

```go
// models/account.go
type NewFeatureRequest struct {
    Param string `json:"param"`
}

// rest/account/account.go
func (c *Client) NewFeature(ctx context.Context, req models.NewFeatureRequest) error {
    return c.doFunc(ctx, http.MethodPost, "/api/v5/account/new-feature", nil, req, nil)
}
```

### Testing

```bash
# Unit tests
go test ./...

# With coverage
go test -cover ./...

# Integration tests (requires credentials)
OKX_API_KEY=... OKX_SECRET_KEY=... OKX_PASSPHRASE=... go test -tags=integration ./...
```

### Documentation

- Update README if adding features
- Add examples for new functionality
- Update wiki if needed

## Pull requests

### Before submitting

- [ ] Tests pass
- [ ] Code is formatted
- [ ] No linter warnings
- [ ] Documentation updated
- [ ] Commit messages are clear

### PR description

Include:
- What changed
- Why it changed
- How to test it
- Related issues

### Review process

1. Automated checks run
2. Maintainer reviews code
3. Address feedback
4. Merge when approved

## Commit messages

Use clear, descriptive messages:

```
Add GetNewFeature endpoint

- Implement new endpoint
- Add request/response models
- Include tests
```

## Code of conduct

- Be respectful
- Be professional
- Help others
- No harassment

## Questions?

Open an issue or email sovletig@gmail.com

## License

By contributing, you agree your code will be licensed under MIT.

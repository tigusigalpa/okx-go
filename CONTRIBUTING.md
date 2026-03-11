# Contributing to okx-go

Thank you for your interest in contributing to okx-go! This document provides guidelines and instructions for contributing.

## Code of Conduct

Be respectful and professional in all interactions. We're all here to build something great together.

## How to Contribute

### Reporting Bugs

Before creating a bug report, please check existing issues to avoid duplicates. When creating a bug report, include:

- **Clear title and description**
- **Steps to reproduce** the issue
- **Expected behavior** vs **actual behavior**
- **Code samples** if applicable
- **Environment details** (Go version, OS, etc.)

### Suggesting Enhancements

Enhancement suggestions are welcome! Please include:

- **Clear use case** for the enhancement
- **Detailed description** of the proposed functionality
- **Examples** of how it would be used
- **Potential implementation approach** (optional)

### Pull Requests

1. **Fork the repository** and create your branch from `main`
2. **Follow Go conventions** and existing code style
3. **Add tests** for new functionality
4. **Update documentation** as needed
5. **Ensure all tests pass** (`go test ./...`)
6. **Run linters** (`golangci-lint run`)
7. **Write clear commit messages**

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Git

### Setup Steps

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/okx-go.git
cd okx-go

# Install dependencies
go mod download

# Run tests
go test ./...

# Run linters (optional but recommended)
golangci-lint run
```

### Running Tests

```bash
# Run all unit tests
go test -v ./...

# Run with coverage
go test -v -cover ./...

# Run integration tests (requires API credentials)
export OKX_API_KEY="your-api-key"
export OKX_SECRET_KEY="your-secret-key"
export OKX_PASSPHRASE="your-passphrase"
go test -v -tags=integration ./...
```

## Code Style Guidelines

### General Principles

- **Idiomatic Go**: Follow standard Go conventions
- **Simplicity**: Prefer simple, readable code over clever solutions
- **Documentation**: All exported functions must have doc comments
- **Error handling**: Always handle errors appropriately
- **Context**: Use `context.Context` for cancellation and timeouts

### Naming Conventions

- **Exported names**: Use PascalCase (e.g., `GetBalance`)
- **Unexported names**: Use camelCase (e.g., `doRequest`)
- **Acronyms**: Keep uppercase (e.g., `HTTPClient`, `APIKey`)
- **Interfaces**: Use `-er` suffix when appropriate (e.g., `Logger`)

### Code Organization

- **One concept per file**: Keep files focused
- **Package structure**: Follow existing package layout
- **Imports**: Group standard library, external, and internal imports
- **Constants**: Define at package level when shared

### Documentation

```go
// GetBalance retrieves the account balance for the specified currency.
// If ccy is nil, returns balances for all currencies.
//
// Example:
//   balances, err := client.Account.GetBalance(ctx, nil)
//   if err != nil {
//       return err
//   }
func (c *Client) GetBalance(ctx context.Context, ccy *string) ([]models.Balance, error) {
    // implementation
}
```

### Testing

- **Table-driven tests** for multiple test cases
- **Descriptive test names** using `TestFunctionName_Scenario`
- **Test coverage** for all public APIs
- **Mock external dependencies** in unit tests
- **Integration tests** tagged with `//go:build integration`

Example:

```go
func TestGetBalance_Success(t *testing.T) {
    tests := []struct {
        name     string
        ccy      *string
        expected int
    }{
        {
            name:     "all currencies",
            ccy:      nil,
            expected: 5,
        },
        {
            name:     "specific currency",
            ccy:      stringPtr("BTC"),
            expected: 1,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

## Adding New Endpoints

When adding support for new OKX API endpoints:

1. **Add models** in `models/` directory
2. **Implement endpoint** in appropriate `rest/` subdirectory
3. **Add tests** for the new endpoint
4. **Update documentation** in README.md
5. **Add example** if it's a commonly used endpoint

Example structure:

```go
// In models/account.go
type NewFeatureRequest struct {
    Param1 string  `json:"param1"`
    Param2 *string `json:"param2,omitempty"`
}

type NewFeatureResponse struct {
    Result string `json:"result"`
    Status string `json:"status"`
}

// In rest/account/account.go
func (c *Client) NewFeature(ctx context.Context, req models.NewFeatureRequest) ([]models.NewFeatureResponse, error) {
    var result []models.NewFeatureResponse
    if err := c.doFunc(ctx, http.MethodPost, "/api/v5/account/new-feature", nil, req, &result); err != nil {
        return nil, err
    }
    return result, nil
}

// In rest/account/account_test.go
func TestNewFeature(t *testing.T) {
    // test implementation
}
```

## Commit Message Guidelines

Use clear, descriptive commit messages:

```
Add support for new trading endpoint

- Implement GetAdvancedOrders method
- Add corresponding request/response models
- Include unit tests and documentation
- Update README with new endpoint info
```

Format:
- **First line**: Brief summary (50 chars or less)
- **Blank line**
- **Body**: Detailed description (wrap at 72 chars)
- **Use imperative mood**: "Add feature" not "Added feature"

## Review Process

1. **Automated checks** must pass (tests, linters)
2. **Code review** by maintainers
3. **Address feedback** promptly
4. **Squash commits** if requested
5. **Merge** once approved

## Questions?

- **Open an issue** for questions about contributing
- **Check existing issues** for similar questions
- **Be patient** - maintainers are volunteers

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

Thank you for contributing to okx-go! 🚀

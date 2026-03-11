package okx

import (
	"log/slog"
	"net/http"
	"time"
)

// Option is a functional option for configuring the Client.
type Option func(*Client)

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithBaseURL overrides the base URL for API requests.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithDemoTrading enables demo trading mode (x-simulated-trading: 1).
func WithDemoTrading() Option {
	return func(c *Client) {
		c.isDemo = true
	}
}

// WithTimeout sets the HTTP request timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		if c.httpClient == nil {
			c.httpClient = &http.Client{}
		}
		c.httpClient.Timeout = d
	}
}

// WithRateLimiter enables or disables the built-in rate limiter.
func WithRateLimiter(enabled bool) Option {
	return func(c *Client) {
		c.rateLimiterEnabled = enabled
	}
}

// WithLogger sets a custom logger.
func WithLogger(logger Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

// Logger is the interface for logging.
type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

// noopLogger is a no-op logger implementation.
type noopLogger struct{}

func (n *noopLogger) Debug(msg string, args ...any) {}
func (n *noopLogger) Info(msg string, args ...any)  {}
func (n *noopLogger) Warn(msg string, args ...any)  {}
func (n *noopLogger) Error(msg string, args ...any) {}

// slogLogger wraps slog.Logger to implement the Logger interface.
type slogLogger struct {
	logger *slog.Logger
}

func (s *slogLogger) Debug(msg string, args ...any) {
	s.logger.Debug(msg, args...)
}

func (s *slogLogger) Info(msg string, args ...any) {
	s.logger.Info(msg, args...)
}

func (s *slogLogger) Warn(msg string, args ...any) {
	s.logger.Warn(msg, args...)
}

func (s *slogLogger) Error(msg string, args ...any) {
	s.logger.Error(msg, args...)
}

// NewSlogLogger creates a new Logger from slog.Logger.
func NewSlogLogger(logger *slog.Logger) Logger {
	return &slogLogger{logger: logger}
}

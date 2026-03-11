package okx

import (
	"errors"
	"fmt"
)

// OKXError represents an error returned by the OKX API.
type OKXError struct {
	Code    string
	Message string
	Raw     []byte
}

func (e *OKXError) Error() string {
	return fmt.Sprintf("OKX API error: code=%s, message=%s", e.Code, e.Message)
}

// Sentinel errors for common failure scenarios
var (
	ErrUnauthorized     = errors.New("unauthorized: invalid API credentials")
	ErrRateLimited      = errors.New("rate limit exceeded")
	ErrInvalidParameter = errors.New("invalid parameter")
	ErrNotFound         = errors.New("resource not found")
	ErrInternalServer   = errors.New("internal server error")
	ErrBadRequest       = errors.New("bad request")
	ErrForbidden        = errors.New("forbidden")
	ErrServiceUnavail   = errors.New("service unavailable")
)

// MapErrorCode maps OKX error codes to sentinel errors
func MapErrorCode(code string) error {
	switch code {
	case "50100", "50101", "50102", "50103", "50104", "50105", "50106", "50107", "50108", "50109", "50110", "50111", "50112", "50113":
		return ErrUnauthorized
	case "50011":
		return ErrRateLimited
	case "50000", "50001", "50002", "50004", "50005", "50006", "50007", "50008", "50009", "50010":
		return ErrBadRequest
	case "50014":
		return ErrNotFound
	case "50012", "50013":
		return ErrInvalidParameter
	case "50003":
		return ErrServiceUnavail
	case "1":
		return ErrInternalServer
	default:
		return nil
	}
}

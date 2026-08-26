package models

// SubAccount represents an OKX API request or response value.
type SubAccount struct {
	Enable      bool   `json:"enable"`
	SubAcct     string `json:"subAcct"`
	Label       string `json:"label"`
	Mobile      string `json:"mobile"`
	GAuth       bool   `json:"gAuth"`
	CanTransOut bool   `json:"canTransOut"`
	TS          string `json:"ts"`
}

// SubAccountAPIKey represents an OKX API request or response value.
type SubAccountAPIKey struct {
	Label  string `json:"label"`
	APIKey string `json:"apiKey"`
	Perm   string `json:"perm"`
	IP     string `json:"ip"`
	TS     string `json:"ts"`
}

// CreateSubAccountRequest represents an OKX API request or response value.
type CreateSubAccountRequest struct {
	SubAcct string  `json:"subAcct"`
	Label   *string `json:"label,omitempty"`
}

// CreateSubAccountResponse represents an OKX API request or response value.
type CreateSubAccountResponse struct {
	SubAcct string `json:"subAcct"`
	Label   string `json:"label"`
	TS      string `json:"ts"`
}

// CreateSubAccountAPIKeyRequest represents an OKX API request or response value.
type CreateSubAccountAPIKeyRequest struct {
	SubAcct    string   `json:"subAcct"`
	Label      string   `json:"label"`
	Passphrase string   `json:"passphrase"`
	IP         *string  `json:"ip,omitempty"`
	Perm       []string `json:"perm"`
}

// CreateSubAccountAPIKeyResponse represents an OKX API request or response value.
type CreateSubAccountAPIKeyResponse struct {
	SubAcct string `json:"subAcct"`
	Label   string `json:"label"`
	APIKey  string `json:"apiKey"`
	Perm    string `json:"perm"`
	IP      string `json:"ip"`
	TS      string `json:"ts"`
}

// ModifySubAccountAPIKeyRequest represents an OKX API request or response value.
type ModifySubAccountAPIKeyRequest struct {
	SubAcct string   `json:"subAcct"`
	APIKey  string   `json:"apiKey"`
	Label   *string  `json:"label,omitempty"`
	Perm    []string `json:"perm,omitempty"`
	IP      *string  `json:"ip,omitempty"`
}

// DeleteSubAccountAPIKeyRequest represents an OKX API request or response value.
type DeleteSubAccountAPIKeyRequest struct {
	SubAcct string `json:"subAcct"`
	APIKey  string `json:"apiKey"`
}

// SetSubAccountTransferOutRequest represents an OKX API request or response value.
type SetSubAccountTransferOutRequest struct {
	SubAcct     string `json:"subAcct"`
	CanTransOut bool   `json:"canTransOut"`
}

// EntrustSubAccountList represents an OKX API request or response value.
type EntrustSubAccountList struct {
	SubAcct string `json:"subAcct"`
}

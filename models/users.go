package models

type SubAccount struct {
	Enable    bool   `json:"enable"`
	SubAcct   string `json:"subAcct"`
	Label     string `json:"label"`
	Mobile    string `json:"mobile"`
	GAuth     bool   `json:"gAuth"`
	CanTransOut bool `json:"canTransOut"`
	TS        string `json:"ts"`
}

type SubAccountAPIKey struct {
	Label      string `json:"label"`
	APIKey     string `json:"apiKey"`
	Perm       string `json:"perm"`
	IP         string `json:"ip"`
	TS         string `json:"ts"`
}

type CreateSubAccountRequest struct {
	SubAcct string `json:"subAcct"`
	Label   *string `json:"label,omitempty"`
}

type CreateSubAccountResponse struct {
	SubAcct string `json:"subAcct"`
	Label   string `json:"label"`
	TS      string `json:"ts"`
}

type CreateSubAccountAPIKeyRequest struct {
	SubAcct    string   `json:"subAcct"`
	Label      string   `json:"label"`
	Passphrase string   `json:"passphrase"`
	IP         *string  `json:"ip,omitempty"`
	Perm       []string `json:"perm"`
}

type CreateSubAccountAPIKeyResponse struct {
	SubAcct    string `json:"subAcct"`
	Label      string `json:"label"`
	APIKey     string `json:"apiKey"`
	Perm       string `json:"perm"`
	IP         string `json:"ip"`
	TS         string `json:"ts"`
}

type ModifySubAccountAPIKeyRequest struct {
	SubAcct    string   `json:"subAcct"`
	APIKey     string   `json:"apiKey"`
	Label      *string  `json:"label,omitempty"`
	Perm       []string `json:"perm,omitempty"`
	IP         *string  `json:"ip,omitempty"`
}

type DeleteSubAccountAPIKeyRequest struct {
	SubAcct string `json:"subAcct"`
	APIKey  string `json:"apiKey"`
}

type SetSubAccountTransferOutRequest struct {
	SubAcct     string `json:"subAcct"`
	CanTransOut bool   `json:"canTransOut"`
}

type EntrustSubAccountList struct {
	SubAcct string `json:"subAcct"`
}

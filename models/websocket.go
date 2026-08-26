package models

// WSRequest represents an OKX API request or response value.
type WSRequest struct {
	Op   string                   `json:"op"`
	Args []map[string]interface{} `json:"args"`
}

// WSResponse represents an OKX API request or response value.
type WSResponse struct {
	Event  string                   `json:"event,omitempty"`
	Code   string                   `json:"code,omitempty"`
	Msg    string                   `json:"msg,omitempty"`
	ConnID string                   `json:"connId,omitempty"`
	Op     string                   `json:"op,omitempty"`
	Data   []map[string]interface{} `json:"data,omitempty"`
	Arg    map[string]interface{}   `json:"arg,omitempty"`
}

// WSLoginRequest represents an OKX API request or response value.
type WSLoginRequest struct {
	Op   string        `json:"op"`
	Args []WSLoginArgs `json:"args"`
}

// WSLoginArgs represents an OKX API request or response value.
type WSLoginArgs struct {
	APIKey     string `json:"apiKey"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}

// WSSubscribeRequest represents an OKX API request or response value.
type WSSubscribeRequest struct {
	Op   string                   `json:"op"`
	Args []map[string]interface{} `json:"args"`
}

// WSUnsubscribeRequest represents an OKX API request or response value.
type WSUnsubscribeRequest struct {
	Op   string                   `json:"op"`
	Args []map[string]interface{} `json:"args"`
}

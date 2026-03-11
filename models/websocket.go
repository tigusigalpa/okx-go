package models

type WSRequest struct {
	Op   string                   `json:"op"`
	Args []map[string]interface{} `json:"args"`
}

type WSResponse struct {
	Event   string                   `json:"event,omitempty"`
	Code    string                   `json:"code,omitempty"`
	Msg     string                   `json:"msg,omitempty"`
	ConnID  string                   `json:"connId,omitempty"`
	Op      string                   `json:"op,omitempty"`
	Data    []map[string]interface{} `json:"data,omitempty"`
	Arg     map[string]interface{}   `json:"arg,omitempty"`
}

type WSLoginRequest struct {
	Op   string        `json:"op"`
	Args []WSLoginArgs `json:"args"`
}

type WSLoginArgs struct {
	APIKey     string `json:"apiKey"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}

type WSSubscribeRequest struct {
	Op   string                   `json:"op"`
	Args []map[string]interface{} `json:"args"`
}

type WSUnsubscribeRequest struct {
	Op   string                   `json:"op"`
	Args []map[string]interface{} `json:"args"`
}

package models

import "encoding/json"

// WSArg identifies an OKX WebSocket channel subscription or push message.
type WSArg struct {
	Channel    string `json:"channel"`
	InstType   string `json:"instType,omitempty"`
	InstID     string `json:"instId,omitempty"`
	InstFamily string `json:"instFamily,omitempty"`
	Uly        string `json:"uly,omitempty"`
	Ccy        string `json:"ccy,omitempty"`
	UID        string `json:"uid,omitempty"`
}

// WSMessage is a typed OKX WebSocket push-data envelope.
//
// Data is decoded without converting financial values to floating-point numbers.
type WSMessage[T any] struct {
	Arg    WSArg  `json:"arg"`
	Action string `json:"action,omitempty"`
	Data   []T    `json:"data"`
}

// DecodeWSMessage decodes an OKX WebSocket push-data envelope into a typed value.
func DecodeWSMessage[T any](payload []byte) (WSMessage[T], error) {
	var message WSMessage[T]
	if err := json.Unmarshal(payload, &message); err != nil {
		return WSMessage[T]{}, err
	}
	return message, nil
}

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
	Action string                   `json:"action,omitempty"`
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

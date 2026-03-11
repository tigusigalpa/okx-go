package okx

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tigusigalpa/okx-go/models"
)

const (
	// WebSocket URLs
	WSPublicURL         = "wss://ws.okx.com:8443/ws/v5/public"
	WSPrivateURL        = "wss://ws.okx.com:8443/ws/v5/private"
	WSBusinessURL       = "wss://ws.okx.com:8443/ws/v5/business"
	WSPublicSBEURL      = "wss://ws.okx.com:8443/ws/v5/public-sbe"
	WSDemoPublicURL     = "wss://wspap.okx.com:8443/ws/v5/public"
	WSDemoPrivateURL    = "wss://wspap.okx.com:8443/ws/v5/private"
	WSDemoBusinessURL   = "wss://wspap.okx.com:8443/ws/v5/business"
	WSDemoPublicSBEURL  = "wss://wspap.okx.com:8443/ws/v5/public-sbe"

	pingInterval = 25 * time.Second
	pongTimeout  = 30 * time.Second
	writeTimeout = 10 * time.Second
	readTimeout  = 60 * time.Second
)

type WSClient struct {
	apiKey     string
	secretKey  string
	passphrase string
	url        string
	conn       *websocket.Conn
	isDemo     bool
	logger     Logger

	mu            sync.RWMutex
	subscriptions map[string]chan []byte
	done          chan struct{}
	reconnect     bool
	authenticated bool
}

type WSOption func(*WSClient)

func WithWSDemo() WSOption {
	return func(ws *WSClient) {
		ws.isDemo = true
	}
}

func WithWSLogger(logger Logger) WSOption {
	return func(ws *WSClient) {
		ws.logger = logger
	}
}

func NewWSClient(apiKey, secretKey, passphrase, url string, opts ...WSOption) *WSClient {
	ws := &WSClient{
		apiKey:        apiKey,
		secretKey:     secretKey,
		passphrase:    passphrase,
		url:           url,
		logger:        &noopLogger{},
		subscriptions: make(map[string]chan []byte),
		done:          make(chan struct{}),
		reconnect:     true,
	}

	for _, opt := range opts {
		opt(ws)
	}

	return ws
}

func (ws *WSClient) Connect(ctx context.Context) error {
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.DialContext(ctx, ws.url, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	ws.mu.Lock()
	ws.conn = conn
	ws.mu.Unlock()

	ws.logger.Info("WebSocket connected", "url", ws.url)

	go ws.readPump()
	go ws.pingPump()

	return nil
}

func (ws *WSClient) Login(ctx context.Context) error {
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	message := timestamp + "GET" + "/users/self/verify"
	h := hmac.New(sha256.New, []byte(ws.secretKey))
	h.Write([]byte(message))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))

	loginReq := models.WSLoginRequest{
		Op: "login",
		Args: []models.WSLoginArgs{
			{
				APIKey:     ws.apiKey,
				Passphrase: ws.passphrase,
				Timestamp:  timestamp,
				Sign:       sign,
			},
		},
	}

	if err := ws.send(loginReq); err != nil {
		return fmt.Errorf("failed to send login request: %w", err)
	}

	ws.mu.Lock()
	ws.authenticated = true
	ws.mu.Unlock()

	ws.logger.Info("WebSocket authenticated")

	return nil
}

func (ws *WSClient) Subscribe(ctx context.Context, channel string, args map[string]interface{}) (<-chan []byte, error) {
	subKey := ws.makeSubKey(channel, args)

	ws.mu.Lock()
	if _, exists := ws.subscriptions[subKey]; exists {
		ws.mu.Unlock()
		return nil, fmt.Errorf("already subscribed to %s", subKey)
	}

	ch := make(chan []byte, 100)
	ws.subscriptions[subKey] = ch
	ws.mu.Unlock()

	subArgs := make(map[string]interface{})
	subArgs["channel"] = channel
	for k, v := range args {
		subArgs[k] = v
	}

	subReq := models.WSSubscribeRequest{
		Op:   "subscribe",
		Args: []map[string]interface{}{subArgs},
	}

	if err := ws.send(subReq); err != nil {
		ws.mu.Lock()
		delete(ws.subscriptions, subKey)
		close(ch)
		ws.mu.Unlock()
		return nil, fmt.Errorf("failed to send subscribe request: %w", err)
	}

	ws.logger.Info("Subscribed to channel", "channel", channel, "args", args)

	return ch, nil
}

func (ws *WSClient) Unsubscribe(channel string, args map[string]interface{}) error {
	subKey := ws.makeSubKey(channel, args)

	ws.mu.Lock()
	ch, exists := ws.subscriptions[subKey]
	if !exists {
		ws.mu.Unlock()
		return fmt.Errorf("not subscribed to %s", subKey)
	}
	delete(ws.subscriptions, subKey)
	close(ch)
	ws.mu.Unlock()

	subArgs := make(map[string]interface{})
	subArgs["channel"] = channel
	for k, v := range args {
		subArgs[k] = v
	}

	unsubReq := models.WSUnsubscribeRequest{
		Op:   "unsubscribe",
		Args: []map[string]interface{}{subArgs},
	}

	if err := ws.send(unsubReq); err != nil {
		return fmt.Errorf("failed to send unsubscribe request: %w", err)
	}

	ws.logger.Info("Unsubscribed from channel", "channel", channel, "args", args)

	return nil
}

func (ws *WSClient) Close() error {
	ws.mu.Lock()
	ws.reconnect = false
	ws.mu.Unlock()

	close(ws.done)

	ws.mu.Lock()
	defer ws.mu.Unlock()

	for _, ch := range ws.subscriptions {
		close(ch)
	}
	ws.subscriptions = make(map[string]chan []byte)

	if ws.conn != nil {
		return ws.conn.Close()
	}

	return nil
}

func (ws *WSClient) send(v interface{}) error {
	ws.mu.RLock()
	conn := ws.conn
	ws.mu.RUnlock()

	if conn == nil {
		return errors.New("WebSocket not connected")
	}

	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	conn.SetWriteDeadline(time.Now().Add(writeTimeout))
	if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	return nil
}

func (ws *WSClient) readPump() {
	defer func() {
		ws.mu.Lock()
		if ws.conn != nil {
			ws.conn.Close()
		}
		ws.mu.Unlock()
	}()

	for {
		select {
		case <-ws.done:
			return
		default:
		}

		ws.mu.RLock()
		conn := ws.conn
		ws.mu.RUnlock()

		if conn == nil {
			return
		}

		conn.SetReadDeadline(time.Now().Add(readTimeout))
		_, message, err := conn.ReadMessage()
		if err != nil {
			ws.logger.Error("WebSocket read error", "error", err)
			if ws.reconnect {
				ws.handleReconnect()
			}
			return
		}

		if string(message) == "pong" {
			ws.logger.Debug("Received pong")
			continue
		}

		ws.handleMessage(message)
	}
}

func (ws *WSClient) pingPump() {
	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ws.done:
			return
		case <-ticker.C:
			ws.mu.RLock()
			conn := ws.conn
			ws.mu.RUnlock()

			if conn == nil {
				return
			}

			conn.SetWriteDeadline(time.Now().Add(writeTimeout))
			if err := conn.WriteMessage(websocket.TextMessage, []byte("ping")); err != nil {
				ws.logger.Error("WebSocket ping error", "error", err)
				return
			}
			ws.logger.Debug("Sent ping")
		}
	}
}

func (ws *WSClient) handleMessage(message []byte) {
	var resp models.WSResponse
	if err := json.Unmarshal(message, &resp); err != nil {
		ws.logger.Error("Failed to unmarshal WebSocket message", "error", err, "message", string(message))
		return
	}

	if resp.Event == "error" {
		ws.logger.Error("WebSocket error event", "code", resp.Code, "msg", resp.Msg)
		return
	}

	if resp.Event == "login" {
		if resp.Code == "0" {
			ws.logger.Info("Login successful")
		} else {
			ws.logger.Error("Login failed", "code", resp.Code, "msg", resp.Msg)
		}
		return
	}

	if resp.Event == "subscribe" || resp.Event == "unsubscribe" {
		ws.logger.Debug("Subscription event", "event", resp.Event, "arg", resp.Arg)
		return
	}

	if resp.Arg != nil {
		channel, ok := resp.Arg["channel"].(string)
		if !ok {
			ws.logger.Warn("No channel in message arg")
			return
		}

		subKey := ws.makeSubKeyFromArg(channel, resp.Arg)

		ws.mu.RLock()
		ch, exists := ws.subscriptions[subKey]
		ws.mu.RUnlock()

		if exists {
			select {
			case ch <- message:
			default:
				ws.logger.Warn("Channel buffer full, dropping message", "channel", channel)
			}
		}
	}
}

func (ws *WSClient) handleReconnect() {
	backoff := time.Second
	maxBackoff := 60 * time.Second

	for {
		select {
		case <-ws.done:
			return
		default:
		}

		ws.logger.Info("Attempting to reconnect", "backoff", backoff)
		time.Sleep(backoff)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := ws.Connect(ctx)
		cancel()

		if err != nil {
			ws.logger.Error("Reconnect failed", "error", err)
			backoff *= 2
			if backoff > maxBackoff {
				backoff = maxBackoff
			}
			continue
		}

		if ws.authenticated {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			if err := ws.Login(ctx); err != nil {
				ws.logger.Error("Re-authentication failed", "error", err)
				cancel()
				continue
			}
			cancel()
		}

		ws.mu.RLock()
		subs := make(map[string]map[string]interface{})
		for key := range ws.subscriptions {
			subs[key] = nil
		}
		ws.mu.RUnlock()

		ws.logger.Info("Reconnected successfully")
		return
	}
}

func (ws *WSClient) makeSubKey(channel string, args map[string]interface{}) string {
	key := channel
	if instID, ok := args["instId"].(string); ok {
		key += ":" + instID
	}
	if instType, ok := args["instType"].(string); ok {
		key += ":" + instType
	}
	if ccy, ok := args["ccy"].(string); ok {
		key += ":" + ccy
	}
	return key
}

func (ws *WSClient) makeSubKeyFromArg(channel string, arg map[string]interface{}) string {
	args := make(map[string]interface{})
	for k, v := range arg {
		if k != "channel" {
			args[k] = v
		}
	}
	return ws.makeSubKey(channel, args)
}

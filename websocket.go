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
	WSPublicURL        = "wss://ws.okx.com:8443/ws/v5/public"
	WSPrivateURL       = "wss://ws.okx.com:8443/ws/v5/private"
	WSBusinessURL      = "wss://ws.okx.com:8443/ws/v5/business"
	WSPublicSBEURL     = "wss://ws.okx.com:8443/ws/v5/public-sbe"
	WSDemoPublicURL    = "wss://wspap.okx.com:8443/ws/v5/public"
	WSDemoPrivateURL   = "wss://wspap.okx.com:8443/ws/v5/private"
	WSDemoBusinessURL  = "wss://wspap.okx.com:8443/ws/v5/business"
	WSDemoPublicSBEURL = "wss://wspap.okx.com:8443/ws/v5/public-sbe"

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
	writeMu       sync.Mutex
	subscriptions map[string]subscription
	done          chan struct{}
	closeOnce     sync.Once
	reconnect     bool
	authenticated bool
}

type subscription struct {
	channel  string
	args     map[string]interface{}
	messages chan []byte
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
		subscriptions: make(map[string]subscription),
		done:          make(chan struct{}),
		reconnect:     true,
	}

	for _, opt := range opts {
		opt(ws)
	}
	if ws.isDemo {
		ws.url = demoWebSocketURL(ws.url)
	}

	return ws
}

func demoWebSocketURL(url string) string {
	switch url {
	case WSPublicURL:
		return WSDemoPublicURL
	case WSPrivateURL:
		return WSDemoPrivateURL
	case WSBusinessURL:
		return WSDemoBusinessURL
	case WSPublicSBEURL:
		return WSDemoPublicSBEURL
	default:
		return url
	}
}

func (ws *WSClient) Connect(ctx context.Context) error {
	return ws.connect(ctx, true)
}

func (ws *WSClient) connect(ctx context.Context, startPingPump bool) error {
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
	if startPingPump {
		go ws.pingPump()
	}

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
	ws.subscriptions[subKey] = subscription{
		channel:  channel,
		args:     cloneArgs(args),
		messages: ch,
	}
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
	sub, exists := ws.subscriptions[subKey]
	if !exists {
		ws.mu.Unlock()
		return fmt.Errorf("not subscribed to %s", subKey)
	}
	delete(ws.subscriptions, subKey)
	close(sub.messages)
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
	var closeErr error
	ws.closeOnce.Do(func() {
		ws.mu.Lock()
		ws.reconnect = false
		close(ws.done)

		for _, sub := range ws.subscriptions {
			close(sub.messages)
		}
		ws.subscriptions = make(map[string]subscription)
		conn := ws.conn
		ws.conn = nil
		ws.mu.Unlock()

		if conn != nil {
			closeErr = conn.Close()
		}
	})
	return closeErr
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

	// gorilla/websocket permits one concurrent reader and one concurrent writer.
	// Subscription, login, and heartbeat messages all share this connection.
	ws.writeMu.Lock()
	defer ws.writeMu.Unlock()

	conn.SetWriteDeadline(time.Now().Add(writeTimeout))
	if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	return nil
}

func (ws *WSClient) readPump() {
	ws.mu.RLock()
	conn := ws.conn
	ws.mu.RUnlock()
	if conn == nil {
		return
	}

	defer func() {
		_ = conn.Close()
	}()

	for {
		select {
		case <-ws.done:
			return
		default:
		}

		conn.SetReadDeadline(time.Now().Add(readTimeout))
		_, message, err := conn.ReadMessage()
		if err != nil {
			ws.logger.Error("WebSocket read error", "error", err)
			ws.mu.RLock()
			reconnect := ws.reconnect
			ws.mu.RUnlock()
			if reconnect {
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

			ws.writeMu.Lock()
			conn.SetWriteDeadline(time.Now().Add(writeTimeout))
			if err := conn.WriteMessage(websocket.TextMessage, []byte("ping")); err != nil {
				ws.writeMu.Unlock()
				ws.logger.Error("WebSocket ping error", "error", err)
				return
			}
			ws.writeMu.Unlock()
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
		sub, exists := ws.subscriptions[subKey]
		if exists {
			select {
			case sub.messages <- message:
			default:
				ws.logger.Warn("Channel buffer full, dropping message", "channel", channel)
			}
		}
		ws.mu.RUnlock()
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
		err := ws.connect(ctx, false)
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
		subs := make([]subscription, 0, len(ws.subscriptions))
		for _, sub := range ws.subscriptions {
			subs = append(subs, sub)
		}
		ws.mu.RUnlock()

		for _, sub := range subs {
			args := make(map[string]interface{}, len(sub.args)+1)
			args["channel"] = sub.channel
			for key, value := range sub.args {
				args[key] = value
			}
			if err := ws.send(models.WSSubscribeRequest{Op: "subscribe", Args: []map[string]interface{}{args}}); err != nil {
				ws.logger.Error("Failed to restore subscription", "channel", sub.channel, "error", err)
			}
		}

		ws.logger.Info("Reconnected successfully")
		return
	}
}

func cloneArgs(args map[string]interface{}) map[string]interface{} {
	cloned := make(map[string]interface{}, len(args))
	for key, value := range args {
		cloned[key] = value
	}
	return cloned
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

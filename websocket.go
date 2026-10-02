package okx

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tigusigalpa/okx-go/models"
)

const (
	// WSPublicURL is the Global production WebSocket endpoint for public channels.
	WSPublicURL = "wss://ws.okx.com/ws/v5/public"
	// WSPrivateURL is the Global production WebSocket endpoint for private channels.
	WSPrivateURL = "wss://ws.okx.com/ws/v5/private"
	// WSBusinessURL is the Global production WebSocket endpoint for business channels.
	WSBusinessURL = "wss://ws.okx.com/ws/v5/business"
	// WSPublicSBEURL is the Global production WebSocket endpoint for public SBE channels.
	// The client only supports JSON protocol messages and does not decode SBE payloads.
	WSPublicSBEURL = "wss://ws.okx.com/ws/v5/public-sbe"

	// WSDemoPublicURL is the Global demo WebSocket endpoint for public channels.
	WSDemoPublicURL = "wss://wspap.okx.com/ws/v5/public"
	// WSDemoPrivateURL is the Global demo WebSocket endpoint for private channels.
	WSDemoPrivateURL = "wss://wspap.okx.com/ws/v5/private"
	// WSDemoBusinessURL is the Global demo WebSocket endpoint for business channels.
	WSDemoBusinessURL = "wss://wspap.okx.com/ws/v5/business"
	// WSDemoPublicSBEURL is the Global demo WebSocket endpoint for public SBE channels.
	WSDemoPublicSBEURL = "wss://wspap.okx.com/ws/v5/public-sbe"

	// WSEEAPublicURL is the EEA production WebSocket endpoint for public channels.
	WSEEAPublicURL = "wss://wseea.okx.com/ws/v5/public"
	// WSEEAPrivateURL is the EEA production WebSocket endpoint for private channels.
	WSEEAPrivateURL = "wss://wseea.okx.com/ws/v5/private"
	// WSEEABusinessURL is the EEA production WebSocket endpoint for business channels.
	WSEEABusinessURL = "wss://wseea.okx.com/ws/v5/business"
	// WSDemoEEAPublicURL is the EEA demo WebSocket endpoint for public channels.
	WSDemoEEAPublicURL = "wss://wseeapap.okx.com/ws/v5/public"
	// WSDemoEEAPrivateURL is the EEA demo WebSocket endpoint for private channels.
	WSDemoEEAPrivateURL = "wss://wseeapap.okx.com/ws/v5/private"
	// WSDemoEEABusinessURL is the EEA demo WebSocket endpoint for business channels.
	WSDemoEEABusinessURL = "wss://wseeapap.okx.com/ws/v5/business"

	// WSUSPublicURL is the United States production WebSocket endpoint for public channels.
	WSUSPublicURL = "wss://wsus.okx.com/ws/v5/public"
	// WSUSPrivateURL is the United States production WebSocket endpoint for private channels.
	WSUSPrivateURL = "wss://wsus.okx.com/ws/v5/private"
	// WSUSBusinessURL is the United States production WebSocket endpoint for business channels.
	WSUSBusinessURL = "wss://wsus.okx.com/ws/v5/business"
	// WSDemoUSPublicURL is the United States demo WebSocket endpoint for public channels.
	WSDemoUSPublicURL = "wss://wsuspap.okx.com/ws/v5/public"
	// WSDemoUSPrivateURL is the United States demo WebSocket endpoint for private channels.
	WSDemoUSPrivateURL = "wss://wsuspap.okx.com/ws/v5/private"
	// WSDemoUSBusinessURL is the United States demo WebSocket endpoint for business channels.
	WSDemoUSBusinessURL = "wss://wsuspap.okx.com/ws/v5/business"

	defaultWSSubscriptionBuffer   = 100
	defaultWSEventBuffer          = 100
	defaultWSPingInterval         = 25 * time.Second
	defaultWSPongTimeout          = 30 * time.Second
	defaultWSWriteTimeout         = 10 * time.Second
	defaultWSReadTimeout          = 60 * time.Second
	defaultWSReconnectBackoff     = time.Second
	defaultWSReconnectMaxBackoff  = 60 * time.Second
	defaultWSReconnectDialTimeout = 10 * time.Second
)

var (
	// ErrWSClientClosed is returned when an operation is attempted after Close.
	ErrWSClientClosed = errors.New("okx: WebSocket client is closed")
	// ErrWSInvalidSubscriptionBuffer is returned for a non-positive subscription buffer size.
	ErrWSInvalidSubscriptionBuffer = errors.New("okx: WebSocket subscription buffer size must be positive")
	// ErrWSLoginInProgress is returned when Login is already awaiting an OKX response.
	ErrWSLoginInProgress = errors.New("okx: WebSocket login is already in progress")
	// ErrWSAuthenticationFailed identifies a rejected WebSocket login.
	ErrWSAuthenticationFailed = errors.New("okx: WebSocket authentication failed")
	// ErrWSSubscriptionRejected identifies an OKX subscription rejection.
	ErrWSSubscriptionRejected = errors.New("okx: WebSocket subscription rejected")
	// ErrWSMessageDropped identifies a message dropped because a subscription buffer is full.
	ErrWSMessageDropped = errors.New("okx: WebSocket message dropped")
	// ErrWSSequenceGap identifies a discontinuity in an incremental order-book stream.
	ErrWSSequenceGap = errors.New("okx: WebSocket order-book sequence gap")
)

// WSRegion identifies an OKX WebSocket deployment.
type WSRegion string

const (
	// WSRegionGlobal is the Global OKX WebSocket deployment.
	WSRegionGlobal WSRegion = "global"
	// WSRegionEEA is the European Economic Area OKX WebSocket deployment.
	WSRegionEEA WSRegion = "eea"
	// WSRegionUS is the United States OKX WebSocket deployment.
	WSRegionUS WSRegion = "us"
	// WSRegionTR is Türkiye. OKX currently uses the Global WebSocket hosts for this region.
	WSRegionTR WSRegion = "tr"
)

// WSService identifies an OKX JSON WebSocket service.
type WSService string

const (
	// WSServicePublic is the public market-data WebSocket service.
	WSServicePublic WSService = "public"
	// WSServicePrivate is the authenticated account-data WebSocket service.
	WSServicePrivate WSService = "private"
	// WSServiceBusiness is the business WebSocket service, including candles.
	WSServiceBusiness WSService = "business"
)

// WSEventType identifies a WebSocket lifecycle or protocol event.
type WSEventType string

const (
	// WSEventConnected reports that a connection was established.
	WSEventConnected WSEventType = "connected"
	// WSEventDisconnected reports that the active connection failed.
	WSEventDisconnected WSEventType = "disconnected"
	// WSEventReconnecting reports a reconnect attempt.
	WSEventReconnecting WSEventType = "reconnecting"
	// WSEventReconnected reports that connection state was restored.
	WSEventReconnected WSEventType = "reconnected"
	// WSEventAuthenticationSucceeded reports an acknowledged OKX login.
	WSEventAuthenticationSucceeded WSEventType = "authentication_succeeded"
	// WSEventAuthenticationFailed reports a rejected OKX login.
	WSEventAuthenticationFailed WSEventType = "authentication_failed"
	// WSEventSubscriptionAcknowledged reports an OKX subscribe acknowledgement.
	WSEventSubscriptionAcknowledged WSEventType = "subscription_acknowledged"
	// WSEventSubscriptionRejected reports an OKX subscription rejection.
	WSEventSubscriptionRejected WSEventType = "subscription_rejected"
	// WSEventSubscriptionRestoreFailed reports a failed write while restoring a subscription.
	WSEventSubscriptionRestoreFailed WSEventType = "subscription_restore_failed"
	// WSEventMessageDropped reports a message dropped because its subscription buffer is full.
	WSEventMessageDropped WSEventType = "message_dropped"
	// WSEventNotice reports an OKX protocol notice, such as an upcoming service disconnect.
	WSEventNotice WSEventType = "notice"
)

// WSEvent is an observable WebSocket lifecycle, protocol, or backpressure event.
type WSEvent struct {
	Type    WSEventType
	Channel string
	Args    map[string]interface{}
	Code    string
	Message string
	Err     error
	Dropped uint64
}

// WSDrop describes a message dropped because a subscription consumer is too slow.
type WSDrop struct {
	Channel string
	Args    map[string]interface{}
	Message []byte
	Count   uint64
}

// WSDropHandler receives observable subscription buffer overflows.
// It is called synchronously by the WebSocket reader and should return quickly.
type WSDropHandler func(WSDrop)

// WSServerError is an error response sent by the OKX WebSocket API.
type WSServerError struct {
	Event   string
	Code    string
	Message string
}

// Error returns the API error in a human-readable form.
func (e *WSServerError) Error() string {
	if e.Code == "" {
		return fmt.Sprintf("okx: WebSocket %s: %s", e.Event, e.Message)
	}
	return fmt.Sprintf("okx: WebSocket %s failed (code %s): %s", e.Event, e.Code, e.Message)
}

// Unwrap returns a sentinel error for recognized OKX WebSocket failures.
func (e *WSServerError) Unwrap() error {
	switch e.Event {
	case "login":
		return ErrWSAuthenticationFailed
	case "subscribe":
		return ErrWSSubscriptionRejected
	default:
		return nil
	}
}

// WebSocketURL returns the supported JSON WebSocket endpoint for a region and service.
// Türkiye currently uses the Global WebSocket hosts according to OKX documentation.
func WebSocketURL(region WSRegion, service WSService, demo bool) (string, error) {
	var host string
	switch region {
	case WSRegionGlobal, WSRegionTR:
		if demo {
			host = "wspap.okx.com"
		} else {
			host = "ws.okx.com"
		}
	case WSRegionEEA:
		if demo {
			host = "wseeapap.okx.com"
		} else {
			host = "wseea.okx.com"
		}
	case WSRegionUS:
		if demo {
			host = "wsuspap.okx.com"
		} else {
			host = "wsus.okx.com"
		}
	default:
		return "", fmt.Errorf("okx: unsupported WebSocket region %q", region)
	}

	switch service {
	case WSServicePublic, WSServicePrivate, WSServiceBusiness:
		return "wss://" + host + "/ws/v5/" + string(service), nil
	default:
		return "", fmt.Errorf("okx: unsupported WebSocket service %q", service)
	}
}

// ValidateOrderBookSequence validates continuity between two incremental order-book updates.
// Snapshot messages have PrevSeqID == -1. Empty keepalive updates and maintenance resets are
// valid when the next PrevSeqID matches the previous SeqID, even if the next SeqID is lower.
func ValidateOrderBookSequence(previous, next models.OrderBook) error {
	if next.PrevSeqID == -1 || next.PrevSeqID == previous.SeqID {
		return nil
	}
	return fmt.Errorf("%w: previous seqId=%d, next prevSeqId=%d", ErrWSSequenceGap, previous.SeqID, next.PrevSeqID)
}

// WSClient manages an OKX WebSocket connection and its subscriptions.
type WSClient struct {
	apiKey     string
	secretKey  string
	passphrase string
	url        string
	conn       *websocket.Conn
	isDemo     bool
	logger     Logger

	mu                    sync.RWMutex
	writeMu               sync.Mutex
	subscriptions         map[string]*subscription
	done                  chan struct{}
	closeOnce             sync.Once
	closed                bool
	reconnecting          bool
	authenticated         bool
	restoreAuthentication bool
	loginWaiter           chan error
	lastReceived          time.Time
	lastPing              time.Time
	awaitingPong          bool
	heartbeatStarted      bool
	configurationErr      error
	subscriptionBuffer    int
	dropHandler           WSDropHandler
	events                chan WSEvent

	pingInterval         time.Duration
	pongTimeout          time.Duration
	writeTimeout         time.Duration
	readTimeout          time.Duration
	reconnectBackoff     time.Duration
	reconnectMaxBackoff  time.Duration
	reconnectDialTimeout time.Duration
}

type subscription struct {
	channel  string
	args     map[string]interface{}
	messages chan []byte
	dropped  atomic.Uint64
}

// WSOption configures a WSClient.
type WSOption func(*WSClient)

// WithWSDemo routes supported production endpoints to their demo equivalents.
func WithWSDemo() WSOption {
	return func(ws *WSClient) {
		ws.isDemo = true
	}
}

// WithWSLogger sets the logger used by a WSClient.
func WithWSLogger(logger Logger) WSOption {
	return func(ws *WSClient) {
		if logger != nil {
			ws.logger = logger
		}
	}
}

// WithWSSubscriptionBuffer sets the bounded per-subscription message buffer size.
// A non-positive value makes Connect and Subscribe return ErrWSInvalidSubscriptionBuffer.
func WithWSSubscriptionBuffer(size int) WSOption {
	return func(ws *WSClient) {
		if size <= 0 {
			ws.configurationErr = fmt.Errorf("%w: %d", ErrWSInvalidSubscriptionBuffer, size)
			return
		}
		ws.subscriptionBuffer = size
	}
}

// WithWSDropHandler sets a callback for observable subscription buffer overflows.
func WithWSDropHandler(handler WSDropHandler) WSOption {
	return func(ws *WSClient) {
		ws.dropHandler = handler
	}
}

// NewWSClient creates a WebSocket client with the provided credentials and endpoint.
func NewWSClient(apiKey, secretKey, passphrase, url string, opts ...WSOption) *WSClient {
	ws := &WSClient{
		apiKey:               apiKey,
		secretKey:            secretKey,
		passphrase:           passphrase,
		url:                  url,
		logger:               &noopLogger{},
		subscriptions:        make(map[string]*subscription),
		done:                 make(chan struct{}),
		subscriptionBuffer:   defaultWSSubscriptionBuffer,
		events:               make(chan WSEvent, defaultWSEventBuffer),
		pingInterval:         defaultWSPingInterval,
		pongTimeout:          defaultWSPongTimeout,
		writeTimeout:         defaultWSWriteTimeout,
		readTimeout:          defaultWSReadTimeout,
		reconnectBackoff:     defaultWSReconnectBackoff,
		reconnectMaxBackoff:  defaultWSReconnectMaxBackoff,
		reconnectDialTimeout: defaultWSReconnectDialTimeout,
	}

	for _, opt := range opts {
		opt(ws)
	}
	if ws.isDemo {
		ws.url = demoWebSocketURL(ws.url)
	}

	return ws
}

// Events returns lifecycle, protocol, and backpressure events. The bounded channel is closed
// when Close returns and should be drained by consumers that need complete observability.
func (ws *WSClient) Events() <-chan WSEvent {
	return ws.events
}

// IsAuthenticated reports whether the current connection has received a successful
// OKX login acknowledgement. It is false while a login is pending or reconnecting.
func (ws *WSClient) IsAuthenticated() bool {
	ws.mu.RLock()
	defer ws.mu.RUnlock()
	return ws.authenticated
}

func demoWebSocketURL(endpoint string) string {
	demoEndpoints := map[string]string{
		WSPublicURL:      WSDemoPublicURL,
		WSPrivateURL:     WSDemoPrivateURL,
		WSBusinessURL:    WSDemoBusinessURL,
		WSPublicSBEURL:   WSDemoPublicSBEURL,
		WSEEAPublicURL:   WSDemoEEAPublicURL,
		WSEEAPrivateURL:  WSDemoEEAPrivateURL,
		WSEEABusinessURL: WSDemoEEABusinessURL,
		WSUSPublicURL:    WSDemoUSPublicURL,
		WSUSPrivateURL:   WSDemoUSPrivateURL,
		WSUSBusinessURL:  WSDemoUSBusinessURL,
	}
	if demo, ok := demoEndpoints[endpoint]; ok {
		return demo
	}
	return endpoint
}

// Connect establishes the WebSocket connection and starts the reader and heartbeat pumps.
func (ws *WSClient) Connect(ctx context.Context) error {
	if err := ws.validateOperation(ctx); err != nil {
		return err
	}
	return ws.connect(ctx, true)
}

func (ws *WSClient) connect(ctx context.Context, startHeartbeat bool) error {
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.DialContext(ctx, ws.url, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	ws.mu.Lock()
	if ws.closed {
		ws.mu.Unlock()
		_ = conn.Close()
		return ErrWSClientClosed
	}
	previous := ws.conn
	ws.conn = conn
	ws.authenticated = false
	ws.lastReceived = time.Now()
	ws.lastPing = time.Time{}
	ws.awaitingPong = false
	if startHeartbeat && !ws.heartbeatStarted {
		ws.heartbeatStarted = true
		go ws.heartbeatPump()
	}
	ws.mu.Unlock()

	if previous != nil && previous != conn {
		_ = previous.Close()
	}

	ws.logger.Info("WebSocket connected", "url", ws.url)
	ws.emit(WSEvent{Type: WSEventConnected})
	go ws.readPump(conn)

	return nil
}

// Login authenticates the client and waits for the OKX login acknowledgement.
func (ws *WSClient) Login(ctx context.Context) error {
	if err := ws.validateOperation(ctx); err != nil {
		return err
	}

	ws.mu.Lock()
	if ws.conn == nil {
		ws.mu.Unlock()
		return errors.New("WebSocket not connected")
	}
	if ws.authenticated {
		ws.mu.Unlock()
		return nil
	}
	if ws.loginWaiter != nil {
		ws.mu.Unlock()
		return ErrWSLoginInProgress
	}
	waiter := make(chan error, 1)
	ws.loginWaiter = waiter
	ws.mu.Unlock()

	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	message := timestamp + "GET" + "/users/self/verify"
	h := hmac.New(sha256.New, []byte(ws.secretKey))
	_, _ = h.Write([]byte(message))

	loginReq := models.WSLoginRequest{
		Op: "login",
		Args: []models.WSLoginArgs{{
			APIKey:     ws.apiKey,
			Passphrase: ws.passphrase,
			Timestamp:  timestamp,
			Sign:       base64.StdEncoding.EncodeToString(h.Sum(nil)),
		}},
	}

	if err := ws.send(loginReq); err != nil {
		ws.clearLoginWaiter(waiter)
		return fmt.Errorf("failed to send login request: %w", err)
	}

	select {
	case err := <-waiter:
		return err
	case <-ctx.Done():
		ws.clearLoginWaiter(waiter)
		return ctx.Err()
	case <-ws.done:
		return ErrWSClientClosed
	}
}

// Subscribe registers a generic subscription and returns its raw JSON message stream.
// A successful return means the request was written to the socket; acknowledgement or rejection
// from OKX is reported asynchronously through Events.
func (ws *WSClient) Subscribe(ctx context.Context, channel string, args map[string]interface{}) (<-chan []byte, error) {
	if err := ws.validateOperation(ctx); err != nil {
		return nil, err
	}
	subKey := ws.makeSubKey(channel, args)

	ws.mu.Lock()
	if _, exists := ws.subscriptions[subKey]; exists {
		ws.mu.Unlock()
		return nil, fmt.Errorf("already subscribed to %s", subKey)
	}

	ch := make(chan []byte, ws.subscriptionBuffer)
	ws.subscriptions[subKey] = &subscription{channel: channel, args: cloneArgs(args), messages: ch}
	ws.mu.Unlock()

	subArgs := subscriptionArgs(channel, args)
	if err := ws.send(models.WSSubscribeRequest{Op: "subscribe", Args: []map[string]interface{}{subArgs}}); err != nil {
		ws.mu.Lock()
		if sub, exists := ws.subscriptions[subKey]; exists && sub.messages == ch {
			delete(ws.subscriptions, subKey)
			close(ch)
		}
		ws.mu.Unlock()
		return nil, fmt.Errorf("failed to send subscribe request: %w", err)
	}

	ws.logger.Info("Subscribed to channel", "channel", channel, "args", args)
	return ch, nil
}

// Unsubscribe removes a subscription from a channel.
func (ws *WSClient) Unsubscribe(channel string, args map[string]interface{}) error {
	if err := ws.validateOperation(context.Background()); err != nil {
		return err
	}
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

	if err := ws.send(models.WSUnsubscribeRequest{Op: "unsubscribe", Args: []map[string]interface{}{subscriptionArgs(channel, args)}}); err != nil {
		return fmt.Errorf("failed to send unsubscribe request: %w", err)
	}

	ws.logger.Info("Unsubscribed from channel", "channel", channel, "args", args)
	return nil
}

// Close stops all background activity, closes subscription streams, and closes the connection.
// It is safe to call multiple times and interrupts a reconnect backoff immediately.
func (ws *WSClient) Close() error {
	var closeErr error
	ws.closeOnce.Do(func() {
		ws.mu.Lock()
		ws.closed = true
		close(ws.done)
		waiter := ws.loginWaiter
		ws.loginWaiter = nil
		for _, sub := range ws.subscriptions {
			close(sub.messages)
		}
		ws.subscriptions = make(map[string]*subscription)
		close(ws.events)
		conn := ws.conn
		ws.conn = nil
		ws.mu.Unlock()
		if waiter != nil {
			waiter <- ErrWSClientClosed
		}

		if conn != nil {
			if err := conn.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
				closeErr = err
			}
		}
	})
	return closeErr
}

func (ws *WSClient) send(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	ws.mu.RLock()
	conn := ws.conn
	closed := ws.closed
	ws.mu.RUnlock()
	if closed {
		return ErrWSClientClosed
	}
	if conn == nil {
		return errors.New("WebSocket not connected")
	}
	if err := ws.write(conn, data); err != nil {
		ws.triggerReconnect(conn, err)
		return err
	}
	return nil
}

func (ws *WSClient) write(conn *websocket.Conn, data []byte) error {
	ws.writeMu.Lock()
	defer ws.writeMu.Unlock()

	if err := conn.SetWriteDeadline(time.Now().Add(ws.writeTimeout)); err != nil {
		return fmt.Errorf("failed to set write deadline: %w", err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}
	return nil
}

func (ws *WSClient) readPump(conn *websocket.Conn) {
	defer func() { _ = conn.Close() }()

	for {
		if err := conn.SetReadDeadline(time.Now().Add(ws.readTimeout)); err != nil {
			ws.triggerReconnect(conn, fmt.Errorf("failed to set read deadline: %w", err))
			return
		}

		_, message, err := conn.ReadMessage()
		if err != nil {
			ws.triggerReconnect(conn, fmt.Errorf("WebSocket read error: %w", err))
			return
		}

		ws.noteActivity(conn)
		if string(message) == "pong" {
			ws.logger.Debug("Received WebSocket pong")
			continue
		}
		ws.handleMessageFromConn(conn, message)
	}
}

func (ws *WSClient) heartbeatPump() {
	tickEvery := ws.pingInterval / 5
	if tickEvery <= 0 || tickEvery > 5*time.Second {
		tickEvery = 5 * time.Second
	}
	ticker := time.NewTicker(tickEvery)
	defer ticker.Stop()

	for {
		select {
		case <-ws.done:
			return
		case now := <-ticker.C:
			ws.heartbeat(now)
		}
	}
}

func (ws *WSClient) heartbeat(now time.Time) {
	ws.mu.Lock()
	if ws.closed || ws.conn == nil {
		ws.mu.Unlock()
		return
	}
	conn := ws.conn
	if ws.awaitingPong {
		if now.Sub(ws.lastPing) >= ws.pongTimeout {
			ws.awaitingPong = false
			ws.mu.Unlock()
			ws.triggerReconnect(conn, errors.New("WebSocket pong timeout"))
			return
		}
		ws.mu.Unlock()
		return
	}
	if now.Sub(ws.lastReceived) < ws.pingInterval {
		ws.mu.Unlock()
		return
	}
	ws.awaitingPong = true
	ws.lastPing = now
	ws.mu.Unlock()

	if err := ws.write(conn, []byte("ping")); err != nil {
		ws.triggerReconnect(conn, fmt.Errorf("WebSocket ping error: %w", err))
		return
	}
	ws.logger.Debug("Sent WebSocket ping")
}

func (ws *WSClient) handleMessage(message []byte) {
	ws.handleMessageFromConn(nil, message)
}

func (ws *WSClient) handleMessageFromConn(source *websocket.Conn, message []byte) {
	if source != nil && !ws.isActiveConnection(source) {
		return
	}

	var response struct {
		Event string                 `json:"event"`
		Code  string                 `json:"code"`
		Msg   string                 `json:"msg"`
		Arg   map[string]interface{} `json:"arg"`
	}
	if err := json.Unmarshal(message, &response); err != nil {
		ws.logger.Error("Failed to unmarshal WebSocket message", "error", err)
		ws.emit(WSEvent{Type: WSEventNotice, Err: fmt.Errorf("failed to decode WebSocket message: %w", err)})
		return
	}

	switch response.Event {
	case "error":
		ws.handleServerError(source, response.Arg, response.Code, response.Msg)
		return
	case "login":
		ws.handleLoginResponse(source, response.Code, response.Msg)
		return
	case "subscribe":
		ws.logger.Debug("WebSocket subscription acknowledged", "arg", response.Arg)
		ws.emit(WSEvent{Type: WSEventSubscriptionAcknowledged, Channel: channelFromArg(response.Arg), Args: cloneArgs(response.Arg)})
		return
	case "unsubscribe":
		ws.logger.Debug("WebSocket unsubscription acknowledged", "arg", response.Arg)
		return
	case "notice":
		err := &WSServerError{Event: response.Event, Code: response.Code, Message: response.Msg}
		ws.emit(WSEvent{Type: WSEventNotice, Code: response.Code, Message: response.Msg, Err: err})
		ws.logger.Warn("WebSocket notice", "code", response.Code, "msg", response.Msg)
		if response.Code == "64008" {
			conn := source
			if conn == nil {
				ws.mu.RLock()
				conn = ws.conn
				ws.mu.RUnlock()
			}
			if conn != nil {
				ws.triggerReconnect(conn, err)
			}
		}
		return
	}

	channel := channelFromArg(response.Arg)
	if channel == "" {
		ws.logger.Warn("No channel in WebSocket message arg")
		return
	}
	subKey := ws.makeSubKeyFromArg(channel, response.Arg)

	ws.mu.RLock()
	sub, exists := ws.subscriptions[subKey]
	if !exists {
		ws.mu.RUnlock()
		return
	}
	select {
	case sub.messages <- message:
		ws.mu.RUnlock()
	default:
		count := sub.dropped.Add(1)
		drop := WSDrop{Channel: sub.channel, Args: cloneArgs(sub.args), Message: append([]byte(nil), message...), Count: count}
		ws.mu.RUnlock()
		err := fmt.Errorf("%w: channel=%s count=%d", ErrWSMessageDropped, drop.Channel, drop.Count)
		ws.logger.Warn("WebSocket message dropped", "channel", drop.Channel, "count", drop.Count)
		ws.emit(WSEvent{Type: WSEventMessageDropped, Channel: drop.Channel, Args: cloneArgs(drop.Args), Err: err, Dropped: drop.Count})
		if ws.dropHandler != nil {
			ws.dropHandler(drop)
		}
	}
}

func (ws *WSClient) handleServerError(source *websocket.Conn, arg map[string]interface{}, code, message string) {
	channel := channelFromArg(arg)
	err := &WSServerError{Event: "subscribe", Code: code, Message: message}
	eventType := WSEventSubscriptionRejected

	ws.mu.RLock()
	active := !ws.closed && (source == nil || ws.conn == source)
	loginPending := ws.loginWaiter != nil
	ws.mu.RUnlock()
	if !active {
		return
	}
	if loginPending && channel == "" {
		ws.handleLoginResponse(source, code, message)
		return
	}

	ws.logger.Error("WebSocket error event", "code", code, "msg", message)
	ws.emit(WSEvent{Type: eventType, Channel: channel, Args: cloneArgs(arg), Code: code, Message: message, Err: err})
}

func (ws *WSClient) handleLoginResponse(source *websocket.Conn, code, message string) {
	ws.mu.Lock()
	if ws.closed || (source != nil && ws.conn != source) {
		ws.mu.Unlock()
		return
	}
	waiter := ws.loginWaiter
	ws.loginWaiter = nil
	if code == "0" {
		ws.authenticated = true
		ws.restoreAuthentication = true
	} else {
		ws.authenticated = false
	}
	ws.mu.Unlock()

	if code == "0" {
		ws.logger.Info("WebSocket authenticated")
		ws.emit(WSEvent{Type: WSEventAuthenticationSucceeded})
		if waiter != nil {
			waiter <- nil
		}
		return
	}

	err := &WSServerError{Event: "login", Code: code, Message: message}
	ws.logger.Error("WebSocket login failed", "code", code, "msg", message)
	ws.emit(WSEvent{Type: WSEventAuthenticationFailed, Code: code, Message: message, Err: err})
	if waiter != nil {
		waiter <- err
	}
}

func (ws *WSClient) triggerReconnect(conn *websocket.Conn, cause error) {
	ws.mu.Lock()
	if ws.closed || ws.reconnecting || ws.conn != conn {
		ws.mu.Unlock()
		return
	}
	ws.reconnecting = true
	ws.authenticated = false
	waiter := ws.loginWaiter
	ws.loginWaiter = nil
	ws.mu.Unlock()
	if waiter != nil {
		waiter <- fmt.Errorf("WebSocket disconnected during login: %w", cause)
	}

	ws.logger.Error("WebSocket disconnected", "error", cause)
	ws.emit(WSEvent{Type: WSEventDisconnected, Err: cause})
	go ws.reconnectLoop()
}

func (ws *WSClient) reconnectLoop() {
	defer func() {
		ws.mu.Lock()
		ws.reconnecting = false
		ws.mu.Unlock()
	}()

	backoff := ws.reconnectBackoff
	for {
		if !ws.waitForReconnect(backoff) {
			return
		}
		ws.emit(WSEvent{Type: WSEventReconnecting})
		ws.logger.Info("Attempting WebSocket reconnect", "backoff", backoff)

		ctx, cancel := context.WithTimeout(context.Background(), ws.reconnectDialTimeout)
		err := ws.connect(ctx, false)
		cancel()
		if err != nil {
			ws.logger.Error("WebSocket reconnect failed", "error", err)
			backoff = nextBackoff(backoff, ws.reconnectMaxBackoff)
			continue
		}

		if err := ws.restoreConnectionState(); err != nil {
			ws.logger.Error("WebSocket connection state restoration failed", "error", err)
			ws.mu.RLock()
			conn := ws.conn
			ws.mu.RUnlock()
			if conn != nil {
				_ = conn.Close()
			}
			backoff = nextBackoff(backoff, ws.reconnectMaxBackoff)
			continue
		}

		ws.logger.Info("WebSocket reconnected")
		ws.emit(WSEvent{Type: WSEventReconnected})
		return
	}
}

func (ws *WSClient) restoreConnectionState() error {
	ws.mu.RLock()
	restoreAuthentication := ws.restoreAuthentication
	ws.mu.RUnlock()
	if restoreAuthentication {
		ctx, cancel := context.WithTimeout(context.Background(), ws.reconnectDialTimeout)
		err := ws.Login(ctx)
		cancel()
		if err != nil {
			return fmt.Errorf("failed to restore authentication: %w", err)
		}
	}

	ws.mu.RLock()
	subs := make([]*subscription, 0, len(ws.subscriptions))
	for _, sub := range ws.subscriptions {
		subs = append(subs, sub)
	}
	ws.mu.RUnlock()

	for _, sub := range subs {
		args := subscriptionArgs(sub.channel, sub.args)
		if err := ws.send(models.WSSubscribeRequest{Op: "subscribe", Args: []map[string]interface{}{args}}); err != nil {
			ws.logger.Error("Failed to restore WebSocket subscription", "channel", sub.channel, "error", err)
			ws.emit(WSEvent{Type: WSEventSubscriptionRestoreFailed, Channel: sub.channel, Args: cloneArgs(sub.args), Err: err})
			return fmt.Errorf("failed to restore %s subscription: %w", sub.channel, err)
		}
	}
	return nil
}

func (ws *WSClient) waitForReconnect(backoff time.Duration) bool {
	timer := time.NewTimer(backoff)
	defer timer.Stop()
	select {
	case <-ws.done:
		return false
	case <-timer.C:
		return true
	}
}

func (ws *WSClient) validateOperation(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	ws.mu.RLock()
	defer ws.mu.RUnlock()
	if ws.closed {
		return ErrWSClientClosed
	}
	return ws.configurationErr
}

func (ws *WSClient) clearLoginWaiter(waiter chan error) {
	ws.mu.Lock()
	if ws.loginWaiter == waiter {
		ws.loginWaiter = nil
	}
	ws.mu.Unlock()
}

func (ws *WSClient) noteActivity(conn *websocket.Conn) {
	ws.mu.Lock()
	if ws.conn == conn {
		ws.lastReceived = time.Now()
		ws.awaitingPong = false
	}
	ws.mu.Unlock()
}

func (ws *WSClient) isActiveConnection(conn *websocket.Conn) bool {
	ws.mu.RLock()
	defer ws.mu.RUnlock()
	return !ws.closed && ws.conn == conn
}

func (ws *WSClient) emit(event WSEvent) {
	ws.mu.RLock()
	if ws.closed {
		ws.mu.RUnlock()
		return
	}
	dropped := false
	select {
	case ws.events <- event:
	default:
		dropped = true
	}
	ws.mu.RUnlock()
	if dropped {
		ws.logger.Warn("WebSocket event buffer full", "type", event.Type)
	}
}

func cloneArgs(args map[string]interface{}) map[string]interface{} {
	cloned := make(map[string]interface{}, len(args))
	for key, value := range args {
		cloned[key] = value
	}
	return cloned
}

func subscriptionArgs(channel string, args map[string]interface{}) map[string]interface{} {
	subArgs := make(map[string]interface{}, len(args)+1)
	subArgs["channel"] = channel
	for key, value := range args {
		subArgs[key] = value
	}
	return subArgs
}

func channelFromArg(arg map[string]interface{}) string {
	channel, _ := arg["channel"].(string)
	return channel
}

func nextBackoff(current, maximum time.Duration) time.Duration {
	if current >= maximum/2 {
		return maximum
	}
	return current * 2
}

func (ws *WSClient) makeSubKey(channel string, args map[string]interface{}) string {
	key := channel
	known := map[string]struct{}{
		"instId":     {},
		"instType":   {},
		"ccy":        {},
		"instFamily": {},
	}
	if instID, ok := args["instId"].(string); ok {
		key += ":" + instID
	}
	if instType, ok := args["instType"].(string); ok {
		key += ":" + instType
	}
	if ccy, ok := args["ccy"].(string); ok {
		key += ":" + ccy
	}
	if instFamily, ok := args["instFamily"].(string); ok {
		key += ":" + instFamily
	}

	// Preserve the historic readable key for common arguments while ensuring that
	// subscriptions differentiated by newer OKX arguments (for example uly or uid)
	// cannot collide. encoding/json orders map keys deterministically.
	extra := make(map[string]interface{})
	for name, value := range args {
		if _, isKnown := known[name]; !isKnown {
			extra[name] = value
		}
	}
	if len(extra) > 0 {
		if encoded, err := json.Marshal(extra); err == nil {
			key += ":" + string(encoded)
		}
	}
	return key
}

func (ws *WSClient) makeSubKeyFromArg(channel string, arg map[string]interface{}) string {
	args := make(map[string]interface{}, len(arg))
	for key, value := range arg {
		if key != "channel" {
			args[key] = value
		}
	}
	return ws.makeSubKey(channel, args)
}

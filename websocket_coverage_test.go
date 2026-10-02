package okx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
	"github.com/tigusigalpa/okx-go/models"
)

func testWebSocketURL(server *httptest.Server) string {
	return "ws" + strings.TrimPrefix(server.URL, "http")
}

func waitForWSEvent(t *testing.T, events <-chan WSEvent, predicate func(WSEvent) bool) WSEvent {
	t.Helper()
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()
	for {
		select {
		case event := <-events:
			if predicate(event) {
				return event
			}
		case <-timer.C:
			t.Fatal("timed out waiting for WebSocket event")
		}
	}
}

func TestWSClientSendsLoginAndSubscriptionMessages(t *testing.T) {
	upgrader := websocket.Upgrader{}
	messages := make(chan []byte, 3)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			messages <- message

			var request struct {
				Op string `json:"op"`
			}
			if json.Unmarshal(message, &request) != nil {
				return
			}
			switch request.Op {
			case "login":
				_ = conn.WriteMessage(websocket.TextMessage, []byte(`{"event":"login","code":"0","msg":""}`))
			case "subscribe":
				_ = conn.WriteMessage(websocket.TextMessage, []byte(`{"event":"subscribe","arg":{"channel":"tickers","instId":"BTC-USDT"}}`))
			}
		}
	}))
	defer server.Close()

	client := NewWSClient("api-key", "secret", "passphrase", testWebSocketURL(server))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	require.NoError(t, client.Connect(ctx))
	defer client.Close()

	require.NoError(t, client.Login(ctx))
	require.True(t, client.IsAuthenticated())
	ch, err := client.Subscribe(ctx, "tickers", map[string]interface{}{"instId": "BTC-USDT"})
	require.NoError(t, err)
	require.NotNil(t, ch)
	require.NoError(t, client.Unsubscribe("tickers", map[string]interface{}{"instId": "BTC-USDT"}))

	for _, operation := range []string{"login", "subscribe", "unsubscribe"} {
		select {
		case message := <-messages:
			require.Contains(t, string(message), `"op":"`+operation+`"`)
		case <-ctx.Done():
			t.Fatalf("did not receive %s message", operation)
		}
	}

	waitForWSEvent(t, client.Events(), func(event WSEvent) bool {
		return event.Type == WSEventSubscriptionAcknowledged && event.Channel == "tickers"
	})
}

func TestWSClientLoginWaitsForAcknowledgement(t *testing.T) {
	upgrader := websocket.Upgrader{}
	loginReceived := make(chan struct{})
	releaseLogin := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		_, message, err := conn.ReadMessage()
		if err != nil || !strings.Contains(string(message), `"op":"login"`) {
			return
		}
		close(loginReceived)
		<-releaseLogin
		_ = conn.WriteMessage(websocket.TextMessage, []byte(`{"event":"login","code":"0"}`))
		_, _, _ = conn.ReadMessage()
	}))
	defer server.Close()

	client := NewWSClient("api-key", "secret", "passphrase", testWebSocketURL(server))
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	require.NoError(t, client.Connect(ctx))

	loginResult := make(chan error, 1)
	go func() { loginResult <- client.Login(ctx) }()
	select {
	case <-loginReceived:
		require.False(t, client.IsAuthenticated(), "authentication must wait for OKX acknowledgement")
	case <-ctx.Done():
		t.Fatal("server did not receive the login request")
	}
	close(releaseLogin)
	require.NoError(t, <-loginResult)
	require.True(t, client.IsAuthenticated())
	waitForWSEvent(t, client.Events(), func(event WSEvent) bool {
		return event.Type == WSEventAuthenticationSucceeded
	})
}

func TestWSClientReportsRejectedLogin(t *testing.T) {
	upgrader := websocket.Upgrader{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		_, _, err = conn.ReadMessage()
		if err == nil {
			_ = conn.WriteMessage(websocket.TextMessage, []byte(`{"event":"error","code":"50113","msg":"Invalid Sign"}`))
		}
	}))
	defer server.Close()

	client := NewWSClient("api-key", "secret", "passphrase", testWebSocketURL(server))
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	require.NoError(t, client.Connect(ctx))
	err := client.Login(ctx)
	require.ErrorIs(t, err, ErrWSAuthenticationFailed)
	require.False(t, client.IsAuthenticated())
	event := waitForWSEvent(t, client.Events(), func(event WSEvent) bool {
		return event.Type == WSEventAuthenticationFailed
	})
	require.ErrorIs(t, event.Err, ErrWSAuthenticationFailed)
}

func TestWSClientCloseInterruptsPendingLogin(t *testing.T) {
	upgrader := websocket.Upgrader{}
	loginReceived := make(chan struct{})
	finish := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		if _, _, err := conn.ReadMessage(); err == nil {
			close(loginReceived)
			<-finish
		}
	}))
	defer server.Close()
	defer close(finish)

	client := NewWSClient("api-key", "secret", "passphrase", testWebSocketURL(server))
	require.NoError(t, client.Connect(context.Background()))
	loginResult := make(chan error, 1)
	go func() { loginResult <- client.Login(context.Background()) }()
	select {
	case <-loginReceived:
	case <-time.After(time.Second):
		t.Fatal("server did not receive the login request")
	}
	require.NoError(t, client.Close())
	select {
	case err := <-loginResult:
		require.ErrorIs(t, err, ErrWSClientClosed)
	case <-time.After(time.Second):
		t.Fatal("Close did not interrupt Login")
	}
}

func TestWSClientHandleMessageDeliversSubscriptionData(t *testing.T) {
	client := NewWSClient("", "", "", WSPublicURL)
	defer client.Close()

	messages := make(chan []byte, 1)
	client.subscriptions["tickers:BTC-USDT"] = &subscription{
		channel:  "tickers",
		args:     map[string]interface{}{"instId": "BTC-USDT"},
		messages: messages,
	}
	payload := []byte(`{"arg":{"channel":"tickers","instId":"BTC-USDT"},"data":[{"last":"1"}]}`)
	client.handleMessage(payload)

	select {
	case received := <-messages:
		require.JSONEq(t, string(payload), string(received))
	case <-time.After(time.Second):
		t.Fatal("subscription message was not delivered")
	}

	client.handleMessage([]byte(`not json`))
	client.handleMessage([]byte(`{"event":"error","code":"1","msg":"failed"}`))
	event := waitForWSEvent(t, client.Events(), func(event WSEvent) bool {
		return event.Type == WSEventSubscriptionRejected
	})
	require.ErrorIs(t, event.Err, ErrWSSubscriptionRejected)
}

func TestWSClientIgnoresLateMessagesFromReplacedConnection(t *testing.T) {
	client := NewWSClient("", "", "", WSPublicURL)
	oldConn := &websocket.Conn{}
	currentConn := &websocket.Conn{}
	client.conn = currentConn
	client.loginWaiter = make(chan error, 1)

	client.handleMessageFromConn(oldConn, []byte(`{"event":"login","code":"0"}`))
	require.False(t, client.IsAuthenticated())
	require.NotNil(t, client.loginWaiter)
}

func TestWSClientReconnectRestoresSubscriptions(t *testing.T) {
	upgrader := websocket.Upgrader{}
	var connectionCount atomic.Int32
	resubscribed := make(chan struct{}, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		connection := connectionCount.Add(1)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			if !strings.Contains(string(message), `"op":"subscribe"`) {
				continue
			}
			if connection == 1 {
				return
			}
			resubscribed <- struct{}{}
			_ = conn.WriteMessage(websocket.TextMessage, []byte(`{"arg":{"channel":"tickers","instId":"BTC-USDT"},"data":[{"instId":"BTC-USDT","last":"42"}]}`))
			_, _, _ = conn.ReadMessage()
		}
	}))
	defer server.Close()

	client := NewWSClient("", "", "", testWebSocketURL(server))
	client.reconnectBackoff = 10 * time.Millisecond
	client.reconnectMaxBackoff = 20 * time.Millisecond
	client.reconnectDialTimeout = time.Second
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	require.NoError(t, client.Connect(ctx))
	stream, err := client.Subscribe(ctx, "tickers", map[string]interface{}{"instId": "BTC-USDT"})
	require.NoError(t, err)

	select {
	case <-resubscribed:
	case <-time.After(2 * time.Second):
		t.Fatal("reconnect did not restore the subscription")
	}
	select {
	case message := <-stream:
		require.Contains(t, string(message), `"last":"42"`)
	case <-time.After(time.Second):
		t.Fatal("restored stream did not receive market data")
	}
	waitForWSEvent(t, client.Events(), func(event WSEvent) bool {
		return event.Type == WSEventReconnected
	})
}

func TestWSClientReconnectRestoresAcknowledgedAuthentication(t *testing.T) {
	upgrader := websocket.Upgrader{}
	var connections atomic.Int32
	closeFirst := make(chan struct{})
	finish := make(chan struct{})
	restoredLogin := make(chan struct{}, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		connection := connections.Add(1)
		_, message, err := conn.ReadMessage()
		if err != nil || !strings.Contains(string(message), `"op":"login"`) {
			return
		}
		_ = conn.WriteMessage(websocket.TextMessage, []byte(`{"event":"login","code":"0"}`))
		if connection == 1 {
			<-closeFirst
			return
		}
		restoredLogin <- struct{}{}
		<-finish
	}))
	defer server.Close()
	defer close(finish)

	client := NewWSClient("api-key", "secret", "passphrase", testWebSocketURL(server))
	client.reconnectBackoff = 10 * time.Millisecond
	client.reconnectMaxBackoff = 20 * time.Millisecond
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	require.NoError(t, client.Connect(ctx))
	require.NoError(t, client.Login(ctx))
	require.True(t, client.IsAuthenticated())

	close(closeFirst)
	select {
	case <-restoredLogin:
	case <-time.After(2 * time.Second):
		t.Fatal("client did not restore a confirmed login")
	}
	waitForWSEvent(t, client.Events(), func(event WSEvent) bool {
		return event.Type == WSEventReconnected
	})
	require.True(t, client.IsAuthenticated())
}

func TestWSClientReportsSubscriptionRestoreFailure(t *testing.T) {
	upgrader := websocket.Upgrader{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}))
	defer server.Close()

	client := NewWSClient("", "", "", testWebSocketURL(server))
	client.reconnectBackoff = time.Hour
	defer client.Close()
	require.NoError(t, client.Connect(context.Background()))
	client.subscriptions["tickers:BTC-USDT"] = &subscription{
		channel:  "tickers",
		args:     map[string]interface{}{"instId": "BTC-USDT"},
		messages: make(chan []byte, 1),
	}

	client.mu.RLock()
	conn := client.conn
	client.mu.RUnlock()
	require.NotNil(t, conn)
	require.NoError(t, conn.Close())
	require.Error(t, client.restoreConnectionState())
	event := waitForWSEvent(t, client.Events(), func(event WSEvent) bool {
		return event.Type == WSEventSubscriptionRestoreFailed
	})
	require.Equal(t, "tickers", event.Channel)
	require.Error(t, event.Err)
}

func TestWSClientHandleMessageDropsFullBuffer(t *testing.T) {
	drops := make(chan WSDrop, 1)
	client := NewWSClient("", "", "", WSPublicURL, WithWSDropHandler(func(drop WSDrop) { drops <- drop }))
	defer client.Close()

	messages := make(chan []byte, 1)
	messages <- []byte("already buffered")
	client.subscriptions["tickers:BTC-USDT"] = &subscription{
		channel:  "tickers",
		args:     map[string]interface{}{"instId": "BTC-USDT"},
		messages: messages,
	}
	client.handleMessage([]byte(`{"arg":{"channel":"tickers","instId":"BTC-USDT"},"data":[]}`))
	require.Equal(t, "already buffered", string(<-messages))

	select {
	case drop := <-drops:
		require.Equal(t, uint64(1), drop.Count)
		require.Equal(t, "tickers", drop.Channel)
	case <-time.After(time.Second):
		t.Fatal("drop handler was not called")
	}
	event := waitForWSEvent(t, client.Events(), func(event WSEvent) bool {
		return event.Type == WSEventMessageDropped
	})
	require.ErrorIs(t, event.Err, ErrWSMessageDropped)
	require.Equal(t, uint64(1), event.Dropped)
}

func TestWSClientHeartbeatUsesTextPingAndReconnectsWithoutPong(t *testing.T) {
	upgrader := websocket.Upgrader{}
	var connections atomic.Int32
	pings := make(chan struct{}, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		connections.Add(1)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			if string(message) == "ping" {
				select {
				case pings <- struct{}{}:
				default:
				}
			}
		}
	}))
	defer server.Close()

	client := NewWSClient("", "", "", testWebSocketURL(server))
	client.pingInterval = 15 * time.Millisecond
	client.pongTimeout = 20 * time.Millisecond
	client.reconnectBackoff = 5 * time.Millisecond
	client.reconnectMaxBackoff = 10 * time.Millisecond
	client.readTimeout = time.Second
	defer client.Close()
	require.NoError(t, client.Connect(context.Background()))

	select {
	case <-pings:
	case <-time.After(time.Second):
		t.Fatal("client did not send a text ping")
	}
	deadline := time.After(time.Second)
	for connections.Load() < 2 {
		select {
		case <-deadline:
			t.Fatal("client did not reconnect after the pong timeout")
		case <-time.After(5 * time.Millisecond):
		}
	}
}

func TestWSClientHeartbeatAcceptsPong(t *testing.T) {
	upgrader := websocket.Upgrader{}
	var connections atomic.Int32
	pings := make(chan struct{}, 2)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		connections.Add(1)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			if string(message) == "ping" {
				_ = conn.WriteMessage(websocket.TextMessage, []byte("pong"))
				pings <- struct{}{}
			}
		}
	}))
	defer server.Close()

	client := NewWSClient("", "", "", testWebSocketURL(server))
	client.pingInterval = 15 * time.Millisecond
	client.pongTimeout = 20 * time.Millisecond
	client.reconnectBackoff = 5 * time.Millisecond
	client.readTimeout = time.Second
	defer client.Close()
	require.NoError(t, client.Connect(context.Background()))
	for i := 0; i < 2; i++ {
		select {
		case <-pings:
		case <-time.After(time.Second):
			t.Fatal("client did not send a text ping")
		}
	}
	require.Equal(t, int32(1), connections.Load(), "a pong must prevent a reconnect")
}

func TestWSClientCloseInterruptsReconnectBackoff(t *testing.T) {
	upgrader := websocket.Upgrader{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err == nil {
			_ = conn.Close()
		}
	}))
	defer server.Close()

	client := NewWSClient("", "", "", testWebSocketURL(server))
	client.reconnectBackoff = 5 * time.Second
	require.NoError(t, client.Connect(context.Background()))
	waitForWSEvent(t, client.Events(), func(event WSEvent) bool {
		return event.Type == WSEventDisconnected
	})
	started := time.Now()
	require.NoError(t, client.Close())
	require.Less(t, time.Since(started), 100*time.Millisecond)
}

func TestWSClientContextAndConcurrentSubscriptionSafety(t *testing.T) {
	canceled, cancel := context.WithCancel(context.Background())
	cancel()
	client := NewWSClient("", "", "", WSPublicURL)
	defer client.Close()
	require.ErrorIs(t, client.Connect(canceled), context.Canceled)
	require.ErrorIs(t, client.Login(canceled), context.Canceled)
	_, err := client.Subscribe(canceled, "tickers", map[string]interface{}{"instId": "BTC-USDT"})
	require.ErrorIs(t, err, context.Canceled)

	upgrader := websocket.Upgrader{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}))
	defer server.Close()

	connected := NewWSClient("", "", "", testWebSocketURL(server))
	defer connected.Close()
	require.NoError(t, connected.Connect(context.Background()))
	var wg sync.WaitGroup
	errs := make(chan error, 40)
	for i := 0; i < 20; i++ {
		id := "BTC-USDT-" + string(rune('A'+i))
		wg.Add(1)
		go func(instID string) {
			defer wg.Done()
			_, err := connected.Subscribe(context.Background(), "tickers", map[string]interface{}{"instId": instID})
			errs <- err
		}(id)
	}
	wg.Wait()
	for i := 0; i < 20; i++ {
		id := "BTC-USDT-" + string(rune('A'+i))
		wg.Add(1)
		go func(instID string) {
			defer wg.Done()
			errs <- connected.Unsubscribe("tickers", map[string]interface{}{"instId": instID})
		}(id)
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		require.NoError(t, err)
	}
}

func TestValidateOrderBookSequence(t *testing.T) {
	previous := models.OrderBook{SeqID: 100}
	require.NoError(t, ValidateOrderBookSequence(previous, models.OrderBook{SeqID: 101, PrevSeqID: 100}))
	require.NoError(t, ValidateOrderBookSequence(previous, models.OrderBook{SeqID: 50, PrevSeqID: 100}))
	require.NoError(t, ValidateOrderBookSequence(previous, models.OrderBook{SeqID: 1, PrevSeqID: -1}))
	err := ValidateOrderBookSequence(previous, models.OrderBook{SeqID: 102, PrevSeqID: 99})
	require.ErrorIs(t, err, ErrWSSequenceGap)
}

func TestWSServerErrorUnwrap(t *testing.T) {
	loginErr := &WSServerError{Event: "login", Code: "50113", Message: "Invalid Sign"}
	require.ErrorIs(t, loginErr, ErrWSAuthenticationFailed)
	subscriptionErr := &WSServerError{Event: "subscribe", Code: "60012", Message: "Invalid request"}
	require.ErrorIs(t, subscriptionErr, ErrWSSubscriptionRejected)
	require.False(t, errors.Is(&WSServerError{Event: "notice"}, ErrWSAuthenticationFailed))
}

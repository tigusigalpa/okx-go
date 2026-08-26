package okx

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

func TestWSClientSendsLoginAndSubscriptionMessages(t *testing.T) {
	upgrader := websocket.Upgrader{}
	messages := make(chan []byte, 3)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		require.NoError(t, err)
		defer conn.Close()
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			messages <- message
		}
	}))
	defer server.Close()

	url := "ws" + strings.TrimPrefix(server.URL, "http")
	client := NewWSClient("api-key", "secret", "passphrase", url)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	require.NoError(t, client.Connect(ctx))
	defer client.Close()

	require.NoError(t, client.Login(ctx))
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
}

func TestWSClientHandleMessageDeliversSubscriptionData(t *testing.T) {
	client := NewWSClient("", "", "", WSPublicURL)
	defer client.Close()

	messages := make(chan []byte, 1)
	client.subscriptions["tickers:BTC-USDT"] = subscription{
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
}

func TestWSClientHelpers(t *testing.T) {
	require.Equal(t, WSDemoPublicURL, demoWebSocketURL(WSPublicURL))
	require.Equal(t, WSDemoPrivateURL, demoWebSocketURL(WSPrivateURL))
	require.Equal(t, WSDemoBusinessURL, demoWebSocketURL(WSBusinessURL))
	require.Equal(t, WSDemoPublicSBEURL, demoWebSocketURL(WSPublicSBEURL))
	require.Equal(t, "wss://custom.example", demoWebSocketURL("wss://custom.example"))

	args := map[string]interface{}{"instId": "BTC-USDT"}
	cloned := cloneArgs(args)
	cloned["instId"] = "ETH-USDT"
	require.Equal(t, "BTC-USDT", args["instId"])

	logger := &noopLogger{}
	client := NewWSClient("", "", "", WSPublicURL, WithWSDemo(), WithWSLogger(logger))
	require.Equal(t, WSDemoPublicURL, client.url)
	require.Same(t, logger, client.logger)
	require.NoError(t, client.Close())
}

func TestWSClientSubscriptionErrorsWithoutConnection(t *testing.T) {
	client := NewWSClient("api-key", "secret", "passphrase", WSPublicURL)
	defer client.Close()

	ctx := context.Background()
	require.ErrorContains(t, client.Login(ctx), "WebSocket not connected")
	_, err := client.Subscribe(ctx, "tickers", map[string]interface{}{"instId": "BTC-USDT"})
	require.ErrorContains(t, err, "WebSocket not connected")
	require.Empty(t, client.subscriptions)
	require.ErrorContains(t, client.Unsubscribe("tickers", map[string]interface{}{"instId": "BTC-USDT"}), "not subscribed")
	client.url = "://invalid"
	require.ErrorContains(t, client.Connect(ctx), "failed to connect to WebSocket")
}

func TestWSClientReconnectRestoresSubscriptions(t *testing.T) {
	upgrader := websocket.Upgrader{}
	messages := make(chan []byte, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		require.NoError(t, err)
		defer conn.Close()
		_, message, err := conn.ReadMessage()
		require.NoError(t, err)
		messages <- message
	}))
	defer server.Close()

	client := NewWSClient("", "", "", "ws"+strings.TrimPrefix(server.URL, "http"))
	client.subscriptions["tickers:BTC-USDT"] = subscription{
		channel:  "tickers",
		args:     map[string]interface{}{"instId": "BTC-USDT"},
		messages: make(chan []byte, 1),
	}
	defer client.Close()

	client.handleReconnect()
	select {
	case message := <-messages:
		require.Contains(t, string(message), `"op":"subscribe"`)
	case <-time.After(3 * time.Second):
		t.Fatal("reconnect did not restore the subscription")
	}
}

func TestWSClientHandleMessageDropsFullBuffer(t *testing.T) {
	client := NewWSClient("", "", "", WSPublicURL)
	defer client.Close()

	messages := make(chan []byte, 1)
	messages <- []byte("already buffered")
	client.subscriptions["tickers:BTC-USDT"] = subscription{
		channel:  "tickers",
		args:     map[string]interface{}{"instId": "BTC-USDT"},
		messages: messages,
	}
	client.handleMessage([]byte(`{"arg":{"channel":"tickers","instId":"BTC-USDT"},"data":[]}`))
	require.Equal(t, "already buffered", string(<-messages))

	client.handleMessage([]byte(`{"arg":{"instId":"BTC-USDT"},"data":[]}`))
}

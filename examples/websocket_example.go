package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tigusigalpa/okx-go"
)

func main() {
	// Example 1: Public WebSocket - Subscribe to ticker
	fmt.Println("=== Example 1: Public WebSocket - Ticker ===")
	publicExample()

	// Example 2: Private WebSocket - Subscribe to account updates
	fmt.Println("\n=== Example 2: Private WebSocket - Account Updates ===")
	privateExample()
}

func publicExample() {
	// Create WebSocket client for public channels
	ws := okx.NewWSClient(
		"", // No credentials needed for public channels
		"",
		"",
		okx.WSPublicURL,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Connect to WebSocket
	if err := ws.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer ws.Close()

	// Subscribe to ticker channel for BTC-USDT
	tickerCh, err := ws.Subscribe(ctx, "tickers", map[string]interface{}{
		"instId": "BTC-USDT",
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to ticker: %v", err)
	}

	// Subscribe to trades channel
	tradesCh, err := ws.Subscribe(ctx, "trades", map[string]interface{}{
		"instId": "BTC-USDT",
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to trades: %v", err)
	}

	// Set up signal handling for graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// Read messages
	messageCount := 0
	maxMessages := 10

	for {
		select {
		case msg := <-tickerCh:
			var data map[string]interface{}
			if err := json.Unmarshal(msg, &data); err != nil {
				log.Printf("Error unmarshaling ticker: %v", err)
				continue
			}
			fmt.Printf("Ticker update: %s\n", string(msg))
			messageCount++

		case msg := <-tradesCh:
			var data map[string]interface{}
			if err := json.Unmarshal(msg, &data); err != nil {
				log.Printf("Error unmarshaling trade: %v", err)
				continue
			}
			fmt.Printf("Trade update: %s\n", string(msg))
			messageCount++

		case <-sigCh:
			fmt.Println("\nReceived interrupt signal, shutting down...")
			return

		case <-ctx.Done():
			fmt.Println("\nContext timeout, shutting down...")
			return
		}

		if messageCount >= maxMessages {
			fmt.Printf("\nReceived %d messages, exiting...\n", messageCount)
			return
		}
	}
}

func privateExample() {
	// Get credentials from environment variables
	apiKey := os.Getenv("OKX_API_KEY")
	secretKey := os.Getenv("OKX_SECRET_KEY")
	passphrase := os.Getenv("OKX_PASSPHRASE")

	if apiKey == "" || secretKey == "" || passphrase == "" {
		log.Println("Skipping private WebSocket example: credentials not set")
		log.Println("Set OKX_API_KEY, OKX_SECRET_KEY, and OKX_PASSPHRASE environment variables")
		return
	}

	// Create WebSocket client for private channels
	ws := okx.NewWSClient(
		apiKey,
		secretKey,
		passphrase,
		okx.WSPrivateURL,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Connect to WebSocket
	if err := ws.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer ws.Close()

	// Authenticate
	if err := ws.Login(ctx); err != nil {
		log.Fatalf("Failed to authenticate: %v", err)
	}

	fmt.Println("Successfully authenticated!")

	// Subscribe to account channel
	accountCh, err := ws.Subscribe(ctx, "account", map[string]interface{}{})
	if err != nil {
		log.Fatalf("Failed to subscribe to account: %v", err)
	}

	// Subscribe to positions channel
	positionsCh, err := ws.Subscribe(ctx, "positions", map[string]interface{}{
		"instType": "SPOT",
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to positions: %v", err)
	}

	// Subscribe to orders channel
	ordersCh, err := ws.Subscribe(ctx, "orders", map[string]interface{}{
		"instType": "SPOT",
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to orders: %v", err)
	}

	// Set up signal handling for graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Listening for account updates... (Press Ctrl+C to exit)")

	// Read messages
	for {
		select {
		case msg := <-accountCh:
			fmt.Printf("\n[ACCOUNT UPDATE]\n%s\n", string(msg))

		case msg := <-positionsCh:
			fmt.Printf("\n[POSITION UPDATE]\n%s\n", string(msg))

		case msg := <-ordersCh:
			fmt.Printf("\n[ORDER UPDATE]\n%s\n", string(msg))

		case <-sigCh:
			fmt.Println("\nReceived interrupt signal, shutting down...")
			return

		case <-ctx.Done():
			fmt.Println("\nContext timeout, shutting down...")
			return
		}
	}
}

// Example 3: Multiple instruments
func multiInstrumentExample() {
	ws := okx.NewWSClient("", "", "", okx.WSPublicURL)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := ws.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer ws.Close()

	// Subscribe to multiple instruments
	instruments := []string{"BTC-USDT", "ETH-USDT", "SOL-USDT"}
	channels := make(map[string]<-chan []byte)

	for _, inst := range instruments {
		ch, err := ws.Subscribe(ctx, "tickers", map[string]interface{}{
			"instId": inst,
		})
		if err != nil {
			log.Fatalf("Failed to subscribe to %s: %v", inst, err)
		}
		channels[inst] = ch
	}

	fmt.Println("Subscribed to multiple instruments, listening for updates...")

	// Read messages from all channels
	for {
		for inst, ch := range channels {
			select {
			case msg := <-ch:
				fmt.Printf("[%s] %s\n", inst, string(msg))
			case <-ctx.Done():
				return
			default:
				// Non-blocking
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

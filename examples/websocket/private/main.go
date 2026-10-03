package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tigusigalpa/okx-go"
)

func main() {
	apiKey := os.Getenv("OKX_API_KEY")
	secretKey := os.Getenv("OKX_SECRET_KEY")
	passphrase := os.Getenv("OKX_PASSPHRASE")
	if apiKey == "" || secretKey == "" || passphrase == "" {
		log.Println("set OKX_API_KEY, OKX_SECRET_KEY, and OKX_PASSPHRASE to run this example")
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	ws := okx.NewWSClient(apiKey, secretKey, passphrase, okx.WSPrivateURL)
	defer ws.Close()
	if err := ws.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	// Login returns only after OKX acknowledges the authentication request.
	if err := ws.Login(ctx); err != nil {
		log.Fatal(err)
	}

	stream, err := ws.Subscribe(ctx, "account", map[string]interface{}{})
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case raw, ok := <-stream:
			if !ok {
				return
			}
			fmt.Println(string(raw))
		case event := <-ws.Events():
			if event.Err != nil {
				log.Printf("WebSocket %s: %v", event.Type, event.Err)
			}
		case <-ctx.Done():
			return
		}
	}
}

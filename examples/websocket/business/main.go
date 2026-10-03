package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tigusigalpa/okx-go"
	"github.com/tigusigalpa/okx-go/models"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	ws := okx.NewWSClient("", "", "", okx.WSBusinessURL)
	defer ws.Close()
	if err := ws.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	stream, err := ws.Subscribe(ctx, "candle1m", map[string]interface{}{"instId": "BTC-USDT"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case raw, ok := <-stream:
			if !ok {
				return
			}
			message, err := models.DecodeWSMessage[models.Candle](raw)
			if err != nil {
				log.Printf("decode candle: %v", err)
				continue
			}
			for _, candle := range message.Data {
				fmt.Printf("ts=%s close=%s confirmed=%t\n", candle.TS, candle.C, candle.IsConfirmed())
			}
		case event := <-ws.Events():
			if event.Err != nil {
				log.Printf("WebSocket %s: %v", event.Type, event.Err)
			}
		case <-ctx.Done():
			return
		}
	}
}

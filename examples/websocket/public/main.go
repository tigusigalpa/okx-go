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

	endpoint, err := okx.WebSocketURL(okx.WSRegionGlobal, okx.WSServicePublic, false)
	if err != nil {
		log.Fatal(err)
	}
	ws := okx.NewWSClient("", "", "", endpoint, okx.WithWSSubscriptionBuffer(256))
	defer ws.Close()
	if err := ws.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	stream, err := ws.Subscribe(ctx, "tickers", map[string]interface{}{"instId": "BTC-USDT"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case raw, ok := <-stream:
			if !ok {
				return
			}
			message, err := models.DecodeWSMessage[models.Ticker](raw)
			if err != nil {
				log.Printf("decode ticker: %v", err)
				continue
			}
			for _, ticker := range message.Data {
				fmt.Printf("%s last=%s at %s\n", ticker.InstID, ticker.Last, ticker.TS)
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

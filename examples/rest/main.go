package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tigusigalpa/okx-go"
	"github.com/tigusigalpa/okx-go/models"
)

func main() {
	// Create a new REST client with demo trading enabled
	client := okx.NewRestClient(
		"your-api-key",
		"your-secret-key",
		"your-passphrase",
		okx.WithDemoTrading(),
		okx.WithTimeout(30*time.Second),
	)

	ctx := context.Background()

	// Example 1: Get account balance
	fmt.Println("=== Example 1: Get Account Balance ===")
	balances, err := client.Account.GetBalance(ctx, nil)
	if err != nil {
		log.Printf("Error getting balance: %v\n", err)
	} else {
		for _, balance := range balances {
			if balance.TotalEq != nil {
				fmt.Printf("Total Equity: %s\n", *balance.TotalEq)
			}
			for _, detail := range balance.Details {
				fmt.Printf("  %s: %s (Available: %s)\n", 
					detail.Ccy, 
					*detail.Eq, 
					*detail.AvailBal)
			}
		}
	}

	// Example 2: Get market ticker
	fmt.Println("\n=== Example 2: Get Market Ticker ===")
	tickers, err := client.Market.GetTicker(ctx, "BTC-USDT")
	if err != nil {
		log.Printf("Error getting ticker: %v\n", err)
	} else {
		for _, ticker := range tickers {
			fmt.Printf("Instrument: %s\n", ticker.InstID)
			fmt.Printf("Last Price: %s\n", ticker.Last)
			fmt.Printf("24h High: %s\n", ticker.High24h)
			fmt.Printf("24h Low: %s\n", ticker.Low24h)
			fmt.Printf("24h Volume: %s\n", ticker.Vol24h)
		}
	}

	// Example 3: Get available instruments
	fmt.Println("\n=== Example 3: Get Available Instruments ===")
	instruments, err := client.Public.GetInstruments(ctx, "SPOT", nil, nil, nil)
	if err != nil {
		log.Printf("Error getting instruments: %v\n", err)
	} else {
		fmt.Printf("Found %d SPOT instruments\n", len(instruments))
		for i, inst := range instruments {
			if i >= 5 {
				fmt.Println("  ...")
				break
			}
			fmt.Printf("  %s: %s/%s (Min: %s, Lot: %s)\n",
				inst.InstID,
				*inst.BaseCcy,
				*inst.QuoteCcy,
				inst.MinSz,
				inst.LotSz)
		}
	}

	// Example 4: Place a limit order
	fmt.Println("\n=== Example 4: Place Limit Order ===")
	px := "30000"
	orderReq := models.PlaceOrderRequest{
		InstID:  "BTC-USDT",
		TdMode:  "cash",
		Side:    "buy",
		OrdType: "limit",
		Px:      &px,
		Sz:      "0.001",
	}

	orderResult, err := client.Trade.PlaceOrder(ctx, orderReq)
	if err != nil {
		log.Printf("Error placing order: %v\n", err)
	} else {
		for _, result := range orderResult {
			fmt.Printf("Order placed successfully!\n")
			fmt.Printf("Order ID: %s\n", result.OrdID)
			fmt.Printf("Client Order ID: %s\n", result.ClOrdID)
			fmt.Printf("Status Code: %s\n", result.SCode)
			fmt.Printf("Status Message: %s\n", result.SMsg)
		}
	}

	// Example 5: Get pending orders
	fmt.Println("\n=== Example 5: Get Pending Orders ===")
	orders, err := client.Trade.GetOrdersPending(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		log.Printf("Error getting pending orders: %v\n", err)
	} else {
		fmt.Printf("Found %d pending orders\n", len(orders))
		for _, order := range orders {
			fmt.Printf("  Order %s: %s %s %s @ %s (Filled: %s/%s)\n",
				order.OrdID,
				order.Side,
				order.InstID,
				order.OrdType,
				order.Px,
				order.AccFillSz,
				order.Sz)
		}
	}

	// Example 6: Get account positions
	fmt.Println("\n=== Example 6: Get Account Positions ===")
	positions, err := client.Account.GetPositions(ctx, nil, nil)
	if err != nil {
		log.Printf("Error getting positions: %v\n", err)
	} else {
		fmt.Printf("Found %d positions\n", len(positions))
		for _, pos := range positions {
			fmt.Printf("  %s: %s %s (Avg Price: %s, P&L: %s)\n",
				pos.InstID,
				pos.PosSide,
				*pos.Pos,
				*pos.AvgPx,
				*pos.Upl)
		}
	}

	// Example 7: Get trade fee
	fmt.Println("\n=== Example 7: Get Trade Fee ===")
	fees, err := client.Account.GetTradeFee(ctx, "SPOT", nil, nil, nil, nil)
	if err != nil {
		log.Printf("Error getting trade fee: %v\n", err)
	} else {
		for _, fee := range fees {
			fmt.Printf("Category: %s, Level: %s\n", fee.Category, fee.Level)
			if fee.Maker != nil {
				fmt.Printf("  Maker: %s, Taker: %s\n", *fee.Maker, *fee.Taker)
			}
		}
	}

	// Example 8: Get candles (OHLCV data)
	fmt.Println("\n=== Example 8: Get Candles ===")
	bar := "1H"
	limit := "10"
	candles, err := client.Market.GetCandles(ctx, "BTC-USDT", &bar, nil, nil, &limit)
	if err != nil {
		log.Printf("Error getting candles: %v\n", err)
	} else {
		fmt.Printf("Got %d candles\n", len(candles))
		for i, candle := range candles {
			if i >= 3 {
				fmt.Println("  ...")
				break
			}
			fmt.Printf("  Time: %s, O: %s, H: %s, L: %s, C: %s, Vol: %s\n",
				candle.TS,
				candle.O,
				candle.H,
				candle.L,
				candle.C,
				candle.Vol)
		}
	}

	// Example 9: Get funding rate (for perpetual contracts)
	fmt.Println("\n=== Example 9: Get Funding Rate ===")
	fundingRates, err := client.Public.GetFundingRate(ctx, "BTC-USDT-SWAP")
	if err != nil {
		log.Printf("Error getting funding rate: %v\n", err)
	} else {
		for _, rate := range fundingRates {
			fmt.Printf("Instrument: %s\n", rate.InstID)
			fmt.Printf("Current Funding Rate: %s\n", rate.FundingRate)
			fmt.Printf("Next Funding Rate: %s\n", rate.NextFundingRate)
			fmt.Printf("Next Funding Time: %s\n", rate.NextFundingTime)
		}
	}

	// Example 10: Get system time
	fmt.Println("\n=== Example 10: Get System Time ===")
	times, err := client.Public.GetSystemTime(ctx)
	if err != nil {
		log.Printf("Error getting system time: %v\n", err)
	} else {
		for _, t := range times {
			fmt.Printf("Server Time: %s\n", t.TS)
		}
	}
}

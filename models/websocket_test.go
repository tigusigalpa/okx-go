package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeWSMessageTickerAndTrades(t *testing.T) {
	tickerPayload := []byte(`{
		"arg":{"channel":"tickers","instType":"SPOT"},
		"data":[
			{"instType":"SPOT","instId":"BTC-USDT","last":"70000.1","lastSz":"0.01","askPx":"70000.2","askSz":"1","bidPx":"70000.1","bidSz":"2","open24h":"69000","high24h":"71000","low24h":"68000","volCcy24h":"100","vol24h":"2","sodUtc0":"69500","sodUtc8":"69400","ts":"1710000000000"},
			{"instType":"SPOT","instId":"ETH-USDT","last":"3500.1","ts":"1710000000001"}
		]
	}`)
	tickers, err := DecodeWSMessage[Ticker](tickerPayload)
	require.NoError(t, err)
	require.Equal(t, "tickers", tickers.Arg.Channel)
	require.Len(t, tickers.Data, 2)
	require.Equal(t, "BTC-USDT", tickers.Data[0].InstID)
	require.Equal(t, "3500.1", tickers.Data[1].Last)

	tradePayload := []byte(`{"arg":{"channel":"trades","instId":"BTC-USDT"},"data":[{"instId":"BTC-USDT","tradeId":"123","px":"70000.1","sz":"0.01","side":"buy","ts":"1710000000000","count":"2","source":"0","seqId":42},{"instId":"BTC-USDT","tradeId":"124","px":"70000.0","sz":"0.02","side":"sell","ts":"1710000000001","count":"1","seqId":43}]}`)
	trades, err := DecodeWSMessage[Trade](tradePayload)
	require.NoError(t, err)
	require.Equal(t, "BTC-USDT", trades.Arg.InstID)
	require.Len(t, trades.Data, 2)
	require.Equal(t, int64(42), trades.Data[0].SeqID)
	require.Equal(t, "2", trades.Data[0].Count)
	require.Equal(t, "sell", trades.Data[1].Side)
}

func TestDecodeWSMessageCandles(t *testing.T) {
	marketPayload := []byte(`{"arg":{"channel":"candle1m","instId":"BTC-USDT"},"data":[["1710000000000","1","3","0.5","2","10","20","30","1"],["1710000060000","2","4","1","3","11","21","31","0","future-field"]]}`)
	market, err := DecodeWSMessage[Candle](marketPayload)
	require.NoError(t, err)
	require.Len(t, market.Data, 2)
	require.True(t, market.Data[0].IsConfirmed())
	require.False(t, market.Data[1].IsConfirmed())
	require.Equal(t, "30", market.Data[0].VolCcyQuote)
	require.Equal(t, "11", market.Data[1].Vol)

	indexPayload := []byte(`{"arg":{"channel":"candle1m","instId":"BTC-USD"},"data":[["1710000000000","1","3","0.5","2","1"]]}`)
	index, err := DecodeWSMessage[Candle](indexPayload)
	require.NoError(t, err)
	require.True(t, index.Data[0].IsConfirmed())
	require.Empty(t, index.Data[0].Vol)

	markPayload := []byte(`{"arg":{"channel":"mark-price-candle1m","instId":"BTC-USD-SWAP"},"data":[["1710000000000","1","3","0.5","2","0"]]}`)
	mark, err := DecodeWSMessage[Candle](markPayload)
	require.NoError(t, err)
	require.False(t, mark.Data[0].IsConfirmed())
}

func TestCandleRejectsMalformedOrIncompleteRows(t *testing.T) {
	tests := []struct {
		name    string
		payload string
	}{
		{"missing fields", `{"data":[["1","2","3","4","5"]]}`},
		{"partial market row", `{"data":[["1","2","3","4","5","6","7"]]}`},
		{"invalid confirm", `{"data":[["1","2","3","4","5","6","7","8","yes"]]}`},
		{"non-string value", `{"data":[[1,"2","3","4","5","1"]]}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DecodeWSMessage[Candle]([]byte(tt.payload))
			require.Error(t, err)
		})
	}
}

func TestDecodeWSMessageBooksAndSequenceFields(t *testing.T) {
	booksPayload := []byte(`{"arg":{"channel":"books","instId":"BTC-USDT"},"action":"snapshot","data":[{"asks":[["70001","1","0","1"]],"bids":[["70000","2","0","2"]],"ts":"1710000000000","checksum":0,"seqId":100,"prevSeqId":-1}]}`)
	books, err := DecodeWSMessage[OrderBook](booksPayload)
	require.NoError(t, err)
	require.Equal(t, "snapshot", books.Action)
	require.Equal(t, int64(-1), books.Data[0].PrevSeqID)
	require.Equal(t, int64(100), books.Data[0].SeqID)

	updatePayload := []byte(`{"arg":{"channel":"books","instId":"BTC-USDT"},"action":"update","data":[{"asks":[],"bids":[["70000","0","0","0"]],"ts":"1710000000001","checksum":0,"seqId":101,"prevSeqId":100}]}`)
	update, err := DecodeWSMessage[OrderBook](updatePayload)
	require.NoError(t, err)
	require.Equal(t, "update", update.Action)
	require.Equal(t, int64(100), update.Data[0].PrevSeqID)

	books5Payload := []byte(`{"arg":{"channel":"books5","instId":"BTC-USDT"},"action":"snapshot","data":[{"asks":[["70001","1","0","1"]],"bids":[["70000","2","0","2"]],"ts":"1710000000000","seqId":555}]}`)
	books5, err := DecodeWSMessage[OrderBook](books5Payload)
	require.NoError(t, err)
	require.Equal(t, "snapshot", books5.Action)
	require.Equal(t, int64(555), books5.Data[0].SeqID)
}

func TestDecodeWSMessageOpenInterestFundingLiquidationsAndMarkPrice(t *testing.T) {
	openInterestPayload := []byte(`{"arg":{"channel":"open-interest","instId":"BTC-USD-SWAP"},"data":[{"instType":"SWAP","instId":"BTC-USD-SWAP","oi":"123.45","oiCcy":"1.2","ts":"1710000000000"}]}`)
	openInterest, err := DecodeWSMessage[OpenInterest](openInterestPayload)
	require.NoError(t, err)
	require.Equal(t, "123.45", openInterest.Data[0].Oi)

	fundingPayload := []byte(`{"arg":{"channel":"funding-rate","instId":"BTC-USD-SWAP"},"data":[{"instType":"SWAP","instId":"BTC-USD-SWAP","fundingRate":"-0.0005","fundingTime":"1710003600000","nextFundingRate":"0.0001","nextFundingTime":"1710032400000","formulaType":"withRate","interestRate":"0.0001","impactValue":"10000","minFundingRate":"-0.003","maxFundingRate":"0.003","settState":"settled","settFundingRate":"-0.0004","ts":"1710000000000"},{"instType":"SWAP","instId":"ETH-USD-SWAP","fundingRate":"0.0002","fundingTime":"1710003600000","ts":"1710000000001"},{"instType":"SWAP","instId":"SOL-USD-SWAP","fundingRate":"0","fundingTime":"1710003600000","ts":"1710000000002"}]}`)
	funding, err := DecodeWSMessage[FundingRate](fundingPayload)
	require.NoError(t, err)
	require.Equal(t, "-0.0005", funding.Data[0].FundingRate)
	require.Equal(t, "0.0001", funding.Data[0].NextFundingRate)
	require.Equal(t, "0.0002", funding.Data[1].FundingRate)
	require.Equal(t, "0", funding.Data[2].FundingRate)

	liquidationPayload := []byte(`{"arg":{"channel":"liquidation-orders","instType":"SWAP"},"data":[{"instType":"SWAP","instId":"BTC-USD-SWAP","uly":"BTC-USD","details":[{"bkLoss":"10","bkPx":"70000","ccy":"BTC","instId":"BTC-USD-SWAP","posSide":"long","side":"sell","sz":"1","ts":"1710000000000"}]}]}`)
	liquidations, err := DecodeWSMessage[LiquidationOrder](liquidationPayload)
	require.NoError(t, err)
	require.Equal(t, "BTC-USD", liquidations.Data[0].Uly)
	require.Equal(t, "70000", liquidations.Data[0].Details[0].BkPx)

	markPricePayload := []byte(`{"arg":{"channel":"mark-price","instId":"BTC-USD-SWAP"},"data":[{"instType":"SWAP","instId":"BTC-USD-SWAP","markPx":"70001.2","ts":"1710000000000"}]}`)
	markPrices, err := DecodeWSMessage[MarkPrice](markPricePayload)
	require.NoError(t, err)
	require.Equal(t, "70001.2", markPrices.Data[0].MarkPx)
}

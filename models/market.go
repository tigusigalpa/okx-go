package models

type Ticker struct {
	InstType  string `json:"instType"`
	InstID    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24h   string `json:"open24h"`
	High24h   string `json:"high24h"`
	Low24h    string `json:"low24h"`
	VolCcy24h string `json:"volCcy24h"`
	Vol24h    string `json:"vol24h"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
	TS        string `json:"ts"`
}

type IndexTicker struct {
	InstID  string `json:"instId"`
	IdxPx   string `json:"idxPx"`
	High24h string `json:"high24h"`
	Low24h  string `json:"low24h"`
	Open24h string `json:"open24h"`
	SodUtc0 string `json:"sodUtc0"`
	SodUtc8 string `json:"sodUtc8"`
	TS      string `json:"ts"`
}

type OrderBook struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
	TS   string     `json:"ts"`
}

type Candle struct {
	TS       string `json:"ts"`
	O        string `json:"o"`
	H        string `json:"h"`
	L        string `json:"l"`
	C        string `json:"c"`
	Vol      string `json:"vol"`
	VolCcy   string `json:"volCcy"`
	VolCcyQuote string `json:"volCcyQuote"`
	Confirm  string `json:"confirm"`
}

type Trade struct {
	InstID  string `json:"instId"`
	TradeID string `json:"tradeId"`
	Px      string `json:"px"`
	Sz      string `json:"sz"`
	Side    string `json:"side"`
	TS      string `json:"ts"`
}

type Platform24Volume struct {
	VolCcy string `json:"volCcy"`
	VolUsd string `json:"volUsd"`
	TS     string `json:"ts"`
}

type OpenOracle struct {
	Messages   []string `json:"messages"`
	Prices     []string `json:"prices"`
	Signatures []string `json:"signatures"`
	Timestamp  string   `json:"timestamp"`
}

type ExchangeRate struct {
	UsdCny string `json:"usdCny"`
}

type IndexComponents struct {
	Index      string           `json:"index"`
	Last       string           `json:"last"`
	Components []IndexComponent `json:"components"`
	TS         string           `json:"ts"`
}

type IndexComponent struct {
	Exch   string `json:"exch"`
	Symbol string `json:"symbol"`
	SymPx  string `json:"symPx"`
	Wgt    string `json:"wgt"`
}

type BlockTicker struct {
	InstID  string `json:"instId"`
	InstType string `json:"instType"`
	Vol24h  string `json:"vol24h"`
	VolCcy24h string `json:"volCcy24h"`
	TS      string `json:"ts"`
}

type BlockTrade struct {
	InstID  string `json:"instId"`
	TradeID string `json:"tradeId"`
	Px      string `json:"px"`
	Sz      string `json:"sz"`
	Side    string `json:"side"`
	TS      string `json:"ts"`
}

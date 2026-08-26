package models

// Instrument represents an OKX API request or response value.
type Instrument struct {
	Alias        *string `json:"alias,omitempty"`
	BaseCcy      *string `json:"baseCcy,omitempty"`
	Category     string  `json:"category"`
	CtMult       *string `json:"ctMult,omitempty"`
	CtType       *string `json:"ctType,omitempty"`
	CtVal        *string `json:"ctVal,omitempty"`
	CtValCcy     *string `json:"ctValCcy,omitempty"`
	ExpTime      *string `json:"expTime,omitempty"`
	InstFamily   *string `json:"instFamily,omitempty"`
	InstID       string  `json:"instId"`
	InstType     string  `json:"instType"`
	Lever        *string `json:"lever,omitempty"`
	ListTime     string  `json:"listTime"`
	LotSz        string  `json:"lotSz"`
	MaxIcebergSz *string `json:"maxIcebergSz,omitempty"`
	MaxLmtAmt    *string `json:"maxLmtAmt,omitempty"`
	MaxLmtSz     *string `json:"maxLmtSz,omitempty"`
	MaxMktAmt    *string `json:"maxMktAmt,omitempty"`
	MaxMktSz     *string `json:"maxMktSz,omitempty"`
	MaxStopSz    *string `json:"maxStopSz,omitempty"`
	MaxTriggerSz *string `json:"maxTriggerSz,omitempty"`
	MaxTwapSz    *string `json:"maxTwapSz,omitempty"`
	MinSz        string  `json:"minSz"`
	OptType      *string `json:"optType,omitempty"`
	QuoteCcy     *string `json:"quoteCcy,omitempty"`
	SettleCcy    *string `json:"settleCcy,omitempty"`
	State        string  `json:"state"`
	Stk          *string `json:"stk,omitempty"`
	TickSz       string  `json:"tickSz"`
	Uly          *string `json:"uly,omitempty"`
}

// DeliveryExerciseHistory represents an OKX API request or response value.
type DeliveryExerciseHistory struct {
	Details []DeliveryExerciseDetail `json:"details"`
	TS      string                   `json:"ts"`
}

// DeliveryExerciseDetail represents an OKX API request or response value.
type DeliveryExerciseDetail struct {
	InsID string `json:"insId"`
	Px    string `json:"px"`
	Type  string `json:"type"`
}

// OpenInterest represents an OKX API request or response value.
type OpenInterest struct {
	InstID   string `json:"instId"`
	InstType string `json:"instType"`
	Oi       string `json:"oi"`
	OiCcy    string `json:"oiCcy"`
	TS       string `json:"ts"`
}

// FundingRate represents an OKX API request or response value.
type FundingRate struct {
	FundingRate     string `json:"fundingRate"`
	FundingTime     string `json:"fundingTime"`
	InstID          string `json:"instId"`
	InstType        string `json:"instType"`
	NextFundingRate string `json:"nextFundingRate"`
	NextFundingTime string `json:"nextFundingTime"`
}

// FundingRateHistory represents an OKX API request or response value.
type FundingRateHistory struct {
	FundingRate  string `json:"fundingRate"`
	FundingTime  string `json:"fundingTime"`
	InstID       string `json:"instId"`
	InstType     string `json:"instType"`
	RealizedRate string `json:"realizedRate"`
}

// PriceLimit represents an OKX API request or response value.
type PriceLimit struct {
	BuyLmt   string `json:"buyLmt"`
	InstID   string `json:"instId"`
	InstType string `json:"instType"`
	SellLmt  string `json:"sellLmt"`
	TS       string `json:"ts"`
}

// OptionSummary represents an OKX API request or response value.
type OptionSummary struct {
	AskVol   string `json:"askVol"`
	BidVol   string `json:"bidVol"`
	Delta    string `json:"delta"`
	DeltaBS  string `json:"deltaBS"`
	Gamma    string `json:"gamma"`
	GammaBS  string `json:"gammaBS"`
	InstID   string `json:"instId"`
	InstType string `json:"instType"`
	MarkVol  string `json:"markVol"`
	RealVol  string `json:"realVol"`
	Theta    string `json:"theta"`
	ThetaBS  string `json:"thetaBS"`
	TS       string `json:"ts"`
	Vega     string `json:"vega"`
	VegaBS   string `json:"vegaBS"`
}

// EstimatedPrice represents an OKX API request or response value.
type EstimatedPrice struct {
	InstID   string `json:"instId"`
	InstType string `json:"instType"`
	SettlePx string `json:"settlePx"`
	TS       string `json:"ts"`
}

// DiscountRateInterestFreeQuota represents an OKX API request or response value.
type DiscountRateInterestFreeQuota struct {
	Amt          string         `json:"amt"`
	Ccy          string         `json:"ccy"`
	DiscountInfo []DiscountInfo `json:"discountInfo"`
	DiscountLv   string         `json:"discountLv"`
}

// DiscountInfo represents an OKX API request or response value.
type DiscountInfo struct {
	DiscountRate string `json:"discountRate"`
	MaxAmt       string `json:"maxAmt"`
	MinAmt       string `json:"minAmt"`
}

// SystemTime represents an OKX API request or response value.
type SystemTime struct {
	TS string `json:"ts"`
}

// LiquidationOrder represents an OKX API request or response value.
type LiquidationOrder struct {
	Details []LiquidationDetail `json:"details"`
	TS      string              `json:"ts"`
}

// LiquidationDetail represents an OKX API request or response value.
type LiquidationDetail struct {
	BkLoss  string `json:"bkLoss"`
	BkPx    string `json:"bkPx"`
	Ccy     string `json:"ccy"`
	InstID  string `json:"instId"`
	PosSide string `json:"posSide"`
	Side    string `json:"side"`
	Sz      string `json:"sz"`
	TS      string `json:"ts"`
}

// MarkPrice represents an OKX API request or response value.
type MarkPrice struct {
	InstID   string `json:"instId"`
	InstType string `json:"instType"`
	MarkPx   string `json:"markPx"`
	TS       string `json:"ts"`
}

// PositionTier represents an OKX API request or response value.
type PositionTier struct {
	BaseMaxLoan  string `json:"baseMaxLoan"`
	BaseMaxSz    string `json:"baseMaxSz"`
	Imr          string `json:"imr"`
	InstID       string `json:"instId"`
	MaxLever     string `json:"maxLever"`
	MaxSz        string `json:"maxSz"`
	MinSz        string `json:"minSz"`
	Mmr          string `json:"mmr"`
	OptMgnFactor string `json:"optMgnFactor"`
	QuoteMaxLoan string `json:"quoteMaxLoan"`
	Tier         string `json:"tier"`
	Uly          string `json:"uly"`
}

// InterestRateLoanQuota represents an OKX API request or response value.
type InterestRateLoanQuota struct {
	Basic []InterestRateBasic `json:"basic"`
	Vip   []InterestRateVIP   `json:"vip"`
}

// InterestRateBasic represents an OKX API request or response value.
type InterestRateBasic struct {
	Ccy   string `json:"ccy"`
	Quota string `json:"quota"`
	Rate  string `json:"rate"`
}

// InterestRateVIP represents an OKX API request or response value.
type InterestRateVIP struct {
	IrDiscount    string `json:"irDiscount"`
	LoanQuotaCoef string `json:"loanQuotaCoef"`
	Level         string `json:"level"`
}

// VIPInterestRateLoanQuota represents an OKX API request or response value.
type VIPInterestRateLoanQuota struct {
	Ccy   string `json:"ccy"`
	Quota string `json:"quota"`
	Rate  string `json:"rate"`
}

// Underlying represents an OKX API request or response value.
type Underlying struct {
	Uly string `json:"uly"`
}

// InsuranceFund represents an OKX API request or response value.
type InsuranceFund struct {
	Details []InsuranceFundDetail `json:"details"`
	Total   string                `json:"total"`
}

// InsuranceFundDetail represents an OKX API request or response value.
type InsuranceFundDetail struct {
	Amt  string `json:"amt"`
	Ccy  string `json:"ccy"`
	Type string `json:"type"`
	TS   string `json:"ts"`
}

// UnitConvert represents an OKX API request or response value.
type UnitConvert struct {
	InstID string `json:"instId"`
	Px     string `json:"px"`
	Sz     string `json:"sz"`
	Type   string `json:"type"`
	Unit   string `json:"unit"`
}

// EconomicCalendar represents an OKX API request or response value.
type EconomicCalendar struct {
	CalendarID string `json:"calendarId"`
	Country    string `json:"country"`
	Region     string `json:"region"`
	Event      string `json:"event"`
	RefValue   string `json:"refValue"`
	Previous   string `json:"previous"`
	Forecast   string `json:"forecast"`
	Actual     string `json:"actual"`
	Impact     string `json:"impact"`
	DateStr    string `json:"dateStr"`
	TS         string `json:"ts"`
}

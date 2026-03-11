package models

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

type DeliveryExerciseHistory struct {
	Details []DeliveryExerciseDetail `json:"details"`
	TS      string                   `json:"ts"`
}

type DeliveryExerciseDetail struct {
	InsID string `json:"insId"`
	Px    string `json:"px"`
	Type  string `json:"type"`
}

type OpenInterest struct {
	InstID   string `json:"instId"`
	InstType string `json:"instType"`
	Oi       string `json:"oi"`
	OiCcy    string `json:"oiCcy"`
	TS       string `json:"ts"`
}

type FundingRate struct {
	FundingRate     string `json:"fundingRate"`
	FundingTime     string `json:"fundingTime"`
	InstID          string `json:"instId"`
	InstType        string `json:"instType"`
	NextFundingRate string `json:"nextFundingRate"`
	NextFundingTime string `json:"nextFundingTime"`
}

type FundingRateHistory struct {
	FundingRate string `json:"fundingRate"`
	FundingTime string `json:"fundingTime"`
	InstID      string `json:"instId"`
	InstType    string `json:"instType"`
	RealizedRate string `json:"realizedRate"`
}

type PriceLimit struct {
	BuyLmt  string `json:"buyLmt"`
	InstID  string `json:"instId"`
	InstType string `json:"instType"`
	SellLmt string `json:"sellLmt"`
	TS      string `json:"ts"`
}

type OptionSummary struct {
	AskVol    string `json:"askVol"`
	BidVol    string `json:"bidVol"`
	Delta     string `json:"delta"`
	DeltaBS   string `json:"deltaBS"`
	Gamma     string `json:"gamma"`
	GammaBS   string `json:"gammaBS"`
	InstID    string `json:"instId"`
	InstType  string `json:"instType"`
	MarkVol   string `json:"markVol"`
	RealVol   string `json:"realVol"`
	Theta     string `json:"theta"`
	ThetaBS   string `json:"thetaBS"`
	TS        string `json:"ts"`
	Vega      string `json:"vega"`
	VegaBS    string `json:"vegaBS"`
}

type EstimatedPrice struct {
	InstID      string `json:"instId"`
	InstType    string `json:"instType"`
	SettlePx    string `json:"settlePx"`
	TS          string `json:"ts"`
}

type DiscountRateInterestFreeQuota struct {
	Amt          string                         `json:"amt"`
	Ccy          string                         `json:"ccy"`
	DiscountInfo []DiscountInfo                 `json:"discountInfo"`
	DiscountLv   string                         `json:"discountLv"`
}

type DiscountInfo struct {
	DiscountRate string `json:"discountRate"`
	MaxAmt       string `json:"maxAmt"`
	MinAmt       string `json:"minAmt"`
}

type SystemTime struct {
	TS string `json:"ts"`
}

type LiquidationOrder struct {
	Details []LiquidationDetail `json:"details"`
	TS      string              `json:"ts"`
}

type LiquidationDetail struct {
	BkLoss   string `json:"bkLoss"`
	BkPx     string `json:"bkPx"`
	Ccy      string `json:"ccy"`
	InstID   string `json:"instId"`
	PosSide  string `json:"posSide"`
	Side     string `json:"side"`
	Sz       string `json:"sz"`
	TS       string `json:"ts"`
}

type MarkPrice struct {
	InstID   string `json:"instId"`
	InstType string `json:"instType"`
	MarkPx   string `json:"markPx"`
	TS       string `json:"ts"`
}

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

type InterestRateLoanQuota struct {
	Basic []InterestRateBasic `json:"basic"`
	Vip   []InterestRateVIP   `json:"vip"`
}

type InterestRateBasic struct {
	Ccy       string `json:"ccy"`
	Quota     string `json:"quota"`
	Rate      string `json:"rate"`
}

type InterestRateVIP struct {
	IrDiscount string `json:"irDiscount"`
	LoanQuotaCoef string `json:"loanQuotaCoef"`
	Level     string `json:"level"`
}

type VIPInterestRateLoanQuota struct {
	Ccy       string `json:"ccy"`
	Quota     string `json:"quota"`
	Rate      string `json:"rate"`
}

type Underlying struct {
	Uly string `json:"uly"`
}

type InsuranceFund struct {
	Details []InsuranceFundDetail `json:"details"`
	Total   string                `json:"total"`
}

type InsuranceFundDetail struct {
	Amt    string `json:"amt"`
	Ccy    string `json:"ccy"`
	Type   string `json:"type"`
	TS     string `json:"ts"`
}

type UnitConvert struct {
	InstID string `json:"instId"`
	Px     string `json:"px"`
	Sz     string `json:"sz"`
	Type   string `json:"type"`
	Unit   string `json:"unit"`
}

type EconomicCalendar struct {
	CalendarID   string `json:"calendarId"`
	Country      string `json:"country"`
	Region       string `json:"region"`
	Event        string `json:"event"`
	RefValue     string `json:"refValue"`
	Previous     string `json:"previous"`
	Forecast     string `json:"forecast"`
	Actual       string `json:"actual"`
	Impact       string `json:"impact"`
	DateStr      string `json:"dateStr"`
	TS           string `json:"ts"`
}

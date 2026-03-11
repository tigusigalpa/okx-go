package models

type Balance struct {
	AdjEq       *string           `json:"adjEq,omitempty"`
	BorrowFroz  *string           `json:"borrowFroz,omitempty"`
	Details     []BalanceDetail   `json:"details,omitempty"`
	Imr         *string           `json:"imr,omitempty"`
	IsoEq       *string           `json:"isoEq,omitempty"`
	MgnRatio    *string           `json:"mgnRatio,omitempty"`
	Mmr         *string           `json:"mmr,omitempty"`
	NotionalUsd *string           `json:"notionalUsd,omitempty"`
	OrdFroz     *string           `json:"ordFroz,omitempty"`
	TotalEq     *string           `json:"totalEq,omitempty"`
	UTime       *string           `json:"uTime,omitempty"`
	UplTotalEq  *string           `json:"uplTotalEq,omitempty"`
}

type BalanceDetail struct {
	AvailBal      *string `json:"availBal,omitempty"`
	AvailEq       *string `json:"availEq,omitempty"`
	BorrowFroz    *string `json:"borrowFroz,omitempty"`
	CashBal       *string `json:"cashBal,omitempty"`
	Ccy           string  `json:"ccy"`
	CrossLiab     *string `json:"crossLiab,omitempty"`
	DisEq         *string `json:"disEq,omitempty"`
	Eq            *string `json:"eq,omitempty"`
	EqUsd         *string `json:"eqUsd,omitempty"`
	FixedBal      *string `json:"fixedBal,omitempty"`
	FrozenBal     *string `json:"frozenBal,omitempty"`
	Imr           *string `json:"imr,omitempty"`
	Interest      *string `json:"interest,omitempty"`
	IsoEq         *string `json:"isoEq,omitempty"`
	IsoLiab       *string `json:"isoLiab,omitempty"`
	IsoUpl        *string `json:"isoUpl,omitempty"`
	Liab          *string `json:"liab,omitempty"`
	MaxLoan       *string `json:"maxLoan,omitempty"`
	MgnRatio      *string `json:"mgnRatio,omitempty"`
	Mmr           *string `json:"mmr,omitempty"`
	NotionalLever *string `json:"notionalLever,omitempty"`
	OrdFrozen     *string `json:"ordFrozen,omitempty"`
	RewardBal     *string `json:"rewardBal,omitempty"`
	SpotInUseAmt  *string `json:"spotInUseAmt,omitempty"`
	StgyEq        *string `json:"stgyEq,omitempty"`
	Twap          *string `json:"twap,omitempty"`
	UTime         *string `json:"uTime,omitempty"`
	Upl           *string `json:"upl,omitempty"`
	UplLiab       *string `json:"uplLiab,omitempty"`
}

type Position struct {
	ADL          *string `json:"adl,omitempty"`
	AvailPos     *string `json:"availPos,omitempty"`
	AvgPx        *string `json:"avgPx,omitempty"`
	CTime        *string `json:"cTime,omitempty"`
	Ccy          *string `json:"ccy,omitempty"`
	DeltaBS      *string `json:"deltaBS,omitempty"`
	DeltaPA      *string `json:"deltaPA,omitempty"`
	GammaBS      *string `json:"gammaBS,omitempty"`
	GammaPA      *string `json:"gammaPA,omitempty"`
	Imr          *string `json:"imr,omitempty"`
	InstID       string  `json:"instId"`
	InstType     string  `json:"instType"`
	Interest     *string `json:"interest,omitempty"`
	Last         *string `json:"last,omitempty"`
	Lever        *string `json:"lever,omitempty"`
	LiabCcy      *string `json:"liabCcy,omitempty"`
	Liab         *string `json:"liab,omitempty"`
	LiqPx        *string `json:"liqPx,omitempty"`
	MarkPx       *string `json:"markPx,omitempty"`
	Margin       *string `json:"margin,omitempty"`
	MgnMode      string  `json:"mgnMode"`
	MgnRatio     *string `json:"mgnRatio,omitempty"`
	Mmr          *string `json:"mmr,omitempty"`
	NotionalUsd  *string `json:"notionalUsd,omitempty"`
	OptVal       *string `json:"optVal,omitempty"`
	PTime        *string `json:"pTime,omitempty"`
	Pos          *string `json:"pos,omitempty"`
	PosCcy       *string `json:"posCcy,omitempty"`
	PosID        *string `json:"posId,omitempty"`
	PosSide      string  `json:"posSide"`
	ThetaBS      *string `json:"thetaBS,omitempty"`
	ThetaPA      *string `json:"thetaPA,omitempty"`
	TradeID      *string `json:"tradeId,omitempty"`
	UTime        *string `json:"uTime,omitempty"`
	Upl          *string `json:"upl,omitempty"`
	UplRatio     *string `json:"uplRatio,omitempty"`
	VegaBS       *string `json:"vegaBS,omitempty"`
	VegaPA       *string `json:"vegaPA,omitempty"`
}

type AccountConfig struct {
	AcctLv      *string `json:"acctLv,omitempty"`
	AutoLoan    *bool   `json:"autoLoan,omitempty"`
	CtIsoMode   *string `json:"ctIsoMode,omitempty"`
	GreeksType  *string `json:"greeksType,omitempty"`
	Level       *string `json:"level,omitempty"`
	LevelTmp    *string `json:"levelTmp,omitempty"`
	MgnIsoMode  *string `json:"mgnIsoMode,omitempty"`
	PosMode     *string `json:"posMode,omitempty"`
	SpotOffsetType *string `json:"spotOffsetType,omitempty"`
	UID         *string `json:"uid,omitempty"`
}

type LeverageInfo struct {
	InstID  string  `json:"instId"`
	MgnMode string  `json:"mgnMode"`
	PosSide *string `json:"posSide,omitempty"`
	Lever   string  `json:"lever"`
}

type MaxSize struct {
	InstID  string  `json:"instId"`
	Ccy     *string `json:"ccy,omitempty"`
	MaxBuy  string  `json:"maxBuy"`
	MaxSell string  `json:"maxSell"`
}

type MaxAvailSize struct {
	InstID    string  `json:"instId"`
	AvailBuy  string  `json:"availBuy"`
	AvailSell string  `json:"availSell"`
}

type MaxLoan struct {
	InstID  string `json:"instId"`
	MgnMode string `json:"mgnMode"`
	MgnCcy  string `json:"mgnCcy"`
	MaxLoan string `json:"maxLoan"`
	Ccy     string `json:"ccy"`
	Side    string `json:"side"`
}

type TradeFee struct {
	Category string      `json:"category"`
	Delivery *string     `json:"delivery,omitempty"`
	Exercise *string     `json:"exercise,omitempty"`
	InstType string      `json:"instType"`
	Level    string      `json:"level"`
	Maker    *string     `json:"maker,omitempty"`
	MakerU   *string     `json:"makerU,omitempty"`
	Taker    *string     `json:"taker,omitempty"`
	TakerU   *string     `json:"takerU,omitempty"`
	TS       string      `json:"ts"`
}

type InterestAccrued struct {
	Ccy       string  `json:"ccy"`
	InstID    *string `json:"instId,omitempty"`
	Interest  string  `json:"interest"`
	InterestRate string `json:"interestRate"`
	Liab      string  `json:"liab"`
	MgnMode   *string `json:"mgnMode,omitempty"`
	TS        string  `json:"ts"`
	Type      string  `json:"type"`
}

type InterestRate struct {
	Ccy          string `json:"ccy"`
	InterestRate string `json:"interestRate"`
}

type Greeks struct {
	GreeksType string `json:"greeksType"`
	ThetaBS    string `json:"thetaBS"`
	ThetaPA    string `json:"thetaPA"`
	DeltaBS    string `json:"deltaBS"`
	DeltaPA    string `json:"deltaPA"`
	GammaBS    string `json:"gammaBS"`
	GammaPA    string `json:"gammaPA"`
	VegaBS     string `json:"vegaBS"`
	VegaPA     string `json:"vegaPA"`
	TS         string `json:"ts"`
}

type Bill struct {
	Bal        string  `json:"bal"`
	BalChg     string  `json:"balChg"`
	BillID     string  `json:"billId"`
	Ccy        string  `json:"ccy"`
	ExecType   string  `json:"execType"`
	Fee        string  `json:"fee"`
	From       *string `json:"from,omitempty"`
	InstID     *string `json:"instId,omitempty"`
	InstType   *string `json:"instType,omitempty"`
	MgnMode    *string `json:"mgnMode,omitempty"`
	Notes      *string `json:"notes,omitempty"`
	OrdID      *string `json:"ordId,omitempty"`
	PnL        string  `json:"pnl"`
	PosBal     string  `json:"posBal"`
	PosBalChg  string  `json:"posBalChg"`
	SubType    string  `json:"subType"`
	Sz         string  `json:"sz"`
	To         *string `json:"to,omitempty"`
	TS         string  `json:"ts"`
	Type       string  `json:"type"`
}

type SetLeverageRequest struct {
	InstID  *string `json:"instId,omitempty"`
	Ccy     *string `json:"ccy,omitempty"`
	Lever   string  `json:"lever"`
	MgnMode string  `json:"mgnMode"`
	PosSide *string `json:"posSide,omitempty"`
}

type SetPositionModeRequest struct {
	PosMode string `json:"posMode"`
}

type SetGreeksRequest struct {
	GreeksType string `json:"greeksType"`
}

type PositionMarginBalanceRequest struct {
	InstID   string  `json:"instId"`
	PosSide  string  `json:"posSide"`
	Type     string  `json:"type"`
	Amt      string  `json:"amt"`
	Ccy      *string `json:"ccy,omitempty"`
	Auto     *bool   `json:"auto,omitempty"`
	LoanTrans *bool  `json:"loanTrans,omitempty"`
}

type MaxWithdrawal struct {
	Ccy       string `json:"ccy"`
	MaxWd     string `json:"maxWd"`
	MaxWdEx   string `json:"maxWdEx"`
	SpotOffsetMaxWd string `json:"spotOffsetMaxWd"`
	SpotOffsetMaxWdEx string `json:"spotOffsetMaxWdEx"`
}

type RiskState struct {
	AtRisk         bool   `json:"atRisk"`
	AtRiskIdx      []string `json:"atRiskIdx,omitempty"`
	AtRiskMgn      []string `json:"atRiskMgn,omitempty"`
	TS             string `json:"ts"`
}

type BorrowRepayRequest struct {
	Ccy    string `json:"ccy"`
	Side   string `json:"side"`
	Amt    string `json:"amt"`
}

type BorrowRepayHistory struct {
	Ccy       string `json:"ccy"`
	Side      string `json:"side"`
	Amt       string `json:"amt"`
	TS        string `json:"ts"`
}

type InterestLimits struct {
	Debt      string `json:"debt"`
	Interest  string `json:"interest"`
	NextDiscountTime string `json:"nextDiscountTime"`
	NextInterestTime string `json:"nextInterestTime"`
	Records   []InterestLimitRecord `json:"records"`
}

type InterestLimitRecord struct {
	Ccy           string `json:"ccy"`
	Interest      string `json:"interest"`
	LoanAlloc     string `json:"loanAlloc"`
	Rate          string `json:"rate"`
	SurplusLmt    string `json:"surplusLmt"`
	UsedLmt       string `json:"usedLmt"`
}

type PositionHistory struct {
	CloseAvgPx   string  `json:"closeAvgPx"`
	CloseTotalPos string `json:"closeTotalPos"`
	CTime        string  `json:"cTime"`
	Direction    string  `json:"direction"`
	InstID       string  `json:"instId"`
	InstType     string  `json:"instType"`
	Lever        string  `json:"lever"`
	MgnMode      string  `json:"mgnMode"`
	OpenAvgPx    string  `json:"openAvgPx"`
	OpenMaxPos   string  `json:"openMaxPos"`
	Pnl          string  `json:"pnl"`
	PnlRatio     string  `json:"pnlRatio"`
	PosID        string  `json:"posId"`
	PosSide      string  `json:"posSide"`
	TriggerPx    *string `json:"triggerPx,omitempty"`
	Type         string  `json:"type"`
	UTime        string  `json:"uTime"`
}

type AccountLevel struct {
	Level    string `json:"level"`
	AutoLoan bool   `json:"autoLoan"`
}

type MMPConfig struct {
	InstFamily string `json:"instFamily"`
	TimeInterval string `json:"timeInterval"`
	FrozenInterval string `json:"frozenInterval"`
	QtyLimit string `json:"qtyLimit"`
}

type MMPState struct {
	InstFamily string `json:"instFamily"`
	Frozen bool `json:"frozen"`
	FrozenUntil string `json:"frozenUntil,omitempty"`
}

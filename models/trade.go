package models

type PlaceOrderRequest struct {
	InstID       string  `json:"instId"`
	TdMode       string  `json:"tdMode"`
	Side         string  `json:"side"`
	OrdType      string  `json:"ordType"`
	Sz           string  `json:"sz"`
	Ccy          *string `json:"ccy,omitempty"`
	ClOrdID      *string `json:"clOrdId,omitempty"`
	Tag          *string `json:"tag,omitempty"`
	PosSide      *string `json:"posSide,omitempty"`
	Px           *string `json:"px,omitempty"`
	ReduceOnly   *bool   `json:"reduceOnly,omitempty"`
	TgtCcy       *string `json:"tgtCcy,omitempty"`
	BanAmend     *bool   `json:"banAmend,omitempty"`
	TpTriggerPx  *string `json:"tpTriggerPx,omitempty"`
	TpOrdPx      *string `json:"tpOrdPx,omitempty"`
	SlTriggerPx  *string `json:"slTriggerPx,omitempty"`
	SlOrdPx      *string `json:"slOrdPx,omitempty"`
	TpTriggerPxType *string `json:"tpTriggerPxType,omitempty"`
	SlTriggerPxType *string `json:"slTriggerPxType,omitempty"`
	QuickMgnType *string `json:"quickMgnType,omitempty"`
	StpID        *string `json:"stpId,omitempty"`
	StpMode      *string `json:"stpMode,omitempty"`
	AttachAlgoClOrdID *string `json:"attachAlgoClOrdId,omitempty"`
}

type PlaceOrderResponse struct {
	ClOrdID string `json:"clOrdId"`
	OrdID   string `json:"ordId"`
	Tag     string `json:"tag"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}

type CancelOrderRequest struct {
	InstID  string  `json:"instId"`
	OrdID   *string `json:"ordId,omitempty"`
	ClOrdID *string `json:"clOrdId,omitempty"`
}

type CancelOrderResponse struct {
	ClOrdID string `json:"clOrdId"`
	OrdID   string `json:"ordId"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}

type AmendOrderRequest struct {
	InstID     string  `json:"instId"`
	CxlOnFail  *bool   `json:"cxlOnFail,omitempty"`
	OrdID      *string `json:"ordId,omitempty"`
	ClOrdID    *string `json:"clOrdId,omitempty"`
	ReqID      *string `json:"reqId,omitempty"`
	NewSz      *string `json:"newSz,omitempty"`
	NewPx      *string `json:"newPx,omitempty"`
	NewTpTriggerPx *string `json:"newTpTriggerPx,omitempty"`
	NewTpOrdPx *string `json:"newTpOrdPx,omitempty"`
	NewSlTriggerPx *string `json:"newSlTriggerPx,omitempty"`
	NewSlOrdPx *string `json:"newSlOrdPx,omitempty"`
	NewTpTriggerPxType *string `json:"newTpTriggerPxType,omitempty"`
	NewSlTriggerPxType *string `json:"newSlTriggerPxType,omitempty"`
}

type AmendOrderResponse struct {
	ClOrdID string `json:"clOrdId"`
	OrdID   string `json:"ordId"`
	ReqID   string `json:"reqId"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}

type Order struct {
	AccFillSz        string  `json:"accFillSz"`
	AlgoClOrdID      *string `json:"algoClOrdId,omitempty"`
	AlgoID           *string `json:"algoId,omitempty"`
	AttachAlgoClOrdID *string `json:"attachAlgoClOrdId,omitempty"`
	AttachAlgoOrds   []AttachAlgoOrder `json:"attachAlgoOrds,omitempty"`
	AvgPx            string  `json:"avgPx"`
	CTime            string  `json:"cTime"`
	Category         string  `json:"category"`
	Ccy              string  `json:"ccy"`
	ClOrdID          string  `json:"clOrdId"`
	Fee              string  `json:"fee"`
	FeeCcy           string  `json:"feeCcy"`
	FillPx           string  `json:"fillPx"`
	FillSz           string  `json:"fillSz"`
	FillTime         string  `json:"fillTime"`
	InstID           string  `json:"instId"`
	InstType         string  `json:"instType"`
	Lever            string  `json:"lever"`
	OrdID            string  `json:"ordId"`
	OrdType          string  `json:"ordType"`
	Pnl              string  `json:"pnl"`
	PosSide          string  `json:"posSide"`
	Px               string  `json:"px"`
	PxType           *string `json:"pxType,omitempty"`
	PxUsd            *string `json:"pxUsd,omitempty"`
	PxVol            *string `json:"pxVol,omitempty"`
	Rebate           string  `json:"rebate"`
	RebateCcy        string  `json:"rebateCcy"`
	ReduceOnly       string  `json:"reduceOnly"`
	Side             string  `json:"side"`
	SlOrdPx          string  `json:"slOrdPx"`
	SlTriggerPx      string  `json:"slTriggerPx"`
	SlTriggerPxType  string  `json:"slTriggerPxType"`
	Source           string  `json:"source"`
	State            string  `json:"state"`
	StpID            *string `json:"stpId,omitempty"`
	StpMode          *string `json:"stpMode,omitempty"`
	Sz               string  `json:"sz"`
	Tag              string  `json:"tag"`
	TdMode           string  `json:"tdMode"`
	TgtCcy           string  `json:"tgtCcy"`
	TpOrdPx          string  `json:"tpOrdPx"`
	TpTriggerPx      string  `json:"tpTriggerPx"`
	TpTriggerPxType  string  `json:"tpTriggerPxType"`
	TradeID          string  `json:"tradeId"`
	UTime            string  `json:"uTime"`
	QuickMgnType     *string `json:"quickMgnType,omitempty"`
}

type AttachAlgoOrder struct {
	AttachAlgoClOrdID string  `json:"attachAlgoClOrdId"`
	AttachAlgoID      *string `json:"attachAlgoId,omitempty"`
	TpOrdPx           string  `json:"tpOrdPx"`
	TpTriggerPx       string  `json:"tpTriggerPx"`
	TpTriggerPxType   string  `json:"tpTriggerPxType"`
	SlOrdPx           string  `json:"slOrdPx"`
	SlTriggerPx       string  `json:"slTriggerPx"`
	SlTriggerPxType   string  `json:"slTriggerPxType"`
	Sz                string  `json:"sz"`
	AmendPxOnTriggerType *string `json:"amendPxOnTriggerType,omitempty"`
}

type Fill struct {
	InstID      string  `json:"instId"`
	InstType    string  `json:"instType"`
	TradeID     string  `json:"tradeId"`
	OrdID       string  `json:"ordId"`
	ClOrdID     string  `json:"clOrdId"`
	BillID      string  `json:"billId"`
	Tag         string  `json:"tag"`
	FillPx      string  `json:"fillPx"`
	FillSz      string  `json:"fillSz"`
	FillPnl     string  `json:"fillPnl"`
	Side        string  `json:"side"`
	PosSide     string  `json:"posSide"`
	ExecType    string  `json:"execType"`
	FeeCcy      string  `json:"feeCcy"`
	Fee         string  `json:"fee"`
	TS          string  `json:"ts"`
	FillTime    string  `json:"fillTime"`
	TdMode      string  `json:"tdMode"`
}

type AlgoOrder struct {
	ActualPx         string  `json:"actualPx"`
	ActualSide       string  `json:"actualSide"`
	ActualSz         string  `json:"actualSz"`
	AlgoClOrdID      *string `json:"algoClOrdId,omitempty"`
	AlgoID           string  `json:"algoId"`
	Ccy              string  `json:"ccy"`
	CTime            string  `json:"cTime"`
	InstID           string  `json:"instId"`
	InstType         string  `json:"instType"`
	Lever            string  `json:"lever"`
	OrdID            *string `json:"ordId,omitempty"`
	OrdPx            string  `json:"ordPx"`
	OrdType          string  `json:"ordType"`
	PosSide          string  `json:"posSide"`
	Px               string  `json:"px"`
	PxLimit          *string `json:"pxLimit,omitempty"`
	PxSpread         *string `json:"pxSpread,omitempty"`
	PxVar            *string `json:"pxVar,omitempty"`
	Side             string  `json:"side"`
	SlOrdPx          *string `json:"slOrdPx,omitempty"`
	SlTriggerPx      *string `json:"slTriggerPx,omitempty"`
	SlTriggerPxType  *string `json:"slTriggerPxType,omitempty"`
	State            string  `json:"state"`
	Sz               string  `json:"sz"`
	SzLimit          *string `json:"szLimit,omitempty"`
	Tag              string  `json:"tag"`
	TdMode           string  `json:"tdMode"`
	TgtCcy           string  `json:"tgtCcy"`
	TimeInterval     *string `json:"timeInterval,omitempty"`
	TpOrdPx          *string `json:"tpOrdPx,omitempty"`
	TpTriggerPx      *string `json:"tpTriggerPx,omitempty"`
	TpTriggerPxType  *string `json:"tpTriggerPxType,omitempty"`
	TriggerPx        string  `json:"triggerPx"`
	TriggerPxType    string  `json:"triggerPxType"`
	TriggerTime      string  `json:"triggerTime"`
	UTime            string  `json:"uTime"`
	CallbackRatio    *string `json:"callbackRatio,omitempty"`
	CallbackSpread   *string `json:"callbackSpread,omitempty"`
	ActivePx         *string `json:"activePx,omitempty"`
	MoveTriggerPx    *string `json:"moveTriggerPx,omitempty"`
}

type PlaceAlgoOrderRequest struct {
	InstID          string  `json:"instId"`
	TdMode          string  `json:"tdMode"`
	Side            string  `json:"side"`
	OrdType         string  `json:"ordType"`
	Sz              string  `json:"sz"`
	Ccy             *string `json:"ccy,omitempty"`
	PosSide         *string `json:"posSide,omitempty"`
	ReduceOnly      *bool   `json:"reduceOnly,omitempty"`
	TpTriggerPx     *string `json:"tpTriggerPx,omitempty"`
	TpOrdPx         *string `json:"tpOrdPx,omitempty"`
	SlTriggerPx     *string `json:"slTriggerPx,omitempty"`
	SlOrdPx         *string `json:"slOrdPx,omitempty"`
	TriggerPx       *string `json:"triggerPx,omitempty"`
	OrderPx         *string `json:"orderPx,omitempty"`
	TpTriggerPxType *string `json:"tpTriggerPxType,omitempty"`
	SlTriggerPxType *string `json:"slTriggerPxType,omitempty"`
	TriggerPxType   *string `json:"triggerPxType,omitempty"`
	TgtCcy          *string `json:"tgtCcy,omitempty"`
	Tag             *string `json:"tag,omitempty"`
	AlgoClOrdID     *string `json:"algoClOrdId,omitempty"`
	QuickMgnType    *string `json:"quickMgnType,omitempty"`
}

type PlaceAlgoOrderResponse struct {
	AlgoID      string `json:"algoId"`
	AlgoClOrdID string `json:"algoClOrdId"`
	SCode       string `json:"sCode"`
	SMsg        string `json:"sMsg"`
}

type CancelAlgoOrderRequest struct {
	AlgoID      *string `json:"algoId,omitempty"`
	AlgoClOrdID *string `json:"algoClOrdId,omitempty"`
	InstID      string  `json:"instId"`
}

type CancelAlgoOrderResponse struct {
	AlgoID      string `json:"algoId"`
	AlgoClOrdID string `json:"algoClOrdId"`
	SCode       string `json:"sCode"`
	SMsg        string `json:"sMsg"`
}

type ClosePositionRequest struct {
	InstID   string  `json:"instId"`
	MgnMode  string  `json:"mgnMode"`
	PosSide  *string `json:"posSide,omitempty"`
	Ccy      *string `json:"ccy,omitempty"`
	AutoCancel *bool `json:"autoCxl,omitempty"`
	ClOrdID  *string `json:"clOrdId,omitempty"`
	Tag      *string `json:"tag,omitempty"`
}

type MassCancelRequest struct {
	InstType   string  `json:"instType"`
	InstFamily *string `json:"instFamily,omitempty"`
}

type MassCancelResponse struct {
	Result bool   `json:"result"`
	SCode  string `json:"sCode"`
	SMsg   string `json:"sMsg"`
}

type CancelAllAfterRequest struct {
	TimeOut string `json:"timeOut"`
}

type CancelAllAfterResponse struct {
	TriggerTime string `json:"triggerTime"`
	TS          string `json:"ts"`
}

type EasyConvertRequest struct {
	FromCcy []string `json:"fromCcy"`
	ToCcy   string   `json:"toCcy"`
}

type EasyConvertResponse struct {
	FillFromSz string `json:"fillFromSz"`
	FillToSz   string `json:"fillToSz"`
	FromCcy    string `json:"fromCcy"`
	Status     string `json:"status"`
	ToCcy      string `json:"toCcy"`
	UTime      string `json:"uTime"`
}

type OneClickRepayRequest struct {
	DebtCcy []string `json:"debtCcy"`
	RepayCcy string  `json:"repayCcy"`
}

type OneClickRepayResponse struct {
	DebtCcy  string `json:"debtCcy"`
	FillDebtSz string `json:"fillDebtSz"`
	FillRepaySz string `json:"fillRepaySz"`
	RepayCcy string `json:"repayCcy"`
	Status   string `json:"status"`
	UTime    string `json:"uTime"`
}

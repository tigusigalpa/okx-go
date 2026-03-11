package models

type AssetBalance struct {
	AvailBal string `json:"availBal"`
	Bal      string `json:"bal"`
	Ccy      string `json:"ccy"`
	FrozenBal string `json:"frozenBal"`
}

type AssetValuation struct {
	Details      []AssetValuationDetail `json:"details"`
	TotalBal     string                 `json:"totalBal"`
	TS           string                 `json:"ts"`
}

type AssetValuationDetail struct {
	Ccy    string `json:"ccy"`
	Bal    string `json:"bal"`
	ValBal string `json:"valBal"`
}

type TransferRequest struct {
	Ccy      string  `json:"ccy"`
	Amt      string  `json:"amt"`
	From     string  `json:"from"`
	To       string  `json:"to"`
	Type     *string `json:"type,omitempty"`
	SubAcct  *string `json:"subAcct,omitempty"`
	InstID   *string `json:"instId,omitempty"`
	ToInstID *string `json:"toInstId,omitempty"`
	LoanTrans *bool  `json:"loanTrans,omitempty"`
	OmitPosRisk *string `json:"omitPosRisk,omitempty"`
	ClientID *string `json:"clientId,omitempty"`
}

type TransferResponse struct {
	TransID  string `json:"transId"`
	Ccy      string `json:"ccy"`
	ClientID string `json:"clientId"`
	From     string `json:"from"`
	Amt      string `json:"amt"`
	To       string `json:"to"`
}

type TransferState struct {
	TransID  string  `json:"transId"`
	Ccy      string  `json:"ccy"`
	ClientID *string `json:"clientId,omitempty"`
	From     string  `json:"from"`
	Amt      string  `json:"amt"`
	To       string  `json:"to"`
	State    string  `json:"state"`
	Type     *string `json:"type,omitempty"`
	SubAcct  *string `json:"subAcct,omitempty"`
	InstID   *string `json:"instId,omitempty"`
	ToInstID *string `json:"toInstId,omitempty"`
}

type WithdrawalRequest struct {
	Ccy      string  `json:"ccy"`
	Amt      string  `json:"amt"`
	Dest     string  `json:"dest"`
	ToAddr   string  `json:"toAddr"`
	Fee      string  `json:"fee"`
	Chain    *string `json:"chain,omitempty"`
	AreaCode *string `json:"areaCode,omitempty"`
	ClientID *string `json:"clientId,omitempty"`
}

type WithdrawalResponse struct {
	Amt      string `json:"amt"`
	WdID     string `json:"wdId"`
	Ccy      string `json:"ccy"`
	ClientID string `json:"clientId"`
	Chain    string `json:"chain"`
}

type WithdrawalHistory struct {
	Chain    string  `json:"chain"`
	Fee      string  `json:"fee"`
	Ccy      string  `json:"ccy"`
	ClientID *string `json:"clientId,omitempty"`
	Amt      string  `json:"amt"`
	TxID     string  `json:"txId"`
	From     string  `json:"from"`
	To       string  `json:"to"`
	State    string  `json:"state"`
	TS       string  `json:"ts"`
	WdID     string  `json:"wdId"`
}

type DepositAddress struct {
	Addr     string           `json:"addr"`
	Tag      *string          `json:"tag,omitempty"`
	Memo     *string          `json:"memo,omitempty"`
	PmtID    *string          `json:"pmtId,omitempty"`
	Ccy      string           `json:"ccy"`
	Chain    string           `json:"chain"`
	To       string           `json:"to"`
	Selected bool             `json:"selected"`
	CtAddr   string           `json:"ctAddr"`
	TS       string           `json:"ts"`
}

type DepositHistory struct {
	ActualDepBlkConfirm string  `json:"actualDepBlkConfirm"`
	Amt                 string  `json:"amt"`
	Ccy                 string  `json:"ccy"`
	Chain               string  `json:"chain"`
	DepID               string  `json:"depId"`
	From                string  `json:"from"`
	FromWdID            *string `json:"fromWdId,omitempty"`
	State               string  `json:"state"`
	To                  string  `json:"to"`
	TS                  string  `json:"ts"`
	TxID                string  `json:"txId"`
}

type Currency struct {
	CanDep       bool   `json:"canDep"`
	CanInternal  bool   `json:"canInternal"`
	CanWd        bool   `json:"canWd"`
	Ccy          string `json:"ccy"`
	Chain        string `json:"chain"`
	LogoLink     string `json:"logoLink"`
	MainNet      bool   `json:"mainNet"`
	MaxFee       string `json:"maxFee"`
	MaxWd        string `json:"maxWd"`
	MinDep       string `json:"minDep"`
	MinDepArrivalConfirm string `json:"minDepArrivalConfirm"`
	MinFee       string `json:"minFee"`
	MinWd        string `json:"minWd"`
	MinWdUnlockConfirm string `json:"minWdUnlockConfirm"`
	Name         string `json:"name"`
	NeedTag      bool   `json:"needTag"`
	UsedDepQuotaFixed string `json:"usedDepQuotaFixed"`
	UsedWdQuota  string `json:"usedWdQuota"`
	WdQuota      string `json:"wdQuota"`
	WdTickSz     string `json:"wdTickSz"`
}

type AssetBill struct {
	BillID   string  `json:"billId"`
	Ccy      string  `json:"ccy"`
	ClientID *string `json:"clientId,omitempty"`
	BalChg   string  `json:"balChg"`
	Bal      string  `json:"bal"`
	Type     string  `json:"type"`
	TS       string  `json:"ts"`
}

type SavingBalance struct {
	Ccy      string `json:"ccy"`
	Amt      string `json:"amt"`
	Earnings string `json:"earnings"`
	Rate     string `json:"rate"`
}

type PurchaseRedemptRequest struct {
	Ccy  string `json:"ccy"`
	Amt  string `json:"amt"`
	Side string `json:"side"`
	Rate string `json:"rate"`
}

type PurchaseRedemptResponse struct {
	Ccy  string `json:"ccy"`
	Amt  string `json:"amt"`
	Side string `json:"side"`
	Rate string `json:"rate"`
}

type LendingRate struct {
	Ccy  string `json:"ccy"`
	Rate string `json:"rate"`
}

type ConvertCurrencyPair struct {
	InstID   string `json:"instId"`
	BaseCcy  string `json:"baseCcy"`
	BaseCcyMax string `json:"baseCcyMax"`
	BaseCcyMin string `json:"baseCcyMin"`
	QuoteCcy string `json:"quoteCcy"`
	QuoteCcyMax string `json:"quoteCcyMax"`
	QuoteCcyMin string `json:"quoteCcyMin"`
}

type ConvertEstimateQuoteRequest struct {
	BaseCcy   string  `json:"baseCcy"`
	QuoteCcy  string  `json:"quoteCcy"`
	Side      string  `json:"side"`
	RfqSz     string  `json:"rfqSz"`
	RfqSzCcy  string  `json:"rfqSzCcy"`
	ClQReqID  *string `json:"clQReqId,omitempty"`
	Tag       *string `json:"tag,omitempty"`
}

type ConvertEstimateQuoteResponse struct {
	BaseCcy    string `json:"baseCcy"`
	BaseSz     string `json:"baseSz"`
	ClQReqID   string `json:"clQReqId"`
	CnvtPx     string `json:"cnvtPx"`
	OrigRfqSz  string `json:"origRfqSz"`
	QuoteCcy   string `json:"quoteCcy"`
	QuoteID    string `json:"quoteId"`
	QuoteSz    string `json:"quoteSz"`
	QuoteTime  string `json:"quoteTime"`
	RfqSz      string `json:"rfqSz"`
	RfqSzCcy   string `json:"rfqSzCcy"`
	Side       string `json:"side"`
	TTLMs      string `json:"ttlMs"`
}

type ConvertTradeRequest struct {
	QuoteID  string  `json:"quoteId"`
	BaseCcy  string  `json:"baseCcy"`
	QuoteCcy string  `json:"quoteCcy"`
	Side     string  `json:"side"`
	Sz       string  `json:"sz"`
	SzCcy    string  `json:"szCcy"`
	ClTReqID *string `json:"clTReqId,omitempty"`
	Tag      *string `json:"tag,omitempty"`
}

type ConvertTradeResponse struct {
	BaseCcy   string `json:"baseCcy"`
	ClTReqID  string `json:"clTReqId"`
	FillBaseSz string `json:"fillBaseSz"`
	FillPx    string `json:"fillPx"`
	FillQuoteSz string `json:"fillQuoteSz"`
	InstID    string `json:"instId"`
	QuoteCcy  string `json:"quoteCcy"`
	QuoteID   string `json:"quoteId"`
	Side      string `json:"side"`
	State     string `json:"state"`
	TradeID   string `json:"tradeId"`
	TS        string `json:"ts"`
}

type MonthlyStatement struct {
	Year  string `json:"year"`
	Month string `json:"month"`
	URL   string `json:"url"`
}

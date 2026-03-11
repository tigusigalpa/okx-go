package models

type SystemStatus struct {
	Title      string              `json:"title"`
	State      string              `json:"state"`
	Begin      string              `json:"begin"`
	End        string              `json:"end"`
	Href       string              `json:"href"`
	ServiceType string             `json:"serviceType"`
	System     string              `json:"system"`
	ScheDesc   string              `json:"scheDesc"`
	PushTime   string              `json:"pushTime"`
	TS         string              `json:"ts"`
}

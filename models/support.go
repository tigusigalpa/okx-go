package models

// AnnouncementType represents an OKX API request or response value.
type AnnouncementType struct {
	AnnID   string `json:"annId"`
	AnnType string `json:"annType"`
}

// Announcement represents an OKX API request or response value.
type Announcement struct {
	AnnID    string `json:"annId"`
	AnnTitle string `json:"annTitle"`
	AnnType  string `json:"annType"`
	PushTime string `json:"pushTime"`
	AnnDesc  string `json:"annDesc"`
}

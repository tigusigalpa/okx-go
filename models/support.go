package models

type AnnouncementType struct {
	AnnID   string `json:"annId"`
	AnnType string `json:"annType"`
}

type Announcement struct {
	AnnID     string `json:"annId"`
	AnnTitle  string `json:"annTitle"`
	AnnType   string `json:"annType"`
	PushTime  string `json:"pushTime"`
	AnnDesc   string `json:"annDesc"`
}

package models

type Activity struct {
	Id           string `json:"id"`
	UserId       string `json:"userId"`
	Description  string `json:"description"`
	RecordedTime string `json:"rt"`
}

// API responses below
type ActivityResponse struct {
	Code       int        `json:"code"`
	Text       string     `json:"text"`
	Activities []Activity `json:"activities"`
}

package models

type Activity struct {
	Id           string `json:"id"`
	UserId       string `json:"userId"`
	Description  string `json:"description"`
	RecordedTime string `json:"rt"`
}

type ActivityDetail struct {
	UserId       string `json:"userId"`
	ActivityId   string `json:"activityId"`
	Description  string `json:"description"`
	ImageUrl     string `json:"imageUrl"`
	Lat          string `json:"lat"`
	Lon          string `json:"lon"`
	RecordedTime int64  `json:"rt"`
}

// API responses below
type ActivityResponse struct {
	Code       int        `json:"code"`
	Text       string     `json:"text"`
	Activities []Activity `json:"activities"`
}

type ActivityDetailResponse struct {
	Code            int              `json:"code"`
	Text            string           `json:"text"`
	ActivityDetails []ActivityDetail `json:"activities"`
}

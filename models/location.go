package models

type Location struct {
	Lat          string `json:"lat"`
	Lon          string `json:"lon"`
	RecordedTime int64  `json:"rt"`
}

// API responses below
type UserLocationResponse struct {
	Code      int        `json:"code"`
	Text      string     `json:"text"`
	Locations []Location `json:"locations"`
}

type LocationCaptureResponse struct {
	Code       int    `json:"code"`
	Text       string `json:"text"`
	IsCaptured bool   `json:"isCaptured"`
}

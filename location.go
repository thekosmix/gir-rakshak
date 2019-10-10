package main

type Location struct {
	Id           int64  `json:"id"`
	UserId       int    `json:"userId"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	RecordedTime int64  `json:"recordedTime"`
	CreatedDate  int64  `json:"createdDate"`
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

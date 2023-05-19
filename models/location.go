package models

import "gir-rakshak/utils"

type Location struct {
	Lat          string `json:"lat"`
	Lon          string `json:"lon"`
	RecordedTime int64  `json:"rt"`
}

type UserLocationResponse struct {
	utils.BaseResponse
	Locations []Location `json:"locations"`
}

type LocationCaptureResponse struct {
	utils.BaseResponse
	IsCaptured bool   `json:"isCaptured"`
}

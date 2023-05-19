package models

import "gir-rakshak/utils"

type Activity struct {
	Id           int    `json:"id"`
	UserId       int    `json:"userId"`
	Description  string `json:"description"`
	RecordedTime int64  `json:"rt"`
	Lat          string `json:"lat"`
	Lon          string `json:"lon"`
}

type ActivityDetail struct {
	UserId       string `json:"userId"`
	ActivityId   int    `json:"activityId"`
	Description  string `json:"description"`
	ImageUrl     string `json:"imageUrl"`
	Lat          string `json:"lat"`
	Lon          string `json:"lon"`
	RecordedTime int64  `json:"rt"`
}

type ActivityResponse struct {
	utils.BaseResponse
	Activities []Activity `json:"activities"`
}

type ActivityDetailResponse struct {
	utils.BaseResponse
	ActivityDetails []ActivityDetail `json:"activities"`
}

type AddActivityResponse struct {
	utils.BaseResponse
	IsAdded bool   `json:"isAdded"`
}

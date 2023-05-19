package models

import "gir-rakshak/utils"

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	AccessToken string `json:"at"`
	Role        string `json:"role"`
	DeviceId    string `json:"deviceId"`
	CreatedDate int64  `json:"createdDate"`
	IsActive    bool   `json:"isActive"`
}

type ApproveUserRequest struct {
	Id         int  `json:"id"`
	IsApproved bool `json:"isApproved"`
}

type LoginUserRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	DeviceId    string `json:"deviceId"`
}

// API responses below
type AllUserResponse struct {
	utils.BaseResponse
	Users []User `json:"users"`
}

type ApproveUserResponse struct {
	utils.BaseResponse
	IsSuccess bool   `json:"isSuccess"`
}

type RegisterUserResponse struct {
	utils.BaseResponse
	IsRegistered bool   `json:"isRegistered"`
}

type LoginUserResponse struct {
	utils.BaseResponse
	User 	User   `json:"user"`
}

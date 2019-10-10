package main

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	DeviceId    string `json:"deviceId"`
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
	Code  int    `json:"code"`
	Text  string `json:"text"`
	Users []User `json:"users"`
}

type ApproveUserResponse struct {
	Code      int    `json:"code"`
	Text      string `json:"text"`
	IsSuccess bool   `json:"isSuccess"`
}

type RegisterUserResponse struct {
	Code         int    `json:"code"`
	Text         string `json:"text"`
	IsRegistered bool   `json:"isRegistered"`
}

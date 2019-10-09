package main

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	PhoneNumber string  `json:"phoneNumber"`
	Password  string 	`json:"password"`
	Role	  string 	`json:"role"`
	DeviceId  string 	`json:"deviceId"`
	CreatedDate int64 	`json:"createdDate"`
	LastUpdatedDate int64 	`json:"lastUpdatedDate"`
	IsActive  bool      `json:"isActive"`
}

type UserApproved struct {
	Id        int       `json:"id"`
	IsApproved  bool      `json:"isApproved"`	
}


// API responses below 
type AllUserResponse struct {
	Code        int      `json:"code"`
	Msg			string   `json:"msg"`
	Users 		[]User 	 `json:"users"`
}

type ApproveUserResponse struct {
	Code        int      `json:"code"`
	Msg			string   `json:"msg"`
	IsSuccess 	bool 	 `json:"IsSuccess"`
}

type RegisterUserResponse struct {
	Code        int      `json:"code"`
	Msg			string   `json:"msg"`
	IsRegistered 	bool `json:"isRegistered"`
}
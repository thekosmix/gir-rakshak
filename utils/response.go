package utils

type BaseResponse struct {
	Code int    `json:"status"`
	Text string `json:"error"`
}

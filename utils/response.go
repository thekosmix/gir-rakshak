package utils

import (
	"encoding/json"
	"net/http"
)
/**
0. success - 200
1. request not valid -
2. data not available - requested data is not present
3. server error - redis/db connection issue
**/
type BaseResponse struct {
	Code int    `json:"status"`
	Text string `json:"error"`
}

func SetResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func SetErroneousResponse(w http.ResponseWriter, err error) {
	json.NewEncoder(w).Encode(BaseResponse{Code: 1, Text: err.Error()})
}
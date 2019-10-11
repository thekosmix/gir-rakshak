package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func SetResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func setErroneousResponse(w http.ResponseWriter, err error) {
	json.NewEncoder(w).Encode(jsonErr{Code: 1, Text: err.Error()})
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w)
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		setErroneousResponse(w, err)
		return
	}

	t, err := RepoRegisterUser(user)
	if err != nil {
		setErroneousResponse(w, err)
		return
	}
	response := RegisterUserResponse{0, "", t}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		setErroneousResponse(w, err)
		panic(err)
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w)
	var request LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		setErroneousResponse(w, err)
		return
	}

	user, err := RepoLoginUser(request.PhoneNumber, request.DeviceId, request.Password)

	if err != nil {
		setErroneousResponse(w, err)
		return
	}
	_, err = AddUserToken(user.Id, user.AccessToken)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		setErroneousResponse(w, err)
		panic(err)
	}

}

func UploadUserLocation(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w)
	var locations []Location
	err := json.NewDecoder(r.Body).Decode(&locations)
	if err != nil {
		setErroneousResponse(w, err)
		return
	}
	uid := r.Header.Get("uid")
	uidInt, _ := strconv.Atoi(uid)

	isUploaded, err := RepoUploadUserLocation(locations, uidInt)

	if err != nil {
		setErroneousResponse(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(LocationCaptureResponse{0, "", isUploaded}); err != nil {
		setErroneousResponse(w, err)
		panic(err)
	}

}

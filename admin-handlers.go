package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Gir-rakshak!\n")
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	users := RepoAllUser()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if(len(users) > 0) {
		response := AllUserResponse{0, "", users}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
		return
	}
	// If we didn't find it, 404
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func ApproveUser(w http.ResponseWriter, r *http.Request) {
	var userApproved UserApproved
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.Unmarshal(body, &userApproved); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoApproveUser(userApproved)
	response := ApproveUserResponse{0,"", t}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func UserLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var userId int
	var err error

	if userId, err = strconv.Atoi(vars["userId"]); err != nil {
		panic(err)
	}

	fromTime := GetTime("fromTime", r)
	toTime := GetTime("toTime", r)

	locations := RepoUserLocation(userId, fromTime, toTime)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if(len(locations) > 0) {
		response := UserLocationResponse{0, "", locations}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
		return
	}
	// If we didn't find it, 404
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func GetTime(timeFormat string, r *http.Request) int64 {
	var requestTime int64
	var err error

	timeArr, ok := r.URL.Query()[timeFormat]
	if !ok || len(timeArr[0]) < 1 {
		panic("toTime is missing")
	}
	if requestTime, err = strconv.ParseInt(timeArr[0], 10, 64); err != nil {
		panic(err)
	}

	return requestTime;
}

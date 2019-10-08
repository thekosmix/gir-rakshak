package main

import (
	"encoding/json"
//	"fmt"
//	"log"
//	"io"
//	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func UserLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var userId int
	var fromTime int64
	var toTime int64
	var err error

	if userId, err = strconv.Atoi(vars["userId"]); err != nil {
		panic(err)
	}

	fromTimeArr, ok := r.URL.Query()["fromTime"]
	if !ok || len(fromTimeArr[0]) < 1 {
		panic("fromTime is missing")
	}
	if fromTime, err = strconv.ParseInt(fromTimeArr[0], 10, 64); err != nil {
		panic(err)
	}
	
	toTimeArr, ok := r.URL.Query()["toTime"]
	if !ok || len(toTimeArr[0]) < 1 {
		panic("toTime is missing")
	}
	if toTime, err = strconv.ParseInt(toTimeArr[0], 10, 64); err != nil {
		panic(err)
	}

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

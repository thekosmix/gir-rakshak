package main

import (
	"encoding/json"
//	"fmt"
//	"io"
//	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func UserLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var userId int
//	var fromTime int64
//	var toTime int64
	var err error
	if userId, err = strconv.Atoi(vars["userId"]); err != nil {
		panic(err)
	}

	locations := RepoUserLocation(userId)
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

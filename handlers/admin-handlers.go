package handlers

import (
	"encoding/json"
	"fmt"
	"gir-rakshak/models"
	"gir-rakshak/repo"
	"gir-rakshak/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Gir-rakshak!\n")
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	users := repo.RepoAllUser()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if len(users) > 0 {
		response := models.AllUserResponse{0, "", users}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf(err.Error())
		}
		return
	}
	// If we didn't find it, 404
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(utils.JsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		log.Printf(err.Error())
	}
}

func ApproveUser(w http.ResponseWriter, r *http.Request) {
	var userApproved models.ApproveUserRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Printf(err.Error())
	}
	if err := r.Body.Close(); err != nil {
		log.Printf(err.Error())
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.Unmarshal(body, &userApproved); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Printf(err.Error())
		}
	}

	t, err := repo.RepoApproveUser(userApproved)
	if err != nil {
		SetErroneousResponse(w, err)
		log.Printf(err.Error())
	}
	response := models.ApproveUserResponse{0, "", t}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf(err.Error())
	}
}

func UserLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var userId int
	var err error

	if userId, err = strconv.Atoi(vars["userId"]); err != nil {
		log.Printf(err.Error())
	}

	fromTime := GetTime("fromTime", r)
	toTime := GetTime("toTime", r)

	locations := repo.RepoUserLocation(userId, fromTime, toTime)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if len(locations) > 0 {
		response := models.UserLocationResponse{0, "", locations}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf(err.Error())
		}
		return
	}
	// If we didn't find it, 404
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(utils.JsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		log.Printf(err.Error())
	}
}

func ViewAllActivity(w http.ResponseWriter, r *http.Request) {

	fromTime := GetTime("fromTime", r)
	toTime := GetTime("toTime", r)

	activities := repo.RepoGetActivity(fromTime, toTime)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if len(activities) > 0 {
		response := models.ActivityResponse{0, "", activities}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf(err.Error())
		}
		return
	}
	// If we didn't find it, 404
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(utils.JsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		log.Printf(err.Error())
	}
}

func ViewActivityDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var activityId int64
	var err error

	if activityId, err = strconv.ParseInt(vars["activityId"], 10, 64); err != nil {
		log.Printf(err.Error())
	}

	activityDetails := repo.RepoGetActivityDetail(activityId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if len(activityDetails) > 0 {
		response := models.ActivityDetailResponse{0, "", activityDetails}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf(err.Error())
		}
		return
	}
	// If we didn't find it, 404
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(utils.JsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		log.Printf(err.Error())
	}
}

func GetTime(timeFormat string, r *http.Request) int64 {
	var requestTime int64
	var err error

	timeArr, ok := r.URL.Query()[timeFormat]
	if !ok || len(timeArr[0]) < 1 {
		log.Printf(timeFormat + " is missing")
	}
	if requestTime, err = strconv.ParseInt(timeArr[0], 10, 64); err != nil {
		log.Printf(err.Error())
	}

	return requestTime
}

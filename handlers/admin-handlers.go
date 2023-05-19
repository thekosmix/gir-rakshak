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
	
	utils.SetResponseHeaders(w);

	if len(users) > 0 {
		var response models.AllUserResponse
		response.Users = users

		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf(err.Error())
		}
		return
	}
	
	if err := json.NewEncoder(w).Encode(utils.BaseResponse{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
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
	utils.SetResponseHeaders(w);

	if err := json.Unmarshal(body, &userApproved); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Printf(err.Error())
		}
	}

	t, err := repo.RepoApproveUser(userApproved)
	if err != nil {
		utils.SetErroneousResponse(w, err)
		log.Printf(err.Error())
	}

	var response models.ApproveUserResponse
	response.IsSuccess = t

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
	
	utils.SetResponseHeaders(w);

	if len(locations) > 0 {

		var response models.UserLocationResponse
		response.Locations = locations

		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf(err.Error())
		}
		return
	}
	
	if err := json.NewEncoder(w).Encode(utils.BaseResponse{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		log.Printf(err.Error())
	}
}

func ViewAllActivity(w http.ResponseWriter, r *http.Request) {

	fromTime := GetTime("fromTime", r)
	toTime := GetTime("toTime", r)

	activities := repo.RepoGetActivity(fromTime, toTime)
	
	utils.SetResponseHeaders(w);

	if len(activities) > 0 {

		var response models.ActivityResponse
		response.Activities = activities

		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf(err.Error())
		}
		return
	}
	
	if err := json.NewEncoder(w).Encode(utils.BaseResponse{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
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
	
	utils.SetResponseHeaders(w);

	if len(activityDetails) > 0 {
		var response models.ActivityDetailResponse
		response.ActivityDetails = activityDetails
		
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf(err.Error())
		}
		return
	}
	
	if err := json.NewEncoder(w).Encode(utils.BaseResponse{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
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

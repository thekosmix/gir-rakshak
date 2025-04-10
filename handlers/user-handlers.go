package handlers

import (
	"encoding/json"
	"gir-rakshak/models"
	"gir-rakshak/repo"
	"gir-rakshak/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.SetErroneousResponse(w, err)
		return
	}

	t, err := repo.RepoRegisterUser(user)
	if err != nil {
		utils.SetErroneousResponse(w, err)
		return
	}

	var response models.RegisterUserResponse
	response.IsRegistered = t

	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.SetErroneousResponse(w, err)
		log.Printf(err.Error())
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	var request models.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.SetErroneousResponse(w, err)
		return
	}

	user, err := repo.RepoLoginUser(request.PhoneNumber, request.DeviceId, request.Password)

	if err != nil {
		utils.SetErroneousResponse(w, err)
		return
	}
	
	err = repo.AddUserToken(user.Id, user.AccessToken)

	var response models.LoginUserResponse
	response.User = user

	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.SetErroneousResponse(w, err)
		log.Printf(err.Error())
	}

}

func UploadUserLocation(w http.ResponseWriter, r *http.Request) {

	var locations []models.Location
	err := json.NewDecoder(r.Body).Decode(&locations)
	if err != nil {
		utils.SetErroneousResponse(w, err)
		return
	}
	uid := r.Header.Get("uid")
	uidInt, _ := strconv.Atoi(uid)

	isCaptured, err := repo.RepoUploadUserLocation(locations, uidInt)

	if err != nil {
		utils.SetErroneousResponse(w, err)
		return
	}

	var response models.LocationCaptureResponse
	response.IsCaptured = isCaptured

	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.SetErroneousResponse(w, err)
		log.Printf(err.Error())
	}
}

func AddActivity(w http.ResponseWriter, r *http.Request) {

	var activity models.Activity
	err := json.NewDecoder(r.Body).Decode(&activity)
	if err != nil {
		utils.SetErroneousResponse(w, err)
		return
	}
	uid := r.Header.Get("uid")
	uidInt, _ := strconv.Atoi(uid)
	log.Printf("%d", uidInt)
	isUploaded, err := repo.RepoAddActivity(activity, uidInt)

	if err != nil {
		utils.SetErroneousResponse(w, err)
		return
	}

	var response models.ApproveUserResponse
	response.IsSuccess = isUploaded

	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.SetErroneousResponse(w, err)
		log.Printf(err.Error())
	}
}

func ViewUserActivity(w http.ResponseWriter, r *http.Request) {

	fromTime := GetTime("fromTime", r)
	toTime := GetTime("toTime", r)
	uid := r.Header.Get("uid")
	uidInt, _ := strconv.Atoi(uid)

	activities := repo.RepoGetActivity(fromTime, toTime, uidInt)

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

func AddActivityDetail(w http.ResponseWriter, r *http.Request) {

	var activityDetail models.ActivityDetail
	err := json.NewDecoder(r.Body).Decode(&activityDetail)
	if err != nil {
		utils.SetErroneousResponse(w, err)
		return
	}

	vars := mux.Vars(r)
	var activityId int

	if activityId, err = strconv.Atoi(vars["activityId"]); err != nil {
		log.Printf(err.Error())
	}

	uid := r.Header.Get("uid")
	uidInt, _ := strconv.Atoi(uid)
	log.Printf("%d", uidInt)
	isAdded, err := repo.RepoAddActivityDetail(activityDetail, uidInt, activityId)

	if err != nil {
		utils.SetErroneousResponse(w, err)
		return
	}

	var response models.AddActivityResponse
	response.IsAdded = isAdded

	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.SetErroneousResponse(w, err)
		log.Printf(err.Error())
	}
}

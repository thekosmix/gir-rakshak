package handlers

import (
	"encoding/json"
	"gir-rakshak/models"
	"gir-rakshak/repo"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		SetErroneousResponse(w, err)
		return
	}

	t, err := repo.RepoRegisterUser(user)
	if err != nil {
		SetErroneousResponse(w, err)
		return
	}
	response := models.RegisterUserResponse{0, "", t}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		SetErroneousResponse(w, err)
		log.Printf(err.Error())
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w)
	var request models.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		SetErroneousResponse(w, err)
		return
	}

	user, err := repo.RepoLoginUser(request.PhoneNumber, request.DeviceId, request.Password)

	if err != nil {
		SetErroneousResponse(w, err)
		return
	}
	_, err = repo.AddUserToken(user.Id, user.AccessToken)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		SetErroneousResponse(w, err)
		log.Printf(err.Error())
	}

}

func UploadUserLocation(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w)
	var locations []models.Location
	err := json.NewDecoder(r.Body).Decode(&locations)
	if err != nil {
		SetErroneousResponse(w, err)
		return
	}
	uid := r.Header.Get("uid")
	uidInt, _ := strconv.Atoi(uid)

	isUploaded, err := repo.RepoUploadUserLocation(locations, uidInt)

	if err != nil {
		SetErroneousResponse(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(models.LocationCaptureResponse{0, "", isUploaded}); err != nil {
		SetErroneousResponse(w, err)
		log.Printf(err.Error())
	}
}

func AddActivity(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w)
	var activity models.Activity
	err := json.NewDecoder(r.Body).Decode(&activity)
	if err != nil {
		SetErroneousResponse(w, err)
		return
	}
	uid := r.Header.Get("uid")
	uidInt, _ := strconv.Atoi(uid)
	log.Printf("%d", uidInt)
	isUploaded, err := repo.RepoAddActivity(activity, uidInt)

	if err != nil {
		SetErroneousResponse(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(models.ApproveUserResponse{0, "", isUploaded}); err != nil {
		SetErroneousResponse(w, err)
		log.Printf(err.Error())
	}
}

func AddActivityDetail(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w)
	var activityDetail models.ActivityDetail
	err := json.NewDecoder(r.Body).Decode(&activityDetail)
	if err != nil {
		SetErroneousResponse(w, err)
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
		SetErroneousResponse(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(models.AddActivityResponse{0, "", isAdded}); err != nil {
		SetErroneousResponse(w, err)
		log.Printf(err.Error())
	}
}

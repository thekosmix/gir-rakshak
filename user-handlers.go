package main

import (
	"encoding/json"
	"net/http"
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
	// var user LoginUserRequest
	// body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	// if err != nil {
	// 	panic(err)
	// }
	// if err := r.Body.Close(); err != nil {
	// 	panic(err)
	// }
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// if err := json.Unmarshal(body, &user); err != nil {
	// 	w.WriteHeader(422) // unprocessable entity
	// 	if err := json.NewEncoder(w).Encode(err); err != nil {
	// 		panic(err)
	// 	}
	// }

	// t := RepoRegisterUser(user)
	// response := RegisterUserResponse{0, "", t}
	// w.WriteHeader(http.StatusCreated)
	// if err := json.NewEncoder(w).Encode(response); err != nil {
	// 	panic(err)
	// }
}

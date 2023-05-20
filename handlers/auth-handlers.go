package handlers

import (
	"encoding/json"
	"gir-rakshak/repo"
	"gir-rakshak/utils"
	"net/http"
	"strconv"
)

var unAuthRoutes = [...]string{
	"Index", "RegisterUser", "LoginUser"}

func AuthHandler(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.SetResponseHeaders(w)
		for _, a := range unAuthRoutes {
			if a == name {
				inner.ServeHTTP(w, r)
				return
			}
		}

		at := r.Header.Get("at")
		uid := r.Header.Get("uid")
		uidInt, _ := strconv.Atoi(uid)
		redisToken := repo.GetUserToken(uidInt)

		if at != "" && uid != "" && at == redisToken {
			inner.ServeHTTP(w, r)
			return
		}

		json.NewEncoder(w).Encode(utils.BaseResponse{Code: 1, Text: "401 Unauthorized"})
	})
}

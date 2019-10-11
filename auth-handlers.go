package main

import (
	"log"
	"net/http"
	"strconv"
)

var unAuthRoutes = [...]string{
	"Index", "RegisterUser", "LoginUser"}

func AuthHandler(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("in auth handler")
		for _, a := range unAuthRoutes {
			if a == name {
				inner.ServeHTTP(w, r)
				return
			}
		}

		at := r.Header.Get("at")
		uid := r.Header.Get("uid")
		uidInt, _ := strconv.Atoi(uid)
		redisToken := GetUserToken(uidInt)

		if at != "" && uid != "" && at == redisToken {
			inner.ServeHTTP(w, r)

			return
		}
		log.Printf("after auth handler")
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
	})
}

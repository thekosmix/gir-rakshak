package main

import (
	"gir-rakshak/handlers"
	"gir-rakshak/repo"
	"log"
	"net/http"
)

func main() {

	repo.InitConfig()

	router := handlers.NewRouter()

	log.Fatal(http.ListenAndServe(":80", router))
}

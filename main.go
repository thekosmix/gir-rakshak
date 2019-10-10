package main

import (
	"log"
	"net/http"
)

func main() {

	InitConfig()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":80", router))
}

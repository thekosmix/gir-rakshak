package main

import (
	"log"
	"net/http"
)

func main() {

	InitDB()
	
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":80", router))
}

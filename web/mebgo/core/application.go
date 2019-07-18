package core

import (
	"log"
	"net/http"
)

func Run() {
	router := LoadRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

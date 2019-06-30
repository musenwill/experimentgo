package sample

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// StartFileServer ...
func StartFileServer() {
	router := httprouter.New()
	router.ServeFiles("/*filepath", http.Dir("/home/musenwill"))
	log.Fatal(http.ListenAndServe(":8080", router))
}

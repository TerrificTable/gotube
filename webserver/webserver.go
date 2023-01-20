package webserver

import (
	"main/webserver/handlers"

	"github.com/gorilla/mux"
)

func NewWebServer(port int) {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.IndexHandler)

}

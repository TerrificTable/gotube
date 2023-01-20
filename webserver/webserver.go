package webserver

import (
	"fmt"
	"main/webserver/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type WebServer struct {
	Router *mux.Router
}

func NewWebServer() *WebServer {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.IndexHandler)
	r.HandleFunc("/search", handlers.SearchHandler)

	return &WebServer{
		Router: r,
	}
}

func (ws *WebServer) Start(port int, handler http.Handler) error {
	http.Handle("/", ws.Router)

	fmt.Printf("Starting webserver: http://localhost:%d\n", port)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}

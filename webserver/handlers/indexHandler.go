package handlers

import (
	"fmt"
	"net/http"
)

func IndexHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Hello World!")
}

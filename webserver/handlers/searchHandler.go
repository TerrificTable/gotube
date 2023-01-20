package handlers

import (
	"net/http"
	"main/data"
	"strconv"
	"fmt"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		term := r.FormValue("s")
		maxResults, err := strconv.Atoi(r.FormValue("max"))
		if err != nil {
			fmt.Fprintf(w, "Failed to convert '%s' to int", r.FormValue("max"))
			return
		}

		response, err := data.Service.Search(term, int64(maxResults))
		if err != nil {
			fmt.Fprintf(w, "Failed to search: %s", err.Error())
			return
		}

		data.YoutubeSearchData = *data.NewSearchData(*response)

		//fmt.Fprintf(w, "Success")

		http.Redirect(w, r, "/", 200)

	} else {
		http.NotFound(w, r)
	}
}

package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"main/data"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		apiKey := r.FormValue("api_key")

		fmt.Fprintf(w, "API key: %s\n", apiKey)

		// http.Redirect(w, r, "/", 301)
		return
	}

	t := template.Must(template.ParseGlob("webserver/static/*"))
	err := t.ExecuteTemplate(w, "index.gohtml", data.YoutubeSearchData)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

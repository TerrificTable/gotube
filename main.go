package main

import (
	"flag"
	"main/util"
	"main/webserver"
	"main/data"
)

const developerKey = ""

func main() {
	flag.Parse()

	data.Service = util.NewService(developerKey)

	ws := webserver.NewWebServer()
	if err := ws.Start(1234, nil); err != nil {
		panic(err)
	}

}

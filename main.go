package main

import (
	"flag"
	"main/util"
)

var (
	query      = flag.String("query", "Google", "Search term")
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

const developerKey = "AIzaSyCCA-ZUq_XZJT-FBoimui783TiST0Zw7Ns"

func main() {
	flag.Parse()

	service := util.NewService(developerKey)

	// Make the API call to YouTube.
	response, err := service.Search(*query, *maxResults)
	if err != nil {
		panic(err)
	}

	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)
	channels := make(map[string]string)
	playlists := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		case "youtube#channel":
			channels[item.Id.ChannelId] = item.Snippet.Title
		case "youtube#playlist":
			playlists[item.Id.PlaylistId] = item.Snippet.Title
		}
	}

	util.PrintSection("Videos", videos)
	util.PrintSection("Channels", channels)
	util.PrintSection("Playlists", playlists)
}

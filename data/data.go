package data

import (
	"google.golang.org/api/youtube/v3"
	"main/util"
)

var (
	YoutubeSearchData SearchData
	Service           util.Service
)

type SearchData struct {
	Videos    []youtube.SearchResult
	Playlists []youtube.SearchResult
	Channels  []youtube.SearchResult
}

func NewSearchData(data youtube.SearchListResponse) *SearchData {
	sd := SearchData{}

	for _, item := range data.Items {
		switch item.Id.Kind {
		case "youtube#video":
			sd.Videos = append(sd.Videos, *item)
			//fmt.Printf("%#+v\n", *item.Id)
		case "youtube#channel":
			sd.Channels = append(sd.Channels, *item)
		case "youtube#playlist":
			sd.Playlists = append(sd.Playlists, *item)
		}
	}

	return &sd
}

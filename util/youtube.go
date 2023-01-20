package util

import (
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

type Service struct {
	service *youtube.Service
}

func NewService(apiKey string) Service {
	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		panic(err)
	}

	return Service{
		service: service,
	}
}

func (s *Service) Search(search string, maxResults int64) (*youtube.SearchListResponse, error) {
	response, err := s.SearchParts(search, maxResults, []string{"id", "snippet"})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *Service) SearchParts(search string, maxResults int64, parts []string) (*youtube.SearchListResponse, error) {
	call := s.service.Search.List(parts).
		Q(search).
		MaxResults(maxResults)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	return response, nil
}

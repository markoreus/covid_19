package main

import (
	"net/http"
)

func getResponse() (*http.Response, error) {
	url := "https://corona-virus-stats.herokuapp.com/api/v1/cases/countries-search?limit=260"
	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

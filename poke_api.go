package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type areaData struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func getAreaData(url string) (areaData, error) {
	res, err := http.Get(url)
	if err != nil {
		return areaData{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return areaData{}, err
	}

	var data areaData
	if err := json.Unmarshal(body, &data); err != nil {
		return areaData{}, err
	}

	return data, nil
}

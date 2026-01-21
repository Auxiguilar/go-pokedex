package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

var startUrl = "https://pokeapi.co/api/v2/location-area/"

type Config struct {
	Client      http.Client
	UrlNext     *string
	UrlPrevious *string
}

type areaData struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func NewClient() Config {
	client := Config{
		Client:      http.Client{},
		UrlNext:     &startUrl,
		UrlPrevious: nil,
	}

	return client
}

func (cfg *Config) GetAreaData(url string) (areaData, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return areaData{}, err
	}

	res, err := cfg.Client.Do(req)
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

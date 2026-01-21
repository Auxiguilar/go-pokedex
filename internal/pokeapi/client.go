package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Auxiguilar/go-pokedex/internal/pokecache"
)

var startUrl = "https://pokeapi.co/api/v2/location-area/"

type Config struct {
	Client      http.Client
	Cache       pokecache.Cache
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

func NewConfig() Config {
	config := Config{
		Client:      http.Client{},
		Cache:       pokecache.NewCache(5 * time.Second),
		UrlNext:     &startUrl,
		UrlPrevious: nil,
	}

	return config
}

func (cfg *Config) GetAreaData(url string) (areaData, error) {
	// check if cached, return cached data
	entry, ok := cfg.Cache.Get(url)
	if ok {
		data, err := decodeAreaData(entry)
		if err != nil {
			return areaData{}, err
		}

		return data, err
	}

	// if not cached, get data
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

	// cache data
	cfg.Cache.Add(url, body)

	data, err := decodeAreaData(body)

	return data, nil
}

func decodeAreaData(body []byte) (areaData, error) {
	var data areaData
	if err := json.Unmarshal(body, &data); err != nil {
		return areaData{}, err
	}

	return data, nil
}

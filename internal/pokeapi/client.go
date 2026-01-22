package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Auxiguilar/go-pokedex/internal/pokecache"
)

var locationsUrl = "https://pokeapi.co/api/v2/location-area/"

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

// giant ugly struct! >:(
type locationData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func NewConfig() Config {
	config := Config{
		Client:      http.Client{},
		Cache:       pokecache.NewCache(5 * time.Second),
		UrlNext:     &locationsUrl,
		UrlPrevious: nil,
	}

	return config
}

func (cfg *Config) makeRequest(method string, url string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return []byte{}, err
	}

	res, err := cfg.Client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	// cache data
	cfg.Cache.Add(url, body)

	return body, nil
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
	resBody, err := cfg.makeRequest(http.MethodGet, url)
	if err != nil {
		return areaData{}, err
	}

	data, err := decodeAreaData(resBody)

	return data, nil
}

func decodeAreaData(body []byte) (areaData, error) {
	var data areaData
	if err := json.Unmarshal(body, &data); err != nil {
		return areaData{}, err
	}

	return data, nil
}

func (cfg *Config) GetLocationData(areaName string) (locationData, error) {
	fullUrl := locationsUrl + areaName

	// check if cached, return cached data
	entry, ok := cfg.Cache.Get(fullUrl)
	if ok {
		data, err := decodeLocationData(entry)
		if err != nil {
			return locationData{}, err
		}

		return data, err
	}

	// if not cached, get data
	resBody, err := cfg.makeRequest(http.MethodGet, fullUrl)
	if err != nil {
		return locationData{}, err
	}

	data, err := decodeLocationData(resBody)

	return data, nil
}

func decodeLocationData(body []byte) (locationData, error) {
	var data locationData
	if err := json.Unmarshal(body, &data); err != nil {
		return locationData{}, err
	}

	return data, nil
}

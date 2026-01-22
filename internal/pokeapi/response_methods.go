package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// it's getting a bit long-winded...

func (cfg *Config) makeRequest(method string, url string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("Creating request: %w", err)
	}

	res, err := cfg.Client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("Getting response: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("Non-OK status code: %v", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Reading response: %w", err)
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
			return areaData{}, fmt.Errorf("Decoding cached response: %w", err)
		}

		cfg.Cache.Add(url, entry)

		return data, nil
	}

	// if not cached, get data
	resBody, err := cfg.makeRequest(http.MethodGet, url)
	if err != nil {
		return areaData{}, fmt.Errorf("Making request: %w", err)
	}

	data, err := decodeAreaData(resBody)
	if err != nil {
		return areaData{}, fmt.Errorf("Decoding response: %w", err)
	}

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
			return locationData{}, fmt.Errorf("Decoding cached response: %w", err)
		}

		cfg.Cache.Add(fullUrl, entry)

		return data, err
	}

	// if not cached, get data
	resBody, err := cfg.makeRequest(http.MethodGet, fullUrl)
	if err != nil {
		return locationData{}, fmt.Errorf("Making request: %w", err)
	}

	data, err := decodeLocationData(resBody)
	if err != nil {
		return locationData{}, fmt.Errorf("Decoding response: %w", err)
	}

	return data, nil
}

func decodeLocationData(body []byte) (locationData, error) {
	var data locationData
	if err := json.Unmarshal(body, &data); err != nil {
		return locationData{}, err
	}

	return data, nil
}

func (cfg *Config) GetPokemonData(pokeName string) (pokemonData, error) {
	fullUrl := "https://pokeapi.co/api/v2/pokemon/" + pokeName

	// check if cached, return cached data
	entry, ok := cfg.Cache.Get(fullUrl)
	if ok {
		data, err := decodePokemonData(entry)
		if err != nil {
			return pokemonData{}, fmt.Errorf("Decoding cached response: %w", err)
		}

		cfg.Cache.Add(fullUrl, entry)

		return data, err
	}

	// if not cached, get data
	resBody, err := cfg.makeRequest(http.MethodGet, fullUrl)
	if err != nil {
		return pokemonData{}, fmt.Errorf("Making request: %w", err)
	}

	data, err := decodePokemonData(resBody)
	if err != nil {
		return pokemonData{}, fmt.Errorf("Decoding response: %w", err)
	}

	return data, nil
}

func decodePokemonData(body []byte) (pokemonData, error) {
	var data pokemonData
	if err := json.Unmarshal(body, &data); err != nil {
		return pokemonData{}, err
	}

	return data, nil
}

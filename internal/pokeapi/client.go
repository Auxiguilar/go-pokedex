package pokeapi

import (
	"net/http"
	"time"

	"github.com/Auxiguilar/go-pokedex/internal/pokecache"
)

var locationsUrl = "https://pokeapi.co/api/v2/location-area/"

type Config struct {
	Client      http.Client
	Cache       pokecache.Cache
	Pokemon     map[string]pokemonData
	UrlNext     *string
	UrlPrevious *string
}

func NewConfig() Config {
	config := Config{
		Client:      http.Client{},
		Cache:       pokecache.NewCache(30 * time.Second),
		Pokemon:     map[string]pokemonData{},
		UrlNext:     &locationsUrl,
		UrlPrevious: nil,
	}

	return config
}

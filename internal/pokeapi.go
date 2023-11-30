package pokeapi

import (
	"net/http"
	"time"
)

// responsible to connect to the PokeApi on the internet

// https://pokeapi.co/api/v2/location-area/?offset=20&limit=20
// JsonLint for better vizualisation
//  https://transform.tools/json-to-go/ Transform Json to Go structure, UnMarshal/parse it into struct

// type LocationAreasResp struct {
// 	Count    int     `json:"count"`
// 	Next     *string `json:"next"`
// 	Previous *string `json:"previous"`
// 	Results  []struct {
// 		Name string `json:"name"`
// 		URL  string `json:"url"`
// 	} `json:"results"`
// }

const baseURL string = "https://pokeapi.co/api/v2/"

// Client: Transform a specific mechanism by which individual http request are made
type Client struct {
	httpClient http.Client
}

// NewClient:  Stop the http request after a minute
func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute, 
		},
	}
}


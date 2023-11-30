package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullUrl := baseURL + endpoint

	// Check the cache , Key = url
	data, ok := c.cache.Get(fullUrl)
	if ok {
		// cache it
		fmt.Println("Cache hit!")
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon) // Transform & Transfer Json Data -> struct

		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil // early return if data already there
	}
	fmt.Println("chach miss !")

	req, err := http.NewRequest("GET", fullUrl, nil)

	// If error return O value of LocationAreasResp and error
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close() // close the Response Object when we finish

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon) // Transform & Transfer Json Data -> struct

	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullUrl, dat)

	return pokemon, nil
}

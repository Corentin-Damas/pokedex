package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseURL + endpoint
	if pageURL != nil {
		fullUrl = *pageURL
	}

	// Check the cache , Key = url
	data, ok := c.cache.Get(fullUrl)
	if ok {
		// cache it
		fmt.Println("Cache hit!")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreasResp) // Transform & Transfer Json Data -> struct

		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreasResp, nil // early return if data already there
	}
	fmt.Println("chach miss !")

	req, err := http.NewRequest("GET", fullUrl, nil)

	// If error return O value of LocationAreasResp and error
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close() // close the Response Object when we finish

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp) // Transform & Transfer Json Data -> struct

	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullUrl, dat)

	return locationAreasResp, nil
}


func (c *Client) GetLocationAreas(locationAreasName string) (LocationAreas, error) {
	endpoint := "/location-area/" + locationAreasName
	fullUrl := baseURL + endpoint

	// Check the cache , Key = url
	data, ok := c.cache.Get(fullUrl)
	if ok {
		// cache it
		fmt.Println("Cache hit!")
		locationAreas := LocationAreas{}
		err := json.Unmarshal(data, &locationAreas) // Transform & Transfer Json Data -> struct

		if err != nil {
			return LocationAreas{}, err
		}
		return locationAreas, nil // early return if data already there
	}
	fmt.Println("chach miss !")

	req, err := http.NewRequest("GET", fullUrl, nil)

	// If error return O value of LocationAreasResp and error
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close() // close the Response Object when we finish

	if resp.StatusCode > 399 {
		return LocationAreas{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	locationAreas := LocationAreas{}
	err = json.Unmarshal(dat, &locationAreas) // Transform & Transfer Json Data -> struct

	if err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(fullUrl, dat)

	return locationAreas, nil
}

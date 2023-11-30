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
	if pageURL != nil{
		fullUrl = *pageURL
	}


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
		return LocationAreasResp{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp) // Transform & Transfer Json Data -> struct 

	if err != nil{
		return LocationAreasResp{}, err
	}
	return  locationAreasResp, nil
}

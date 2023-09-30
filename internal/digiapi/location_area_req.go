package digiapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (c *Client) ListOptions(pageURL *string, page int) (digiAPI, error) {
	endpoint := "digimon?page=" + strconv.Itoa(page)
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return digiAPI{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return digiAPI{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return digiAPI{}, fmt.Errorf("Something went wrong, status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return digiAPI{}, err
	}

	listOptions := digiAPI{}
	err = json.Unmarshal(data, &listOptions)
	if err != nil {
		return digiAPI{}, err
	}
	return listOptions, nil
}

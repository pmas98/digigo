package digiapi

import (
	"net/http"
	"time"
)

const baseURL = "https://www.digi-api.com/api/v1/"

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

package digiapi

import (
	"net/http"
	"time"

	"github.com/pmas98/digigo/internal/digicache"
)

const baseURL = "https://www.digi-api.com/api/v1/"

type Client struct {
	cache      digicache.Cache
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		cache: digicache.NewCache(),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

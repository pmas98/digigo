package main

import (
	"github.com/pmas98/digigo/internal/digiapi"
)

type config struct {
	digiapiClient digiapi.Client
	nextListURL   *string
	prevListURL   *string
}

func main() {
	cfg := config{
		digiapiClient: digiapi.NewClient(),
	}

	start(&cfg)
}

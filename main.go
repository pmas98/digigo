package main

import (
	"github.com/pmas98/digigo/internal/digiapi"
)

type config struct {
	digiapiClient digiapi.Client
	nextListURL   *string
	prevListURL   *string
	caughtDigimon map[string]digiapi.DigimonStruct
}

func main() {
	cfg := config{
		digiapiClient: digiapi.NewClient(),
		caughtDigimon: make(map[string]digiapi.DigimonStruct),
	}

	start(&cfg)
}

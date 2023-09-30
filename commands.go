package main

import (
	"errors"
	"fmt"
	"os"
)

func callbackHelp(cfg *config) {
	fmt.Println("Welcome to the DigiGO menu:")
	fmt.Println("Here are some commands:")

	availableCommands := getCommands()
	for _, command := range availableCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
}

func callbackExit(cfg *config) {
	fmt.Println("Exiting program...")
	os.Exit(0)
}

func callbackList(cfg *config, page int) error {
	digiapiClient := cfg.digiapiClient

	resp, err := digiapiClient.ListOptions(cfg.nextListURL, page)

	if err != nil {
		panic(err)
	}
	for _, digimon := range resp.Content {
		fmt.Println("- " + digimon.Name)
	}
	cfg.nextListURL = &resp.Pageable.NextPage
	cfg.prevListURL = &resp.Pageable.PreviousPage
	return nil
}

func callbackListB(cfg *config, page int) error {
	if cfg.prevListURL == nil {
		return errors.New("No previous page")
	}
	digiapiClient := cfg.digiapiClient

	resp, err := digiapiClient.ListOptions(cfg.prevListURL, page)
	if err != nil {
		panic(err)
	}
	for _, digimon := range resp.Content {
		fmt.Println("- " + digimon.Name)
	}
	cfg.nextListURL = &resp.Pageable.NextPage
	cfg.prevListURL = &resp.Pageable.PreviousPage
	return nil
}

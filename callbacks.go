package main

import (
	"errors"
	"fmt"
	"math/rand"
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

func callbackDigimon(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Invalid number of arguments")
	}

	digimonName := args[0]

	digiapiClient := cfg.digiapiClient

	resp, err := digiapiClient.GetDigimon(digimonName)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n" + resp.Name)
	for i := 0; i < len(resp.Descriptions); i++ {
		if resp.Descriptions[i].Language == "en_us" {
			fmt.Print("\n" + resp.Descriptions[i].Description[0:140] + "\n")
		}
	}

	fmt.Print("\nType\n")
	fmt.Println("\n" + resp.Types[0].Type)

	fmt.Print("\nImage\n")
	fmt.Println(resp.Images[0].Href)

	fmt.Println("\nSkills")

	for i := 0; i < 5; i++ {
		fmt.Println(resp.Skills[i].Skill)
	}
	fmt.Println("\nPrior evolutions")

	for i := 0; i < len(resp.PriorEvolutions); i++ {
		fmt.Println(resp.PriorEvolutions[i].Digimon)
	}

	return nil
}

func callbackCatchDigimon(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Invalid number of arguments")
	}

	digimonName := args[0]

	if _, ok := cfg.caughtDigimon[digimonName]; ok {
		fmt.Printf("\nYou already caught %s\n", digimonName)
		return nil
	}

	digiapiClient := cfg.digiapiClient

	digimonStats := map[string]int{
		"Digitama": 1,
		"Baby I":   5,
		"Baby II":  10,
		"Child":    20,
		"Adult":    100,
		"Perfect":  500,
		"Ultimate": 10000,
		"Mega":     500000,
	}

	multiplier := map[string]float64{
		"Digitama": 10,
		"Baby I":   8,
		"Baby II":  6,
		"Child":    5,
		"Adult":    3,
		"Perfect":  2,
		"Ultimate": 1.1,
		"Mega":     1.05,
	}

	resp, err := digiapiClient.GetDigimon(digimonName)
	if err != nil {
		panic(err)
	}

	level := resp.Levels[0].Level

	randMult := multiplier[level] * rand.Float64()

	randNum := rand.Intn(int(float64(digimonStats[level]) * randMult))

	if randNum > digimonStats[level] {
		fmt.Printf("You caught %s\n", digimonName)
		cfg.caughtDigimon[digimonName] = resp
	} else {
		fmt.Printf("You failed to catch the digimon, your exp is %d and the digimon's exp is %d\n", randNum, digimonStats[level])
	}
	return nil
}

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Invalid number of arguments")
	}

	digimonName := args[0]

	if _, ok := cfg.caughtDigimon[digimonName]; !ok {
		fmt.Printf("\nYou haven't caught %s yet\n", digimonName)
	} else {
		callbackDigimon(cfg, args...)
	}

	return nil
}

func callbackCaughtDigimons(cfg *config, args ...string) error {
	if len(cfg.caughtDigimon) == 0 {
		return errors.New("You haven't caught any digimon yet")
	}
	fmt.Println("\nHere are the digimon you caught:")
	for digimon := range cfg.caughtDigimon {
		fmt.Println("- " + digimon)
	}
	return nil
}

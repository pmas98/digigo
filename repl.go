package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	function    func(*config, ...string)
}

var page int = 0

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints this help message",
			function:    func(cfg *config, args ...string) { callbackHelp(cfg) },
		},
		"exit()": {
			name:        "exit()",
			description: "Exits the program",
			function:    func(cfg *config, args ...string) { callbackExit(cfg) },
		},
		"digimon": {
			name:        "digimon {name}",
			description: "Gets a digimon's skills, image, stats and previous evolutions",
			function:    func(cfg *config, args ...string) { callbackDigimon(cfg, args...) },
		},
		"catch": {
			name:        "catch {name}",
			description: "Try to catch a digimon and put it on your team",
			function:    func(cfg *config, args ...string) { callbackCatchDigimon(cfg, args...) },
		},
		"inspect": {
			name:        "inspect {name}",
			description: "Inspect a digimon you caught!",
			function:    func(cfg *config, args ...string) { callbackInspect(cfg, args...) },
		},
		"digimons": {
			name:        "digimons",
			description: "Get a list of the digimons you caught!",
			function:    func(cfg *config, args ...string) { callbackCaughtDigimons(cfg) },
		},
		"list": {
			name:        "list",
			description: "Get a list of digimon according to a page number",
			function:    func(cfg *config, args ...string) { callbackList(cfg, page) },
		},
		"listb": {
			name:        "listb",
			description: "Get the previous list of digimon according to a page number",
			function:    func(cfg *config, args ...string) { callbackListB(cfg, 0) },
		},
	}
}
func start(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nPlease enter some text > ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]
		args := []string{}

		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		commands, ok := availableCommands[command]

		if !ok {
			fmt.Printf("I don't know what the command \"%s\" is\n", command)
			continue
		}
		if command != "list" {
			commands.function(cfg, args...)
		} else {
			fmt.Print("Please enter the page > ")
			scanner.Scan()
			text := scanner.Text()
			page, _ = strconv.Atoi(text)
			fmt.Print("\n")
			commands.function(cfg, args...)
		}
	}
}
func cleanInput(str string) []string {
	loweredString := strings.ToLower(str)
	words := strings.Fields(loweredString)
	return words
}

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
	function    func(*config)
}

var page int = 0

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints this help message",
			function:    callbackHelp,
		},
		"exit()": {
			name:        "exit()",
			description: "Exits the program",
			function:    callbackExit,
		},
		"list": {
			name:        "list",
			description: "Get a list of digimon according to a page number",
			function:    func(cfg *config) { callbackList(cfg, page) }, //Gambiarra
		},
		"listb": {
			name:        "listb",
			description: "Get the previous list of digimon according to a page number",
			function:    func(cfg *config) { callbackListB(cfg, 0) }, //Gambiarra
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

		availableCommands := getCommands()

		commands, ok := availableCommands[command]

		if !ok {
			fmt.Printf("I don't know what the command \"%s\" is\n", command)
			continue
		}
		if command != "list" {
			commands.function(cfg)
		} else {
			fmt.Print("Please enter the page > ")
			scanner.Scan()
			text := scanner.Text()
			page, _ = strconv.Atoi(text)
			fmt.Print("\n")
			commands.function(cfg)
		}
	}
}
func cleanInput(str string) []string {
	loweredString := strings.ToLower(str)
	words := strings.Fields(loweredString)
	return words
}

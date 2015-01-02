// Copyright 2015 Joseph Spurrier
// Author: Joseph Spurrier (http://josephspurrier.com)
// License: http://www.apache.org/licenses/LICENSE-2.0.html

/*
Package goconsole provides simple command line functionality.

Console output after typing a few commands:

	*** Go Console ***
	> help
	Available commands:
	clear - Clear the screen
	exit - Exit the console
	hello - Prints: world
	help - Show a list of available commands

	> hello
	world

	>

*/
package goconsole

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
)

// Map of all the callable commands
var commands = make(map[string]command)

// Controls the console loop
var active = true

// Displayed at the beginning of a line. Defaults to no linebreak. Possible value is "\n".
var Prompt = ""

// Displayed when a command is not found.
var NotFound = "Not found:"

// Displayed when first opened on the top line.
var Title = ""

// Displayed between each line
var NewLine = ""

// Structure of a command
type command struct {
	name        string
	description string
	function    func(string)
}

// Load the default commands at startup
func init() {
	loadDefaultCommands()
}

// *****************************************************************************
// Core
// *****************************************************************************

// The Start function starts the console loop where the user is prompted for keywords and then runs the registered commands.
func Start() {
	ClearScreen()
	fmt.Print(Title)

	// Set the initial values
	var typed string
	active = true

	// Loop while the value is true
	for active {
		// Prompt the user for a keyword
		fmt.Print(Prompt)
		typed = Readline()

		// If at least a character is typed
		if arr := strings.Fields(typed); len(arr) > 0 {

			// If the keyword is found
			if cmd, ok := commands[arr[0]]; ok {
				// Call the function
				cmd.function(typed)
				fmt.Println()
				// If the keyword is not found
			} else {
				// Output the NotFound message
				fmt.Println(NotFound + arr[0])
			}

			fmt.Print(NewLine)
		}
	}
}

// Load the default commands
func loadDefaultCommands() {
	Add("clear", "Clear the screen", func(typed string) {
		ClearScreen()
	})
	Add("exit", "Exit the console", func(typed string) {
		Leave()
	})
	Add("help", "Show a list of available commands", func(typed string) {
		// Sort by keywords
		keys := make([]string, 0)
		for key := range commands {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		// Output the commands
		fmt.Println("Available commands:")
		for i, val := range keys {
			if i == len(keys)-1 {
				fmt.Print(commands[val].name, " - ", commands[val].description)
			} else {
				fmt.Println(commands[val].name, "-", commands[val].description)
			}
		}
	})
}

// *****************************************************************************
// Console Configuration
// *****************************************************************************

// The Add function registers a new console keyword, description (used in help), and function.
func Add(keyword string, description string, function func(string)) {
	commands[keyword] = command{keyword, description, function}
}

// The Remove function unregisters a console keyword so it cannot be called.
func Remove(keyword string) {
	delete(commands, keyword)
}

// The Clear function unregisters all the console keywords so they cannot be called.
func Clear() {
	commands = make(map[string]command)
}

// *****************************************************************************
// Helpers
// *****************************************************************************

// The Readline function waits for the user to type and then press Enter. Readline returns the typed string.
func Readline() string {
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	return string(line)
}

// The ClearScreen function clears the screen. It is platform independent.
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// The Leave function ends the console loop
func Leave() {
	active = false
}

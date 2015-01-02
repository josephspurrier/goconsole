// Copyright 2015 Joseph Spurrier
// Author: Joseph Spurrier (http://josephspurrier.com)
// License: http://www.apache.org/licenses/LICENSE-2.0.html

package goconsole_test

import (
	"fmt"
	"github.com/josephspurrier/goconsole"
)

func Example() {
	goconsole.Title = "*** Go Console ***\n"
	goconsole.Prompt = "> "
	goconsole.NotFound = "Command not found: "
	goconsole.NewLine = "\n"
	goconsole.Add("hello", "Prints: world", func(typed string) {
		fmt.Print("world")
	})
	goconsole.Start()
}

func ExampleAdd() {
	goconsole.Add("hello", "Prints: world", func(typed string) {
		fmt.Print("world")
	})
}

func ExampleClear() {
	goconsole.Clear()
}

func ExampleClearScreen() {
	goconsole.ClearScreen()
}

func ExampleLeave() {
	goconsole.Leave()
}

func ExampleReadline() {
	fmt.Print("What is your favorite color? ")
	typed = Readline()
}

func ExampleRemove() {
	goconsole.Remove("help")
}

func ExampleStart() {
	goconsole.Start()
}

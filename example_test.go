// Copyright 2015 Joseph Spurrier
// Author: Joseph Spurrier (http://josephspurrier.com)
// License: http://www.apache.org/licenses/LICENSE-2.0.html

package goconsole_test

import (
	"fmt"
	"github.com/josephspurrier/goconsole"
)

func Example() {
	con := goconsole.New()
	con.Title = "*** Go Console ***\n"
	con.Prompt = "> "
	con.NotFound = "Command not found: "
	con.NewLine = "\n"
	con.Add("hello", "Prints: world", func(typed string) {
		fmt.Print("world")
	})
	con.Start()
}

func ExampleAdd() {
	con := goconsole.New()
	con.Add("hello", "Prints: world", func(typed string) {
		fmt.Print("world")
	})
}

func ExampleClear() {
	con := goconsole.New()
	con.Clear()
}

func ExampleClearScreen() {
	goconsole.ClearScreen()
}

func ExampleReadline() {
	fmt.Print("What is your favorite color? ")
	typed := goconsole.Readline()
	fmt.Print(typed)
}

func ExampleRemove() {
	con := goconsole.New()
	con.Remove("help")
}

func ExampleStart() {
	con := goconsole.New()
	con.Start()
}

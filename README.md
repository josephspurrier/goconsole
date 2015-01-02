goregistry
==========
[![GoDoc](https://godoc.org/github.com/josephspurrier/goconsole?status.svg)](https://godoc.org/github.com/josephspurrier/goconsole)

Golang Command Line Functionality

Package provides simple command line functionality with the ability to easily add commands and customize the interface. Complete documentation available on [GoDoc](https://godoc.org/github.com/josephspurrier/goconsole).

## Code

```
package main

import (
	"fmt"
	"github.com/josephspurrier/goconsole"
)

func main() {
	goconsole.Title = "*** Go Console ***\n"
	goconsole.Prompt = "> "
	goconsole.NotFound = "Command not found: "
	goconsole.NewLine = "\n"
	goconsole.Add("hello", "Prints: world", func(typed string) {
		fmt.Print("world")
	})
	goconsole.Start()
}
```

## Interface
```
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
```
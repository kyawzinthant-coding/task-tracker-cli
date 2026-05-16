package main

import (
	"fmt"
	"os"
	"task-cli/cli"
)

func main() {

	command, err := cli.ParseCommand(os.Args)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	err = command.Validate()
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	switch command.Name {
	case "add":
		fmt.Println("Add")
	default:
		fmt.Println("Unknown command")
	}
}

package main

import (
	"fmt"
	"os"
	"task-cli/cli"
	"task-cli/storage"
	task "task-cli/tasks"
)

func main() {

	store := storage.NewStorage("task.json")
	taskService := task.NewService(store)

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
		task, err := taskService.AddTask(command.Args[0])
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		fmt.Println("AddTask", task.ID)
	case "list":
		tasks, err := taskService.ListTasks()
		if err != nil {
			fmt.Println("Error", err)
		}
		cli.PrintTasks(tasks)
	default:
		fmt.Println("Unknown command")
	}
}

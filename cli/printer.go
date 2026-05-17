package cli

import (
	"fmt"
	task "task-cli/tasks"
)

func PrintTasks(tasks []task.Task) {
	for _, t := range tasks {
		fmt.Printf("ID: %d | %s | %s\n", t.ID, t.Status, t.Description)
	}
}

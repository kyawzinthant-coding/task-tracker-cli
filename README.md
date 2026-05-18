# Task Tracker CLI

A simple command-line app to manage your tasks. Built with Go — no external dependencies.

This is my solution to the [Task Tracker](https://roadmap.sh/projects/task-tracker) project on roadmap.sh.

## Features

- Add, update, and delete tasks
- Mark tasks as in-progress or done
- List all tasks or filter by status
- Tasks are saved to a local `task.json` file automatically

## Installation

Make sure you have [Go](https://golang.org/dl/) installed (1.18+).

```bash
git clone https://github.com/kyawzinthant-coding/task-cli.git
cd task-cli
go build -o task-cli .
```

## Usage

```bash
# Add a new task
./task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Update a task
./task-cli update 1 "Buy groceries and cook dinner"

# Delete a task
./task-cli delete 1

# Mark a task as in progress
./task-cli mark-in-progress 1

# Mark a task as done
./task-cli mark-done 1

# List all tasks
./task-cli list

# List tasks by status
./task-cli list todo
./task-cli list in-progress
./task-cli list done
```

## Task Properties

Each task stored in `task.json` has the following fields:

| Field | Description |
|-------|-------------|
| `id` | Unique identifier |
| `description` | Task description |
| `status` | `todo`, `in-progress`, or `done` |
| `createdAt` | When the task was created |
| `updatedAt` | When the task was last updated |

## Project Structure

```
task-cli/
├── main.go           # Entry point, wires everything together
├── cli/
│   ├── parser.go     # Parses command-line arguments
│   ├── command.go    # Validates commands
│   └── printer.go    # Prints tasks to terminal
├── tasks/
│   ├── task.go       # Task model and status constants
│   └── service.go    # Business logic (add, delete, update, list)
└── storage/
    └── storage.go    # Reads and writes task.json
```

## Running Tests

```bash
go test ./...
```

## What I Learned

- Dependency inversion with interfaces to avoid import cycles
- Test-driven development (TDD) with mock implementations
- Error handling patterns in Go
- Working with the filesystem using only the standard library

package tasks

import "time"

type Status string

const (
	Todo       Status = "Todo"
	InProgress Status = "in-progress"
	Done       Status = "Done"
)

type Task struct {
	ID          int
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

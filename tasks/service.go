package tasks

import "time"

type Service struct {
	tasks  []Task
	nextID int
}

func NewService() *Service {
	return &Service{
		tasks:  []Task{},
		nextID: 1,
	}
}

func (s *Service) AddTask(description string) Task {
	task := Task{
		ID:          s.nextID,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.tasks = append(s.tasks, task)
	s.nextID++
	return task
}

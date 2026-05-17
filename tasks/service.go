package tasks

import (
	"time"
)

type Store interface {
	Load() ([]Task, error)
	Save([]Task) error
}

type Service struct {
	tasks   []Task
	nextID  int
	storage Store
}

func NewService(store Store) *Service {
	tasks, err := store.Load()
	if err != nil {
		panic(err)
	}

	nextID := 1
	for _, t := range tasks {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}

	return &Service{
		tasks:   tasks,
		nextID:  nextID,
		storage: store,
	}
}

func (s *Service) AddTask(description string) (Task, error) {
	task := Task{
		ID:          s.nextID,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.tasks = append(s.tasks, task)
	s.nextID++
	err := s.storage.Save(s.tasks)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *Service) ListTasks() ([]Task, error) {
	return s.tasks, nil
}

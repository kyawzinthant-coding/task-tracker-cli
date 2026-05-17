package storage

import (
	"encoding/json"
	"os"

	task "task-cli/tasks"
)

type Storage struct {
	filePath string
}

func NewStorage(filePath string) *Storage {
	return &Storage{
		filePath: filePath,
	}
}

func (s *Storage) Load() ([]task.Task, error) {
	file, err := os.Open(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []task.Task{}, nil
		}
		return nil, err
	}

	defer file.Close()

	var loadedTasks []task.Task
	err = json.NewDecoder(file).Decode(&loadedTasks)
	if err != nil {
		return nil, err
	}

	return loadedTasks, nil
}

func (s *Storage) Save(tasks []task.Task) error {
	file, err := os.Create(s.filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}

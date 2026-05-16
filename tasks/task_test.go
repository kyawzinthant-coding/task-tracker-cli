package tasks

import "testing"

func TestAddTask(t *testing.T) {
	s := NewService()

	task := s.AddTask("buy milk")

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}

	if task.Description != "buy milk" {
		t.Errorf("expected description 'buy milk', got %s", task.Description)
	}

	if task.Status != Todo {
		t.Errorf("expected status todo, got %s", task.Status)
	}

	if len(s.tasks) != 1 {
		t.Errorf("expected 1 task in service, got %d", len(s.tasks))
	}
}

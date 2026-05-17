package tasks

import "testing"

type mockStorage struct {
	saved []Task
}

func (m *mockStorage) Load() ([]Task, error) {
	return m.saved, nil
}

func (m *mockStorage) Save(task []Task) error {
	m.saved = task
	return nil
}

func TestAddTask(t *testing.T) {
	store := &mockStorage{}
	s := NewService(store)

	task, err := s.AddTask("buy milk")
	if err != nil {
		t.Errorf("Error adding task: %s", err)
	}

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}

	if task.Description != "buy milk" {
		t.Errorf("expected description 'buy milk', got %s", task.Description)
	}

	if task.Status != Todo {
		t.Errorf("expected status todo, got %s", task.Status)
	}

	if (len(store.saved) != 1) || (store.saved[0].Description != "buy milk") {
		t.Errorf("expected saved tasks to contain 1 task, got %d", len(store.saved))
	}
}

func TestListTask(t *testing.T) {
	store := &mockStorage{
		saved: []Task{
			{
				ID:          1,
				Description: "buy milk",
				Status:      "Todo",
			},
			{
				ID:          2,
				Description: "buy book",
				Status:      "Done",
			},
		},
	}
	s := NewService(store)

	list, err := s.ListTasks()
	if err != nil {
		t.Errorf("Error listing tasks: %s", err)
	}

	if len(list) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(list))
	}

}

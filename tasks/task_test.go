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
			{
				ID:          3,
				Description: "buy CD",
				Status:      "in-progress",
			},
		},
	}
	s := NewService(store)

	t.Run("Done Task", func(t *testing.T) {
		tasks, err := s.ListTasks("Done")
		if err != nil {
			t.Errorf("Error listing tasks: %s", err)
		}

		for _, task := range tasks {
			if task.Status != Done {
				t.Errorf("expected status done, got %s", task.Status)
			}
		}
	})

	t.Run("in-progress Task", func(t *testing.T) {
		tasks, err := s.ListTasks("in-progress")
		if err != nil {
			t.Errorf("Error listing tasks: %s", err)
		}

		for _, task := range tasks {
			if task.Status != InProgress {
				t.Errorf("expected status done, got %s", task.Status)
			}
		}
	})

	t.Run("all", func(t *testing.T) {
		tasks, err := s.ListTasks("")
		if err != nil {
			t.Errorf("Error listing tasks: %s", err)
		}

		if len(tasks) != 3 {
			t.Errorf("expected 2 tasks, got %d", len(tasks))
		}

	})
}

func TestRemoveTask(t *testing.T) {
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

	err := s.DeleteTask(1)
	if err != nil {
		t.Errorf("Error removing task: %s", err)
	}
	list, err := s.ListTasks()
	if err != nil {
		t.Errorf("Error listing tasks: %s", err)
	}
	if len(list) != 1 {
		t.Errorf("expected 1 task, got %d", len(list))
	}

	if list[0].ID != 2 {
		t.Errorf("expected ID 2, got %d", list[0].ID)
	}
}

func TestStatusUpdate(t *testing.T) {
	t.Run("Done Task", func(t *testing.T) {
		store := &mockStorage{
			saved: []Task{
				{
					ID:          1,
					Description: "buy milk",
					Status:      Todo,
				},
			},
		}

		s := NewService(store)

		err := s.UpdateStatusDone(1)

		if err != nil {
			t.Fatalf("error updating status: %v", err)
		}

		if store.saved[0].Status != Done {
			t.Errorf("expected status %s, got %s", Done, store.saved[0].Status)
		}
	})

	t.Run("InProgress Task", func(t *testing.T) {
		store := &mockStorage{
			saved: []Task{
				{
					ID:          1,
					Description: "buy milk",
					Status:      Todo,
				},
			},
		}

		s := NewService(store)

		err := s.UpdateStatusInProgress(1)

		if err != nil {
			t.Fatalf("error updating status: %v", err)
		}

		if store.saved[0].Status != InProgress {
			t.Errorf("expected status %s, got %s", InProgress, store.saved[0].Status)
		}
	})
}

func TestUpdateTask(t *testing.T) {

	store := &mockStorage{
		saved: []Task{
			{
				ID:          1,
				Description: "buy milk",
				Status:      Todo,
			},
		},
	}

	s := NewService(store)

	err := s.UpdateTask(1, "update PR")
	if err != nil {
		t.Fatalf("error updating status: %v", err)
	}

	if store.saved[0].Description != "update PR" {
		t.Errorf("expected description 'update PR', got %s", store.saved[0].Description)
	}

}

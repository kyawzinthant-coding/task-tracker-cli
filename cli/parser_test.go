package cli

import "testing"

func TestParseCommand(t *testing.T) {

	tests := []struct {
		name      string
		args      []string
		wantName  string
		wantArgs  []string
		wantError bool
	}{
		{
			name:      "add",
			args:      []string{"task-cli", "add", "book"},
			wantName:  "add",
			wantArgs:  []string{"book"},
			wantError: false,
		},
		{
			name:      "delete",
			args:      []string{"task-cli", "delete", "1"},
			wantName:  "delete",
			wantArgs:  []string{"1"},
			wantError: false,
		},
		{
			name:      "update",
			args:      []string{"task-cli", "update", "1", "new text"},
			wantName:  "update",
			wantArgs:  []string{"1", "new text"},
			wantError: false,
		},
		{
			name:      "list",
			args:      []string{"task-cli", "list"},
			wantName:  "list",
			wantArgs:  nil,
			wantError: false,
		},
		{
			name:      "mark-in-progress",
			args:      []string{"task-cli", "mark-in-progress", "1"},
			wantName:  "mark-in-progress",
			wantArgs:  []string{"1"},
			wantError: false,
		},
		{
			name:      "mark-done",
			args:      []string{"task-cli", "mark-done", "1"},
			wantName:  "mark-done",
			wantArgs:  []string{"1"},
			wantError: false,
		},
		{
			name:      "no command",
			args:      []string{"task-cli"},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			cmd, err := ParseCommand(tt.args)

			if tt.wantError {
				if err == nil {
					t.Error("want error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("want no error, got %v", err)
			}
			assertEqual(t, cmd.Name, tt.wantName)

			for i := range tt.wantArgs {
				assertEqual(t, cmd.Args[i], tt.wantArgs[i])
			}
		})
	}

}

func assertEqual(t *testing.T, got, want string) {
	if want != got {
		t.Errorf("got %s want %s", got, want)
	}
}

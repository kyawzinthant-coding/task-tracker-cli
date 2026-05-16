package cli

import "testing"

func TestParseCommand(t *testing.T) {

	t.Run("add", func(t *testing.T) {
		addArgs := []string{
			"task-cli",
			"add",
			"book",
		}
		command := ParseCommand(addArgs)

		assertEqual(t, command.Name, addArgs[1])
		assertEqual(t, command.Args[0], addArgs[2])
	})

}

func assertEqual(t *testing.T, got, want string) {
	if want != got {
		t.Errorf("got %s want %s", got, want)
	}
}

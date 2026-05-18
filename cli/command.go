package cli

import "errors"

func (c Command) Validate() error {
	switch c.Name {

	case "add":
		if len(c.Args) < 1 {
			return errors.New("missing task description")
		}

	case "delete":
		if len(c.Args) < 1 {
			return errors.New("missing task id")
		}

	case "update":
		if len(c.Args) < 2 {
			return errors.New("update requires: id + new text")
		}

	case "list":
		// no args needed
	case "mark-in-progress":
		if len(c.Args) < 1 {
			return errors.New("missing task id")
		}
	case "mark-done":
		if len(c.Args) < 1 {
			return errors.New("missing task id")
		}
	default:
		return errors.New("unknown command: " + c.Name)
	}

	return nil
}

package cli

import "errors"

type Command struct {
	Name string
	Args []string
}

func ParseCommand(args []string) (Command, error) {

	if len(args) < 2 {
		return Command{}, errors.New("not enough arguments")
	}

	return Command{
		Name: args[1],
		Args: args[2:],
	}, nil
}

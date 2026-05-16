package cli

type Command struct {
	Name string
	Args []string
}

func ParseCommand(args []string) Command {
	return Command{
		Name: args[1],
		Args: args[2:],
	}
}

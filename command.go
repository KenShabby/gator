package main

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (cmds *commands) run(s *state, cmd command) error {
	return nil
}

func (cmds *commands) register(name string, f func(*state, command) error) {
}

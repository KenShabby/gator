package main

import (
	"errors"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (cmds *commands) register(name string, f func(*state, command) error) {
	cmds.commandMap[name] = f
}

func (cmds *commands) run(s *state, cmd command) error {
	f, ok := cmds.commandMap[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}


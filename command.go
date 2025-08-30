package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (cmds *commands) run(s *state, cmd command) error {
	value, exists := cmds.commandMap[cmd.name]
	if exists {
		err := value(s, cmd)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("command does not exist")
	}
	return nil
}

func (cmds *commands) register(name string, f func(*state, command) error) {
	cmds.commandMap[name] = f
}

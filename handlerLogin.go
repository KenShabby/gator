package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no arguments given")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("could not set username: %v", err)
	}

	fmt.Printf("user has been set to: %s\n", cmd.args[0])
	return nil

}

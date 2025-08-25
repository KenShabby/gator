package main

import (
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		log.Println("no arguments given")
	}

	s.cfg.SetUser(cmd.name)

	fmt.Printf("User has been set to: %s", cmd.name)

	return nil

}

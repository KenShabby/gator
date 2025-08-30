package main

import (
	"gator/internal/config"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	st := state{&cfg}

	cliCommands := commands{
		commandMap: make(map[string]func(*state, command) error),
	}
	cliCommands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("too few command line arguments")
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	err = cliCommands.run(&st, cmd)
	if err != nil {
		log.Fatalf("could not run command %v", err)
	}
}

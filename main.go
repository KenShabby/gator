package main

import (
	"fmt"
	"log"

	"gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("Brendan")
	if err != nil {
		log.Fatal("could not set user: %v\n", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal("Error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)

}

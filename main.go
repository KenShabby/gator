package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {

	// Read the config file
	var cfg config.Config
	cfg, err := cfg.Read()
	if err != nil {
		return
	}
	fmt.Println(cfg)
	// Set the current user to "brendan" and update the config file on disk.
	// Read the config file again and print the contents of the config struct
	// to the terminal.
	// Adding another comment to test merging branches.
}

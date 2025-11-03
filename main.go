package main

import (
	"database/sql"
	"gator/internal/config"
	"gator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}


func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("could not open db: %s", cfg.DBURL)
	}
	dbQueries := database.New(db)
	defer db.Close()

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cliCommands := commands{
		commandMap: make(map[string]func(*state, command) error),
	}
	cliCommands.register("login", handlerLogin)
	cliCommands.register("register", handlerRegister)
	cliCommands.register("reset", handlerReset)
	cliCommands.register("users", handlerGetUsers)
	cliCommands.register("agg", handlerAgg)

	if len(os.Args) < 2 {
		log.Fatalf("too few command line arguments")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cliCommands.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatalf("could not run command %v", err)
	}
}

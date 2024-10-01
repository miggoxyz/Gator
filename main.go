package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/miggoxyz/gator/internal/config"
	"github.com/miggoxyz/gator/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal("error opening database")
	}
	
	dbQueries := database.New(db)

	currentState := &state{
		cfg: &cfg,
		db: dbQueries,
	}

	cmds := commands {
		availableCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerDeleteAll)
	cmds.register("users", handlerGetUsers)

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := []string{}

	if len(os.Args) > 2 {
		cmdArgs = os.Args[2:]
	}

	err = cmds.run(currentState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)  // Exit with code 1 when there's an error running the command
	}
}
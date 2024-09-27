package main

import (
	"fmt"
	"log"
	"os"

	"github.com/miggoxyz/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	currentState := &state{
		cfg: &cfg,
	}

	cmds := commands {
		availableCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

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
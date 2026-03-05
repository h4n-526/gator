package main

import (
	"gator/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("error reading config: ", err)
	}
	s := &state{cfg: &cfg}
	cmds := commands{cmdNames: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	err = cmds.run(s, command{name: os.Args[1], args: os.Args[2:]})
	if err != nil {
		log.Fatal(err)
	}
}

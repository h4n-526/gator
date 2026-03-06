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
		log.Fatal("error reading config: ", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal("error opening database: ", err)
	}
	s := &state{
		db:  database.New(db),
		cfg: &cfg,
	}
	cmds := commands{cmdNames: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerListFeeds)
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	err = cmds.run(s, command{name: os.Args[1], args: os.Args[2:]})
	if err != nil {
		log.Fatal(err)
	}
}

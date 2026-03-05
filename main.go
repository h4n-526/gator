package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		return
	}
	err = cfg.SetUser("test")
	if err != nil {
		return
	}
	fmt.Println(cfg.DBURL)
	fmt.Println(cfg.CurrentUserName)
}

package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("handler expects a single argument, the username")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("User has been set to", cmd.args[0])
	return nil
}

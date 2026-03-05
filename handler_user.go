package main

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("handler login expects a single argument, the username")
	}
	name := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("user %q does not exist, you can't login to an account that doesn't exist", name)
	}

	if err = s.cfg.SetUser(name); err != nil {
		return err
	}
	fmt.Println("User has been set to", name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("handler register expects a single argument, the username")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	if err = s.cfg.SetUser(user.Name); err != nil {
		return err
	}
	fmt.Println("User created and set to", user.Name)
	fmt.Printf("ID: %v\nCreated: %v\n", user.ID, user.CreatedAt)
	return nil
}

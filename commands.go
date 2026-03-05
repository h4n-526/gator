package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	cmdNames map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.cmdNames[cmd.name]
	if !ok {
		return errors.New("unknown command: " + cmd.name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, handler func(*state, command) error) {
	c.cmdNames[name] = handler
}

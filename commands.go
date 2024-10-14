package main

import (
	"errors"
	"fmt"
)

type Command struct {
	name string
	args []string
}

type Commands struct {
	commands map[string]func(*State, Command) error
}

func (c *Commands)Register(name string, f func(*State, Command) error) {
	c.commands[name] = f
}

func (c *Commands)Run(s *State, cmd Command) error {
	f, exists := c.commands[cmd.name] 
	if !exists {
		return fmt.Errorf("the command %v does not exist", cmd.name)
	}
	err := f(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func NewCommands() *Commands {
	return &Commands{commands: make(map[string]func(*State, Command) error)}
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}

	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("Username has been set")
	return nil
}


package main

import (
	"fmt"
	"os"

	"github.com/lazyturtlez/gator/internal/config"
)

type State struct {
	config *config.Config
}

func main(){
	currentConfig, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := &State{config: currentConfig}
	commands := NewCommands()
	commands.Register("login", HandlerLogin)
	args := os.Args
	if len(args) < 3 {
		fmt.Println("not enough args provided")
		os.Exit(1)
	}
	command := Command{name: args[1], args: args[2:]}
	err = commands.Run(s, command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
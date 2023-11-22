package main

import (
	"errors"
	"log"
	"os/exec"
)

type node struct {
	parent      *node
	name        string
	description string
	children    map[rune]node
	command     string
}

func (nd node) run_command() error {
	if nd.command == "" {
		return errors.New("command is empty")
	}
	cmd := exec.Command(nd.command)
	err := cmd.Run()
	if err != nil {
		log.Fatal("sddf", err)
	}
	return nil
}

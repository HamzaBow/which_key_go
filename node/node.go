package node

import (
	"errors"
	"log"
	"os/exec"
)

type Node struct {
	parent      *Node
	Name        string
	Description string
	Children    map[rune]Node
	command     string
}

func (nd Node) run_command() error {
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

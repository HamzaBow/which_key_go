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
	Command     string
}

func (nd Node) RunCommand() error {
	if nd.Command == "" {
		return errors.New("command is empty")
	}
	cmd := exec.Command("bash", "-c", nd.Command)
	err := cmd.Run()
	if err != nil {
		log.Fatal("sddf", err)
	}
	return nil
}

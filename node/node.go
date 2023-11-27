package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"which_key_go/util"
)

type Node struct {
	parent      *Node
	Name        string
	Description string
	Children    map[string]Node
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

func (nd Node) PrintSubTree() error {
	util.ClearTerminal()
	b, err := json.MarshalIndent(nd, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

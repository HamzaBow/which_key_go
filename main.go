package main

import (
	"encoding/json"
	"fmt"
	"os"
	"which_key_go/node"
	"which_key_go/prompt"
	"which_key_go/serialize"
)

func main() {
	rt := startUp()
	serialize.SerializeTreeToGob(*rt)
	prompt.PromptPrefixNode(*rt, []*node.Node{})
}

func startUp() *node.Node {

	var rt node.Node
	treeBytes, err := os.ReadFile("tree.json")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Creating default tree")
		dt := getDefaultTree()
		return &dt
	}
	err = json.Unmarshal(treeBytes, &rt)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Creating default tree")
		dt := getDefaultTree()
		return &dt
	}
	fmt.Println("Using serialized tree")
	return &rt
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"which_key_go/node"
	"which_key_go/prompt"
)

func main() {
	rt := startUp()
	prompt.PromptPrefixNode(*rt, []*node.Node{})
}

// func serializeTree(rt node.Node) {

// 	f, err := os.Create("tree.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	// bts, err := json.Marshal(rt)
// 	bts, err := json.MarshalIndent(rt, "", "\t")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = f.Write(bts)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("done")
// }

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

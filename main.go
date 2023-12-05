package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"which_key_go/node"
)

func main() {
	// rt := getDefaultTree()
	// serializeTree(rt)
	var rt node.Node
	treeBytes, err := os.ReadFile("tree.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(treeBytes, &rt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rt)
	rt.PromptPrefixNode()
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

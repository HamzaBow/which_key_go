package main

import (
	"which_key_go/node"
	"which_key_go/prompt"
)

func main() {
	rt := getDefaultTree()
	prompt.PromptPrefixNode(rt, []*node.Node{})
	// serializeTree(rt)
	// var rt node.Node
	// treeBytes, err := os.ReadFile("tree.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = json.Unmarshal(treeBytes, &rt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(rt)
	// rt.PromptPrefixNode()
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

package main

import (
	"fmt"
	"which_key_go/node"
)

func main() {
	var rt = node.Node{
		Name:        "Root",
		Description: "Desc of Root",
		Children:    map[rune]node.Node{},
	}
	fmt.Println(rt)
}

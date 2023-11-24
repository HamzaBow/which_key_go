package main

import (
	"encoding/json"
	"fmt"
	"which_key_go/node"
)

func main() {
	var rt = node.Node{
		Name:        "Root",
		Description: "Desc of Root",
		Children:    map[rune]node.Node{},
	}
	rt.Children['a'] = node.Node{
		Name:        "Child A",
		Description: "Desc for Child A",
		Children:    map[rune]node.Node{},
	}
	rt.Children['b'] = node.Node{
		Name:        "Child B",
		Description: "Desc for Child B",
		Children:    map[rune]node.Node{},
	}
	// fmt.Println(rt)
	b, _ := json.MarshalIndent(rt, "", "\t")
	fmt.Println(string(b))
	var v interface{}
	err := json.Unmarshal(b, v)
	fmt.Println("err:", err)
	fmt.Println(v)

}

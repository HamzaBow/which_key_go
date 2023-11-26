package main

import "which_key_go/node"

func getDefaultTree() node.Node {

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
	return rt
}

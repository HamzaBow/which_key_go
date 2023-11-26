package main

import "which_key_go/node"

func getDefaultTree() node.Node {

	var rt = node.Node{
		Name:        "Root",
		Description: "Desc of Root",
		Children:    map[rune]node.Node{},
	}
	rt.Children['a'] = node.Node{
		Name:     "Chrome Profile 1",
		Children: map[rune]node.Node{},
		Command:  "google-chrome --new-window --profile-directory=\"Profile 1\"",
	}
	rt.Children['b'] = node.Node{
		Name:     "Chrome Profile 2",
		Children: map[rune]node.Node{},
		Command:  "google-chrome --new-window --profile-directory=\"Profile 2\"",
	}
	return rt
}

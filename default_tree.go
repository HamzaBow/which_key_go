package main

import "which_key_go/node"

func getDefaultTree() node.Node {

	var rt = node.Node{
		Name:        "Root",
		Description: "Desc of Root",
		Children:    map[string]node.Node{},
	}
	rt.AddChild(node.Node{
		Name:     "Chrome Profile 1",
		Children: map[string]node.Node{},
		Command:  "google-chrome --new-window --profile-directory=\"Profile 1\"",
	}, "a")
	rt.AddChild(node.Node{
		Name:     "Chrome Profile 2",
		Children: map[string]node.Node{},
		Command:  "google-chrome --new-window --profile-directory=\"Profile 2\"",
	}, "b")
	return rt
}

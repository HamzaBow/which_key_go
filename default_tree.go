package main

import "which_key_go/node"

func getDefaultTree() node.Node {

	var rt = node.Node{
		Name:        "Root",
		Description: "Desc of Root",
		Children:    map[string]*node.Node{},
	}
	var browserNode = node.Node{
		Name:        "Browser",
		Description: "Desc of Browser",
		Children:    map[string]*node.Node{},
	}
	chromeNode := node.Node{
		Name:        "Chrome",
		Description: "Desc of Chrome",
		Children:    map[string]*node.Node{},
	}
	chromeNodeProfile1 := node.Node{
		Name:     "Chrome Profile 1",
		Children: map[string]*node.Node{},
		Command:  "google-chrome --new-window --profile-directory=\"Profile 1\"",
	}
	chromeNodeProfile2 := node.Node{
		Name:     "Chrome Profile 2",
		Children: map[string]*node.Node{},
		Command:  "google-chrome --new-window --profile-directory=\"Profile 2\"",
	}

	rt.Children["b"] = &browserNode
	browserNode.Children["c"] = &chromeNode
	chromeNode.Children["a"] = &chromeNodeProfile1
	chromeNode.Children["b"] = &chromeNodeProfile2

	return rt
}

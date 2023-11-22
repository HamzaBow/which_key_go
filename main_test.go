package main

import (
	"testing"
	"which_key_go/node"
)

func TestMapDefault(t *testing.T) {

	var rt1 = node.Node{
		Name:        "Root 1",
		Description: "Desc of Root 1",
		Children:    map[rune]node.Node{},
	}
	var rt2 = node.Node{
		Name:        "Root 2",
		Description: "Desc of Root 2",
	}

	var x = rt1.Children == nil
	t.Log("rt1.Children == nil", x)

	var y = rt2.Children == nil
	t.Log("rt2.Children == nil", y)

	if rt1.Children == nil {
		t.Fatal("Children of rt1 is nil, but shouldn't be ")
	}
	if rt2.Children != nil {
		t.Fatal("Children of rt2 isn't nil, but should be")
	}
}

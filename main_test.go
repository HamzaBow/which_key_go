package main

import "testing"

func TestMapDefault(t *testing.T) {

	var rt1 = node{
		name:        "Root 1",
		description: "Desc of Root 1",
		children:    map[rune]node{},
	}
	var rt2 = node{
		name:        "Root 2",
		description: "Desc of Root 2",
	}

	var x = rt1.children == nil
	t.Log("rt1.children == nil", x)

	var y = rt2.children == nil
	t.Log("rt2.children == nil", y)

	if rt1.children == nil {
		t.Fatal("children of rt1 is nil, but shouldn't be ")
	}
	if rt2.children != nil {
		t.Fatal("children of rt2 isn't nil, but should be")
	}
}

package main

type node struct {
	parent      *node
	name        string
	description string
	children    map[rune]node
	command     string
}

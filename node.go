package main

type node struct {
	name        string
	description string
	children    map[rune]node
	command     string
}

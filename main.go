package main

import "fmt"

func main() {
	var rt = node{
		name:        "Root",
		description: "Desc of Root",
		children:    map[rune]node{},
	}
	fmt.Println(rt)
}

package main

import "fmt"

func main() {
	var rt = getDefaultTree()
	// rt.PromptPrefixNode()
	var err = rt.PrintSubTree()
	fmt.Println(err)
}

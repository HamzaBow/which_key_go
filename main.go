package main

import "fmt"

func main() {
	var rt = getDefaultTree()
	rt.PrintSubTree()
	fmt.Println("-------------------------")
	rt.Children['a'].PrintSubTree()
}

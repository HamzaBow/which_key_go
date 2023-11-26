package main

import (
	"fmt"
	"time"
)

func main() {
	var rt = getDefaultTree()
	rt.PrintSubTree()
	time.Sleep(1 * time.Second)
	fmt.Println("-------------------------")
	rt.Children['a'].PrintSubTree()
}

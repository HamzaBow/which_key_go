package main

import "fmt"

func main() {
	child_nd1 := node{
		name: "Child A",
	}
	child_nd2 := node{
		name: "Child B",
	}
	nd := node{
		name: "Root",
		children: map[rune]node{
			'a': child_nd1,
			'b': child_nd2,
		},
	}
	fmt.Println(nd.children['b'])
	fmt.Println(nd.children['b'].children)
	fmt.Println(nd)
}

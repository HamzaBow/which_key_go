package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	nd := node{
		name:        "ND",
		description: "Desc ND",
		command:     "google-chrome",
	}
	fmt.Println(nd)
	cmd := exec.Command(nd.command)
	err := cmd.Run()
	if err != nil {
		log.Fatal("sddf", err)
	}
}

//
//
//
//
//
//
//
//
//
//

// func main() {
// 	child_nd1 := node{
// 		name: "Child A",
// 	}
// 	child_nd2 := node{
// 		name: "Child B",
// 	}
// 	nd := node{
// 		name: "Root",
// 		children: map[rune]node{
// 			'a': child_nd1,
// 			'b': child_nd2,
// 		},
// 	}
// 	fmt.Println(nd.children['b'])
// 	fmt.Println(nd.children['b'].children)
// 	fmt.Println(nd)
// }

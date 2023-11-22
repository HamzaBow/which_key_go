package main

import (
	"fmt"
	"log"
)

func main() {
	// code1()
	code2()
	// code3()
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
func code3() {
	var x = map[rune]string{
		'a': "abcd",
	}
	x['b'] = "bcde"
	fmt.Println(x)
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

func code2() {
	var rt = node{
		name:        "Root",
		description: "Desc of Root",
		children:    map[rune]node{},
	}
	var ch1 = node{
		parent:  &rt,
		command: "google-chrome",
	}
	rt.children['a'] = ch1
	fmt.Println(rt.children)
	// rt.children['a'].run_command()
	err := rt.run_command()
	if err != nil {
		log.Fatal(err)
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

func code1() {

	// child := node{
	// 	name:        "Child A",
	// 	description: "Desc Ch A",
	// }
	nd := node{
		name:        "ND",
		description: "Desc ND",
		command:     "google-chrome",
	}
	fmt.Println(nd.parent)
	// cmd := exec.Command(nd.command)
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal("sddf", err)
	// }
}

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

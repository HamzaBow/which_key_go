package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var rt = getDefaultTree()
	b, _ := json.MarshalIndent(rt, "", "\t")
	fmt.Println(string(b))
	var v interface{}
	err := json.Unmarshal(b, v)
	fmt.Println("err:", err)
	fmt.Println(v)
	rt.Children['b'].RunCommand()
}

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var rt = getDefaultTree()
	b, _ := json.MarshalIndent(rt, "", "\t")
	fmt.Println(string(b))
}

package node

import (
	"fmt"
	"which_key_go/util"
)

var notification = ""

func (nd Node) PromptPrefixNode() {
	util.ClearTerminal()
	fmt.Println(notification)
	fmt.Println("----------------------")

	if nd.Command != "" {
		fmt.Println("Executing command")
		nd.RunCommand()
		return
	}
	for k, v := range nd.Children {
		fmt.Println(k, " -> ", v.Name)
	}

	fmt.Print("Enter key: ")

	var char = string(util.ReadRune())

	// if err != nil {
	// 	fmt.Println("Error reading input:", err)
	// 	return
	// }

	if len(char) != 1 {
		notification = "key should be of length 1"
		nd.PromptPrefixNode()
		return
	}

	if char == "q" {
		fmt.Println("Quitting...")
		return
	}

	if char == "K" {
		if nd.parent == nil {
			notification = "Already at root node!"
			nd.PromptPrefixNode()
			return
		}
		nd := *nd.parent
		nd.PromptPrefixNode()
		// nd.PromptPrefixNode()
		return
	}

	newNode, exists := nd.Children[char]

	if !exists {
		notification = fmt.Sprint("Key \"", char, "\" does't exist")
		nd.PromptPrefixNode()
		return
	}

	newNode.PromptPrefixNode()

}

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

	var userInput string
	_, err := fmt.Scanln(&userInput)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if len(userInput) != 1 {
		notification = "key should be of length 1"
		nd.PromptPrefixNode()
		return
	}

	if userInput == "q" {
		fmt.Println("Quitting...")
		return
	}

	if userInput == "K" {
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

	newNode, exists := nd.Children[userInput]

	if !exists {
		notification = fmt.Sprint("Key \"", userInput, "\" does't exist")
		nd.PromptPrefixNode()
		return
	}

	newNode.PromptPrefixNode()

}

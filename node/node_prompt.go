package node

import (
	"fmt"
	"time"
)

func (nd Node) PromptPrefixNode() {

	if nd.Command != "" {
		fmt.Println("this is a command, we will execute it in a second:")
		time.Sleep(time.Second)
		nd.RunCommand()
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
		fmt.Println("key should be of length 1")
		nd.PromptPrefixNode()
		return
	}

	newNode, exists := nd.Children[userInput]

	if !exists {
		fmt.Println("Key \"", userInput, "\" does't exist")
		nd.PromptPrefixNode()
		return
	}

	newNode.PromptPrefixNode()

}

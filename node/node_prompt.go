package node

import (
	"fmt"
	"which_key_go/util"

	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#FF0000")).
	PaddingLeft(1).
	PaddingRight(1)
	// Background(lipgloss.Color("#7D56F4")).
	// PaddingTop(2).
	// Width(22)

var notification = ""
var prevNotif = ""

func (nd Node) PromptPrefixNode() {
	util.ClearTerminal()
	fmt.Println("path")
	fmt.Println("----------------------")

	if nd.Command != "" {
		fmt.Println("Executing command")
		nd.RunCommand()
		return
	}
	for k, v := range nd.Children {
		fmt.Println("", k, "➞", v.Name)
	}
	var diff = 10 - len(nd.Children)
	lineBreaks := ""
	for i := 0; i < diff; i++ {
		lineBreaks += "\n"
	}
	fmt.Print(lineBreaks)

	fmt.Println("----------------------")

	if notification != "" {
		fmt.Println(style.Render(notification))
	} else {

		fmt.Println()
	}

	if (prevNotif != notification) && notification != "" {
		prevNotif = notification
		notification = ""
	}
	fmt.Print("Enter key: ")

	var char = string(util.ReadRune())

	// if err != nil {
	// 	fmt.Println("Error reading input:", err)
	// 	return
	// }

	if len(char) != 1 {
		prevNotif = notification
		notification = "key should be of length 1"
		nd.PromptPrefixNode()
		return
	}

	if char == "q" {
		fmt.Println("\nQuitting...")
		return
	}

	if char == "K" {
		if nd.parent == nil {
			prevNotif = notification
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
		prevNotif = notification
		notification = fmt.Sprint("Key \"", char, "\" does't exist")
		nd.PromptPrefixNode()
		return
	}

	newNode.PromptPrefixNode()

}

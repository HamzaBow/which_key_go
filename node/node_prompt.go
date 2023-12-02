package node

import (
	"fmt"
	"which_key_go/util"

	"github.com/charmbracelet/lipgloss"
)

var keyStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFF00"))

var prefixStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#0000FF"))

var errorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#FF0000")).
	PaddingLeft(1).
	PaddingRight(1)

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
	printChildrenPairs(nd)
	var diff = 10 - len(nd.Children)
	lineBreaks := ""
	for i := 0; i < diff; i++ {
		lineBreaks += "\n"
	}
	fmt.Print(lineBreaks)

	fmt.Println("----------------------")

	if notification != "" {
		fmt.Println(errorStyle.Render(notification))
	} else {

		fmt.Println()
	}

	if (prevNotif != notification) && notification != "" {
		prevNotif = notification
		notification = ""
	}
	fmt.Print("Enter key: ")

	var char = string(util.ReadRune())

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

func printChildrenPairs(nd Node) {

	// nd.Display()

	for k, v := range nd.Children {
		fmt.Println("", keyStyle.Render(k), "➞", prefixStyle.Render(v.Name))
	}
}

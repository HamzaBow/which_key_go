package node

import (
	"fmt"
	"which_key_go/util"

	"github.com/charmbracelet/lipgloss"
)

type PathElement struct {
	Key  string
	Name string
}

var path = []PathElement{}

var KeyStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00FF00"))

var PrefixStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#2680d9"))

var GrayStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#777777"))

var commandStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#d9d9268"))

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
	// fmt.Println(path)
	printPath(path)
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
		path = path[:len(path)-1]
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

	path = append(path, PathElement{Key: char, Name: newNode.Name})
	newNode.PromptPrefixNode()

}

func printChildrenPairs(nd Node) {

	// nd.Display()

	for k, v := range nd.Children {
		var formattedName string
		if v.Command == "" {
			formattedName = PrefixStyle.Render(v.Name)
		} else {
			formattedName = commandStyle.Render(v.Name)

		}
		fmt.Println("", KeyStyle.Render(k), "➞", formattedName)
	}
}

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

var pathInfo = []PathElement{}

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

// TODO: might be better to use *[]Node instead of []Node to
// prevent copying data unnecessarily
func (nd Node) PromptPrefixNode(pathStack []*Node) {
	util.ClearTerminal()
	// fmt.Println(path)
	printPath(pathInfo)
	util.PrintDivider()

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

	util.PrintDivider()

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
		nd.PromptPrefixNode(pathStack)
		return
	}

	if char == "q" {
		fmt.Println("\nQuitting...")
		return
	}

	if char == "K" {
		if len(pathStack) == 0 {
			prevNotif = notification
			notification = "Already at root node!"
			nd.PromptPrefixNode(pathStack)
			return
		}
		pathInfo = pathInfo[:len(pathInfo)-1]
		parent := pathStack[len(pathStack)-1]
		pathStack = pathStack[:len(pathStack)-1]
		parent.PromptPrefixNode(pathStack)
		return
	}

	newNode, exists := nd.Children[char]

	if !exists {
		prevNotif = notification
		notification = fmt.Sprint("Key \"", char, "\" does't exist")
		nd.PromptPrefixNode(pathStack)
		return
	}

	pathInfo = append(pathInfo, PathElement{Key: char, Name: newNode.Name})
	pathStack = append(pathStack, &nd)
	newNode.PromptPrefixNode(pathStack)

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

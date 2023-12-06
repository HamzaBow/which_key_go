package node

import (
	"fmt"
	"which_key_go/styles"
	"which_key_go/util"
)

type PathElement struct {
	Key  string
	Name string
}

var pathInfo = []PathElement{}

var notification = ""
var prevNotif = ""

func (nd Node) PromptPrefixNode(pathStack []*Node) {
	promptPrepareAndPrintHeader()

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
		fmt.Println(styles.ErrorStyle.Render(notification))
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
	for k, v := range nd.Children {
		var formattedName string
		if v.Command == "" {
			formattedName = styles.PrefixStyle.Render(v.Name)
		} else {
			formattedName = styles.CommandStyle.Render(v.Name)
		}
		fmt.Println("", styles.KeyStyle.Render(k), "➞", formattedName)
	}
}

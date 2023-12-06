package node

import (
	"fmt"
	"which_key_go/styles"
)

func printPath(path []PathElement) {
	if len(path) == 0 {
		fmt.Println(styles.PrefixStyle.Render(" Root"))
		return
	}
	var names = ""
	var keychain = " "
	for i, p := range path {
		keychain += styles.KeyStyle.Render(p.Key)
		names = names + styles.PrefixStyle.Render(p.Name)
		if i < len(path)-1 {
			names += styles.GrayStyle.Render(" / ")
		}
	}
	fmt.Println(keychain, "|", names)
}

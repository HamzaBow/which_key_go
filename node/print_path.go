package node

import "fmt"

func printPath(path []PathElement) {
	var names = ""
	var keychain = ""
	for i, p := range path {
		keychain += KeyStyle.Render(p.Key)
		names = names + PrefixStyle.Render(p.Name)
		if i < len(path)-1 {
			names += GrayStyle.Render(" / ")
		}
	}
	fmt.Println(keychain, "|", names)
}

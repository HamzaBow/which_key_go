package node

import "which_key_go/util"

func promptPrepareAndPrintHeader() {
	util.ClearTerminal()
	printPath(pathInfo)
	util.PrintDivider()

}

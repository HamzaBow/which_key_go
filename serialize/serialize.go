package serialize

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"which_key_go/node"
)

var filenameWithoutExtension = "tree"

func serializeTree(tree node.Node) {

	f, err := os.Create(filenameWithoutExtension + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// bts, err := json.Marshal(rt)
	bts, err := json.MarshalIndent(tree, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(bts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")
}

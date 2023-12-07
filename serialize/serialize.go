package serialize

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"which_key_go/node"
)

var filenameWithoutExtension = "tree"

func SerializeTree(tree node.Node) {

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

func SerializeTreeToGob(tree node.Node) error {
	file, err := os.Create(filenameWithoutExtension + ".gob")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(tree)
	if err != nil {
		return err
	}

	return nil
}

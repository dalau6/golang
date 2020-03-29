package main

import (
	"fmt"

	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("./filehandling") // A box represents a folder whose contents will be embedded in the binary.
	data := box.String("test.txt")        // read the contents of the file.
	fmt.Println("Contents of file:", data)
}

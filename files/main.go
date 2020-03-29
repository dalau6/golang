package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	// "github.com/gobuffalo/packr"
)

func main() {
	// box := packr.NewBox("./filehandling") // A box represents a folder whose contents will be embedded in the binary.
	// data := box.String("test.txt")        // read the contents of the file.
	// fmt.Println("Contents of file:", data)

	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b[0:n]))
	}
}

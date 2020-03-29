package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	/* This function accepts 3 arguments. The first is the name of the flag, second is the default value and the third is a short description of the flag. */
	/* When this program is run using the command `-fpath=/path-of-file/test.txt`, we pass `/path-of-file/test.txt` as the value of the flag fpath. */
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)

	data, err := ioutil.ReadFile("./filehandling/test.txt") // ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

package main

import (
	"fmt"
	"os"
)

func writeStringBytes() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// writing string to a file
	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	// writing bytes to a file
	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	n2, err := f.Write(d2)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")
	fmt.Println(n2, "bytes written successfully")

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func writeLines() {
	f, err := os.Create("lines")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}

	for _, v := range d {
		// use the Fprintln function to write the lines to a file.
		// The Fprintln function takes a io.writer as parameter and appends a new line
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}

func appendFile() {
	f, err := os.OpenFile("lines", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}

func main() {
	writeStringBytes()
	writeLines()
	appendFile()
}

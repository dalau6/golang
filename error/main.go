package main

import (
	"fmt"
	"math"
	"net"
	"os"
	"path/filepath"
)

/*
	// Package errors implements functions to manipulate errors.
	package errors

	// New returns an error that formats as the given text.
	func New(text string) error {
		return &errorString{text}
	}

	// errorString is a trivial implementation of error.
	type errorString struct {
		s string
	}

	func (e *errorString) Error() string {
		return e.s
	}
*/

/*
	It contains a single method with signature Error() string.
	Any type which implements this interface can be used as an error.
	This method provides the description of the error.

	type error interface {
		Error() string
	}

	func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
*/
func open() bool {
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return true
	}
	fmt.Println(f.Name(), "opened successfully")
	return false
}

/*
	DNS lookups

	If you read the documentation of the Open function carefully,
	you can see that it returns an error of type *PathError.
	PathError is a struct type and its implementation in the standard library is as follows,

	type PathError struct {
		Op   string
		Path string
		Err  error
	}
*/
func lookup() {
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}

/*
	ErrBadPattern is defined in the filepath package as follows.
	var ErrBadPattern = errors.New("syntax error in pattern")
	This function returns an error ErrBadPattern when the pattern is malformed.
*/
func badpattern() {
	files, error := filepath.Glob("[")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	fmt.Println("matched files", files)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

func main() {
	open()
	lookup()
	badpattern()

	// custom errors
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

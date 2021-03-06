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
		// return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func calcCircle() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

type areaRecError struct {
	err    string  //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

func (e *areaRecError) Error() string {
	return e.err
}

func (e *areaRecError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaRecError) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaRecError{err, length, width}
	}
	return length * width, nil
}

func calcRec() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaRecError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)

			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)

			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}

func main() {
	open()
	lookup()
	badpattern()

	// custom errors
	calcCircle()
	calcRec()
}

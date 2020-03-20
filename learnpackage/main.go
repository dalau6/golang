package main

import (
	"fmt"
	"learnpackage/simpleinterest" //importing custom package
	"learnpackage/structs/computer"
	"log"
)

var p, r, t = 5000.0, 10.0, 1.0

/*
* init function to check if p, r and t are greater than zero
 */
func init() {
	println("Main package initialized")
	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}

func main() {
	fmt.Println("Simple interest calculation")
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)

	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
	structs()
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

func structs() {
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	fmt.Println("Spec:", spec)
}

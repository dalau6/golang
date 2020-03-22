package main

import (
	"fmt"
	"learnpackage/simpleinterest" //importing custom package
	"learnpackage/structs/computer"
	"log"
	"time"
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

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func digits(number int, dchnl chan<- int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func calcSquares(number int, squareop chan<- int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
	close(squareop)
}

func calcCubes(number int, cubeop chan<- int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
	close(cubeop)
}

func producer(chnl chan<- int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	fmt.Println("Simple interest calculation")
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)

	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
	structs()

	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)

	for {
		squares, ok := <-sqrch
		cubes, okay := <-cubech
		if ok == false && okay == false {
			break
		}
		fmt.Println("Received ", squares, cubes)
		// squares, cubes := <-sqrch, <-cubech
		fmt.Println("Final output", squares+cubes)
	}

	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received ", v)
	}

	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
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

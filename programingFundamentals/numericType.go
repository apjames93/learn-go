package main

import "fmt"

// integers have no decimals also known as whole numbers
var x int

// floating point numbers do have decimals known as real numbers
var y float64

func main() {
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	y = 99.66
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	var z int8 = -128 // cant be -129 or more than 128 bites
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}

// https://golang.org/ref/spec#Numeric_types

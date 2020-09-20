package main

import "fmt"

// closure: enclose the scope of a var and contain it to a certain area. so the scope of the var is limited
// you want to keep the scope of your vars as narrow as possable

var x int

// https://golang.org/ref/spec#Defer_statements
func main() {
	fmt.Println("the the scope of x is the entire package", x)
	foo()

	{
		y := 42
		fmt.Println(y)
	}

	fmt.Println("you cant run 'y' here or code would break because the scope is limited to the block above")

	// a and b have different var x when incrementing
	a := incrementor()
	fmt.Println("a=====", a())
	fmt.Println("a=====", a())
	fmt.Println("a=====", a())
	fmt.Println("a=====", a())

	b := incrementor()
	fmt.Println("b ======", b())
	fmt.Println("b ======", b())
	fmt.Println("b ======", b())
	fmt.Println("b ======", b())

}

func foo() {
	x++
	fmt.Println("adding to x without passing in as an arg", x)
}

func incrementor() func() int {
	var x int

	return func() int {
		// func dosent need to have x passed down because it is scoped to the incrementer block
		x++
		return x
	}
}

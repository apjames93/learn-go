package main

import "fmt"

// https://golang.org/ref/spec#Passing_arguments_to_..._parameters
func main() {
	x := sum(1, 2, 3, 4, 5) // have to invoke fuction to call it even if there are not any arguments
	fmt.Println("total is", x)

	y := sum() // can pass 0 values in and get the zero value of the argument type
	fmt.Println("total is", y)

	// first argument is for paramter string, anything after is variadic parameter for t ...int
	nameAndTimes("bradford", 5, 6, 7)
	nameAndTimes("alex")
}

// unlimited number of ints
func sum(x ...int) int {
	// x is a slice of values [1, 2, 3, 4, 5]
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for i, v := range x {
		sum += v
		fmt.Println("for item in index position, ", i, " we are now adding.", v, " to the total", sum)
	}
	return sum
}

// variadic parameters has to be the final paramter func sum(s string, x ...int) int { ... }
func nameAndTimes(n string, t ...int) {
	fmt.Println(n, " times :", t)
}

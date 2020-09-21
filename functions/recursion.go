package main

import "fmt"

// recursion: when a function calls itself a certain number of times then stops

// https://golang.org/ref/spec#Defer_statements
func main() {

	n := factorial(4)
	fmt.Println(n)

	n2 := factorial(4)
	fmt.Println(n2)
}

// the product of an integer and all the integers below it; e.g. factorial four ( 4! ) is equal to 24. (factorial of 4) = 4! = (4*3*2*1)
func factorial(n int) int {
	// will keep calling factorial until the arg passed in is 0
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
	// 4 * (3 * (2 * 1) ) || 4 * factorial(4-1) * factorial(3-1) * factorial(2-1) * 1
	// each of these function need to be put into memory and then executed to return the value
}

// same thing with a loop
func loopFact(n int) int {
	total := 1
	// leave init empty with a condition to loop over
	for ; n > 0; n-- {
		total *= n
	}
	return total
}

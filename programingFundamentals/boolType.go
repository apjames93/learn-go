package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) // zerovalue false
	x = true

	fmt.Println(x)

	a := 7
	b := 42
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a <= b) // true
	fmt.Println(a >= b) // false
}

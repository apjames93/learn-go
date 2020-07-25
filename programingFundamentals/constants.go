package main

import "fmt"

const a = 42 // untyped constant

// typed constant
const (
	b float64 = 42.78
	c         = "james bond"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

package main

import "fmt"

var a int

// create a type of hotdog that is of type int
type hotdog int

// make var b of type hotdog (that is type int)
var b hotdog

// main operator
func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// you can not assign "a" to "b" even though hotdog has an underlining type of int and "a" is of type int
	// a = b

	// conversion can take type hotdog and change it to type int becuse hotdogs underling type is int
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

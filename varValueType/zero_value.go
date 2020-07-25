package main

import "fmt"

// main ..
func main() {
	// DECLARE a VARIABLE to be of a certain TYPE and then ASSIGN a VALUE of that TYPE to the VARIABLE
	// use var when you wnt to delcare a ZERO VALUE for a VARIABLE

	var a string
	fmt.Println("zero value of stirng", a)
	fmt.Printf("%T\n", a)

	a = "not empty string"
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var b int
	fmt.Println("zero value of int", b)
	fmt.Printf("%T\n", b)

	b = 3030
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	var c bool
	fmt.Println("zero value of bool", c)
	fmt.Printf("%T\n", c)

	c = true
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	var d float64
	fmt.Println("zero value of float64", d)
	fmt.Printf("%T\n", d)

	d = 1.01
	fmt.Println(d)
	fmt.Printf("%T\n", d)

	var e float32
	fmt.Println("zero value of float32", e)
	fmt.Printf("%T\n", e)

	e = 1.01
	fmt.Println(e)
	fmt.Printf("%T\n", e)

}

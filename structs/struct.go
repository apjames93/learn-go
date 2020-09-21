package main

import "fmt"

// https://golang.org/ref/spec#Struct_types
type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "alex",
		last:  "james",
	}
	p2 := person{
		first: "brooke",
		last:  "higgins",
	}

	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.last)

	fmt.Println(p2)
	fmt.Println(p2.first)
	fmt.Println(p2.last)
}

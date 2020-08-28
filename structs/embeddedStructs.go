package main

import "fmt"

// https://golang.org/ref/spec#Struct_types
type person struct {
	first string
	last  string
}

type programer struct {
	person
	golan bool
	first string
}

func main() {
	p1 := programer{
		person: person{
			first: "alex",
			last:  "james",
		},
		golan: true,
	}

	p2 := programer{
		person: person{
			first: "brooke",
			last:  "higgins",
		},
		first: "buthead",
		golan: false,
	}

	fmt.Println(p1)
	// inner type gets promoted to to outer type unless name collision
	fmt.Println(p1.first, p1.person.first)

	fmt.Println(p2)
	// name collision inner type has to be accessed by dot notation
	fmt.Println(p2.first, p2.person.first)
}

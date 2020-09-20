package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// when you have a reciver it will attach this func to the type (secretAgent) so any value of they type has access to this method
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.person.last)

}

// https://golang.org/ref/spec#Defer_statements
func main() {
	as1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	// secretAgent type can use method speak with dot notation
	as1.speak()

	as2 := secretAgent{
		person: person{
			"Some",
			"Dude",
		},
		ltk: true,
	}
	as2.speak()
}

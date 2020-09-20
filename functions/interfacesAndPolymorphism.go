package main

import "fmt"

// interfaces allows us to define behaviour and to use polymorphism
// allows a value to be more than one type
// joke: a interface says "hey baby if you got this method then you are my type"

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}

type person struct {
	first string
	last  string
}

type clothed struct {
	person
	shirt bool
	pants bool
	shoes bool
}

type human interface {
	speak()
}

func canLeave(h human) {
	switch h.(type) {
	case person:
		fmt.Println(h.(person).first, "you have to be clothed to leave")
	case clothed:
		if h.(clothed).shirt == true && h.(clothed).pants == true && h.(clothed).shoes == true {
			fmt.Println("I am out of here")
		} else {
			fmt.Println("shirt ", h.(clothed).shirt, "pants ", h.(clothed).pants, "shoes ", h.(clothed).shoes)
		}
	}
}

// https://golang.org/ref/spec#Interface_types
func main() {
	// a value can be more than one type ex: p is type person and human
	p := person{
		first: "James",
		last:  "Bond",
	}

	p.speak()
	canLeave(p)

	p1 := clothed{
		person: person{
			first: "bradford",
			last:  "james",
		},
		shirt: false,
		pants: false,
		shoes: true,
	}

	p1.speak()
	canLeave(p1)

	p2 := clothed{
		person: person{
			first: "jo",
			last:  "james",
		},
		shirt: true,
		pants: true,
		shoes: true,
	}

	p2.speak()
	canLeave(p2)

}

// a variadic paramiter of a empty interface evey value can be passed in

// TODO: https://www.ardanlabs.com/blog/2015/09/composition-with-go.html

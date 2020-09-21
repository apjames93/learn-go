package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first:      "alex",
		last:       "james",
		favFlavors: []string{"chocolate", "vanilla"},
	}

	p2 := person{
		first:      "jo",
		last:       "james",
		favFlavors: []string{"chocolate", "cheesecake"},
	}

	fmt.Println(p1.first)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}

}

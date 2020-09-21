package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) sayHello() {
	fmt.Println("hello i am", p.first, p.last)
}

func main() {
	p := person{
		first: "guy",
		last:  "dude",
		age:   90,
	}
	p.sayHello()
}

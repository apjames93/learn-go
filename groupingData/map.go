package main

import "fmt"

// https://golang.org/doc/effective_go.html#maps
func main() {
	m := map[string]int{
		"alex":   27,
		"brooke": 21,
	}
	fmt.Println(m)

	fmt.Println(m["alex"])

	fmt.Println(m["bradford"]) // will retun zero value of type

	v, ok := m["jo"]
	fmt.Println(v)  // zero value
	fmt.Println(ok) // false doesn't exist in map

	if v, ok := m["brooke"]; ok {
		fmt.Println("if brooke", v)
	}

}

package main

import "fmt"

func main() {
	var person = struct {
		first   string
		last    string
		friends map[string]int
	}{
		first: "alex",
		last:  "james",
		friends: map[string]int{
			"bradford": 666,
		},
	}
	fmt.Println(person.first)
	fmt.Println(person.friends)

	for k, v := range person.friends {
		fmt.Println(k, v)
	}

}

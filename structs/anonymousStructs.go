package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "alex",
		last:  "james",
		age:   27,
	}

	fmt.Println(p1)

}

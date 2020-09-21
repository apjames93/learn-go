package main

import "fmt"

// https://golang.org/doc/effective_go.html#slices
func main() {
	m := map[string]int{
		"alex":   27,
		"brooke": 21,
	}
	fmt.Println(m)

	m["bradford"] = 9000

	for k, v := range m {
		println(k, v)
	}

}

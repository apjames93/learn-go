package main

import "fmt"

func main() {
	m := map[string][]string{
		`alex`:     []string{"cars", "code", "crepes"},
		`bradford`: []string{"buble", "binary", "baboons"},
		`jo`:       []string{"jam", "jars", "jac"},
	}
	// fmt.Println(m)

	for k, v := range m {
		fmt.Println("this is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

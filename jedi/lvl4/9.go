package main

import "fmt"

func main() {
	m := map[string][]string{
		`alex`:     []string{"cars", "code", "crepes"},
		`bradford`: []string{"buble", "binary", "baboons"},
		`jo`:       []string{"jam", "jars", "jac"},
	}

	m["brooke"] = []string{"bathing", "barking", "biting"}

	for k, v := range m {
		fmt.Println("this is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

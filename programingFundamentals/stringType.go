package main

import (
	"fmt"
)

func main() {
	s := "helllloooo"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	rawStringLiteral := `
	raw
		string
	literal
	`
	fmt.Println(rawStringLiteral)
	fmt.Printf("%T\n", rawStringLiteral)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i]) // utf8 code point
	}
	fmt.Println("")

	for i, v := range s {
		fmt.Println(i, v)
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}

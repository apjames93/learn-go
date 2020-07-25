package main

import (
	"fmt"
)

func main() {
	s := "rsx"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	for i := 0; i < len(bs); i++ {
		fmt.Printf("%T\n", bs[i])  // unit8
		fmt.Printf("%b\n", bs[i])  // binary
		fmt.Printf("%#X\n", bs[i]) // hex
		fmt.Printf("%#U", bs[i])   // utf8
	}
}

package main

import "fmt"

const (
	_  = iota             // 0
	kb = 1 << (iota * 10) // 1 * 10 shift 10
	mb = 1 << (iota * 10) // 2 * 10 shift 20
	gb = 1 << (iota * 10) // 3 * 10 shift 30
)

func main() {
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}

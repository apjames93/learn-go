package main

import "fmt"

// https://golang.org/ref/spec#Function_declarations
func main() {
	foo()
	bar("harrison")
	hello := woo("alex")
	fmt.Println(hello)
	x, y := mouse("Harrison", "ford")
	fmt.Println(x)
	fmt.Println(y)

}

// func (r receiver) identifier(parameters) (return(s)) { ... }

func foo() {
	fmt.Println("foo is called")
}

// eveything is Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("hello from woo,", s)
}

func mouse(fn string, ln string) (string, bool) {
	s := fmt.Sprint(fn, " ", ln, `, says "hello"`)
	b := false
	return s, b
}

// func (r receiver) identifier(parameters) (return(s)) { ... }

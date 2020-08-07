package main

// https://golang.org/doc/effective_go.html#allocation_make
func main() {
	aj := []string{"alex", "james", "honda", "rsx"}
	println(aj)
	bl := []string{"bradford", "lamson", "subaru", "sti"}
	println(bl)

	buddys := [][]string{aj, bl}
	println(buddys)
}

package main

// https://golang.org/doc/effective_go.html#maps
func main() {
	m := map[string]int{
		"alex":   27,
		"brooke": 21,
	}

	delete(m, "alex")
	println(m)

	if _, ok := m["alex"]; ok {
		delete(m, "alex")
	}
}

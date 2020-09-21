package main

func main() {
	x := []int{42, 43, 45, 46, 47, 48, 48, 49, 50, 51}
	println(x)
	x = append(x[:3], x[6:]...)
	println(x)
}

package main

import (
	"fmt"
	"runtime"
)

func main() {
	// runs what the operating system and architecture that the system is using
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

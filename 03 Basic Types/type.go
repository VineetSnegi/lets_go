package main

import (
	"fmt"
)

func main() {
	a := 2
	b := 3.1
	fmt.Printf("a: %T %[1]v\n", a)
	fmt.Printf("b: %T %[1]v\n", b)
}
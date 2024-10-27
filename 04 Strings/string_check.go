package main

import (
	"fmt"
)

func main () {
	s := "hello"
	t := "こんにちは"
	
	fmt.Printf("s: %8T %[1]v\n", s)
	fmt.Printf("s: %8T %[1]v\n", []rune(s))
	fmt.Printf("s: %8T %[1]v\n", []byte(s))
	fmt.Printf("\n")
	fmt.Printf("t: %8T %[1]v\n", t)
	fmt.Printf("t: %8T %[1]v\n", []rune(t))
	fmt.Printf("t: %8T %[1]v\n", []byte(t))
}
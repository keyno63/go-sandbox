package main

import "fmt"

func main() {
	fmt.Printf("Hello world!")

	var (
		a1 int = 123
		a2 int = 456
	)
	fmt.Printf(string(rune(a1)))
	fmt.Printf(string(rune(a2)))
}

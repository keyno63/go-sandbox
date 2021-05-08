package main

import "fmt"

func main() {
	fmt.Printf("Hello world!")

	var (
		a1 int = 123
		//a2 int = 456
	)
	fmt.Printf(string(rune(a1)))
	//fmt.Printf(string(rune(a2)))

	p1 := map[string]interface{}{
		"name": "Yamada",
		"age":  26,
	}
	name := p1["name"]
	fmt.Printf(name.(string))

	fmt.Printf(sampleFunc())
}

func sampleFunc() string {
	return "string"
}

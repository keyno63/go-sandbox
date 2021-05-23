package main

import "fmt"

func main() {
	fmt.Printf("Hello world!")

	fmt.Printf(sampleFunc())

	pointerFunc()
}

func sampleFunc() string {
	return "string"
}

func pointerFunc() {
	var a1 int
	var p1 *int

	p1 = &a1
	*p1 = 123
	fmt.Println(a1)
}

func defineVar() {
	var (
		a1 int = 123
		//a2 int = 456
	)
	fmt.Printf(string(rune(a1)))
	//fmt.Printf(string(rune(a2)))
}

func defineMap() {
	p1 := map[string]interface{}{
		"name": "Yamada",
		"age":  26,
	}
	name := p1["name"]
	fmt.Printf(name.(string))
}

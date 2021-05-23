package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello world!")

	fmt.Printf(sampleFunc())

	pointerFunc()

	sampleGoroutine()
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

func sampleGoroutine() {
	// channel の生成
	ch1 := make(chan int)
	f := func() {
		defer close(ch1)
		a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for _, v := range a {
			fmt.Printf("add channel: %d\n", v)
			// channel への書き込み
			ch1 <- v
		}
		//close(ch1)
	}
	go f()
	//time.Sleep(time.Second * 1)
	// channel から読み出し
	for c := range ch1 {
		fmt.Printf("read channel: %d\n", c)
	}
	fmt.Printf("go out\n")
	//value := <-ch1
	//fmt.Printf("%d", value)
}

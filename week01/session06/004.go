package main

import "fmt"

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	myAdder10 := Adder{start: 10}
	fmt.Println(myAdder10.AddTo(5)) // 15

	f1 := myAdder10.AddTo
	fmt.Println(f1(5)) // 15

	f2 := Adder.AddTo
	fmt.Println(f2(myAdder10, 5)) // 15
}

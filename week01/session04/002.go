package main

import (
	"fmt"
	"strconv"
)

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

type opFuncType func(int, int) int

func main() {
	var opMap = map[string]opFuncType{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		{"1", "+", "2"},
		{"1", "-", "2"},
		{"1", "*", "2"},
		{"1", "/", "2"},
		{"1", "%", "2"},
		{"1", "2", "3"},
		{"one", "+", "two"},
		{"1", "2"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression")
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("invalid operator")
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}

	for i := 0; i < 10; i++ {
		func(j int) {
			fmt.Println(j)
		}(i)
	}
}

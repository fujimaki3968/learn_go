package main

import (
	"errors"
	"fmt"
)

type divFuncOpts struct {
	debug bool
}

func div(a int, b int, opts divFuncOpts) (int, error) {
	if opts.debug {
		fmt.Println("debug: ", a, "/", b)
	}
	if b == 0 {
		return 0, errors.New("divied by zero")
	}
	return a / b, nil
}

func div2(a int, b int, opts divFuncOpts) (result int, err error) {
	if opts.debug {
		fmt.Println("debug: ", a, "/", b)
	}
	if b == 0 {
		return 0, errors.New("divied by zero") // 別にreturn 0, err と書いても良い
	}
	result = a / b
	return result, nil
}

func addAll(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func main() {
	fmt.Println(div(10, 2, divFuncOpts{})) // 5

	fmt.Println(div(10, 2, divFuncOpts{debug: true})) // 10 / 0

	fmt.Println(div2(10, 2, divFuncOpts{})) // 5

	result, err := div(10, 0, divFuncOpts{})
	if err != nil {
		fmt.Println(err) // divided by zero
	}
	fmt.Println(result) // 0

	fmt.Println(addAll(1, 2, 3, 4, 5)) // 15
}

package main

import "fmt"

func div60(i int) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	fmt.Println(60 / i)
}

func main() {
	for _, val := range []int{1, 2, 3, 0, 5} {
		div60(val)
	}
}

package main

import "fmt"

func main() {
	var i any
	i = 1
	fmt.Println(i) // 1

	i = "hello"
	fmt.Println(i) // hello

	i = struct {
		name string
	}{name: "world"}
	fmt.Println(i) // {world}
}

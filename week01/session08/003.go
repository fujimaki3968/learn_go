package main

import "fmt"

func doTypeSwitch(i any) {
	switch v := i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknown")
	}
}

func main() {
	doTypeSwitch(nil)
	doTypeSwitch(1)
	doTypeSwitch("hello")
	doTypeSwitch(struct {
		name string
	}{name: "world"})
}

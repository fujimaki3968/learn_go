package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s (%d)", p.FirstName, p.LastName, p.Age)
}

func main() {
	p := Person{"Taro", "Yamada", 42}
	fmt.Println(p) // Taro Yamada (42)

	output := p.String()
	fmt.Println(output) // Taro Yamada (42)
}

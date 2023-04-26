package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
	}

	makiart := person{
		name: "makiart",
		age:  20,
	}

	fmt.Println(makiart) // {makiart 20}

	var pet struct {
		name string
		age  int
	}

	pet.name = "makiart's pet"
	pet.age = 1

	fmt.Println(pet) // {makiart's pet 1}
}

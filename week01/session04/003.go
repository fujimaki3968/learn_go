package main

import (
	"fmt"
	"sort"
)

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Bob", 30},
		{"Alice", 20},
		{"Charlie", 40},
	}

	fmt.Println("initial:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("sorted:", people)

	twoBase := makeMult(2)
	threeBase := makeMult(3)

	for i := 0; i < 3; i++ {
		fmt.Println(i, twoBase(i), threeBase(i))
	}
}

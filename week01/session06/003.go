package main

import "fmt"

type IntTree struct {
	Value       int
	Left, Right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{Value: val}
	}
	if val < it.Value {
		it.Left = it.Left.Insert(val)
	} else {
		it.Right = it.Right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.Value:
		return it.Left.Contains(val)
	case val > it.Value:
		return it.Right.Contains(val)
	default:
		return true
	}
}

func main() {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(7)
	it = it.Insert(1)
	it = it.Insert(4)
	it = it.Insert(6)

	fmt.Println(it.Contains(5))  // true
	fmt.Println(it.Contains(3))  // true
	fmt.Println(it.Contains(7))  // true
	fmt.Println(it.Contains(10)) // false˜
}

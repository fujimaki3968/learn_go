package main

import "fmt"

func main() {
	var array01 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array01) // [1 2 3 4 5]

	var array02 = []int{1, 5: 4}
	fmt.Println(array02) // [1 0 0 0 0 4]

	array02 = append(array02, 5)
	fmt.Println(array02) // [1 0 0 0 0 4 5]

	var array03 = make([]int, 5)
	fmt.Println(array03) // [0 0 0 0 0]

	var array04 = make([]int, 0, 5)
	fmt.Println(array04) // []

	array04 = append(array04, 1, 2, 3, 4, 5)
	fmt.Println(array04) // [1 2 3 4 5]

	// サブスライス
	var array05 = array04[1:3]
	fmt.Println(array05) // [2 3]

	array04[1] = 100
	fmt.Println(array05) // [100 3] サブスライスも値が変わる
}

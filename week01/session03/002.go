package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	if n == 5 {
		fmt.Println("n is 5")
	} else if n < 5 {
		fmt.Println("n is less than 5")
	} else {
		fmt.Println("n is greater than 5")
	}

	// You can also do this:
	if m := rand.Intn(10); m == 5 {
		fmt.Println("m is 5")
	} else if m < 5 {
		fmt.Println("m is less than 5")
	} else {
		fmt.Println("m is greater than 5")
	}
	// fmt.Println(m) <- errorになる

}

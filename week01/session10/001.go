package main

import (
	"errors"
	"fmt"
	"os"
)

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return -1, errors.New("Can't double odd number")
	}
	return i * 2, nil
}

func main() {
	i := 19
	double, err := doubleEven(i)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Double of %d is %d\n", i, double)
}

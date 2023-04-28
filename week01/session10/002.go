package main

import "fmt"

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return -1, fmt.Errorf("Can't double odd number %d", i)
	}
	return i * 2, nil
}

func main() {
	i := 19
	double, err := doubleEven(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Double of %d is %d\n", i, double)
}

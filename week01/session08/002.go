package main

import (
	"fmt"
	"os"
)

type MyInt int

func main() {
	var i any
	var mine MyInt = 20

	i = mine
	i2 := i.(MyInt) //型アサーション

	fmt.Println(i2) // 20

	//wrong pattern
	//i3 := i.(string) <- panic
	//i4 := i.(int) <- panic

	//型アサーションの戻り値は2つ
	i5, ok := i.(string)
	if !ok {
		fmt.Println("not string")
		os.Exit(1)
	}
	fmt.Println("i5 is string:", i5)
}

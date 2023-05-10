package main

import (
	"fmt"

	print "example.com/sample_packages/formatter"
	"example.com/sample_packages/math"
)

func main() {
	num := math.Double(20)
	fmt.Println(num)
	output := print.Format(num)
	fmt.Println(output)
}

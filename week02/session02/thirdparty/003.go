package main

import (
	"fmt"
	"os"

	"github.com/learning-go-book/formatter"
	"github.com/shopspring/decimal"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: amount and percent")
		os.Exit(1)
	}
	amoount, err := decimal.NewFromString(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	percent, err := decimal.NewFromString(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	percent = percent.Div(decimal.NewFromInt(100))
	total := amoount.Add(amoount.Mul(percent)).Round(2)
	fmt.Println(formatter.Space(80, os.Args[1], os.Args[2], total.StringFixed(2)))

}

//❯ ./thirdparty 100 5
//100                                   5                                   105.00

package main

import "fmt"

func main() {
	texts := []string{"開け", "ぼっちざ", "ほげ"}
	for _, v := range texts {
		fmt.Println(v)
		switch v {
		case "開け":
			fmt.Println("ごま")
		case "ぼっちざ":
			fmt.Println("ろっく")
		default:
			fmt.Println("I dont know what you said")
		}
	}

	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		fmt.Println(v)
		switch n := v * 5; {
		case n < 10:
			fmt.Println("v * 5 is less than 10")
		case n > 10:
			fmt.Println("v * 5 is greater than 10")
		default:
			fmt.Println("v * 5 is equal to 10")
		}
	}

	// case内でbreakを使ってもforから抜けられない
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		switch {
		case i%2 == 0:
			fmt.Println("i is even")
		case i%7 == 0:
			fmt.Println("want to end of loop...but can't")
			break // ループから抜けることができない
		default:
			fmt.Println("i is odd")
		}
	}

loop:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		switch {
		case i%2 == 0:
			fmt.Println("i is even")
		case i%7 == 0:
			fmt.Println("end of loop")
			break loop // ループから抜けられる
		default:
			fmt.Println("i is odd")
		}
	}
}

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 1
	for i < 100 {
		fmt.Println(i)
		i *= 4
	}

	j := 0
	for {
		fmt.Println("loop")
		if j > 10 {
			break
		}
		j++
	}

	values := []string{"a", "b", "c"}
	for i, v := range values {
		fmt.Println(i, v)
	}

	text := "Hello World"
	for i, v := range text {
		fmt.Println(i, v, string(v)) // string()で文字列に変換
	}

	texts := []string{"Hello", "World"}
outer:
	for _, v := range texts {
		for _, c := range v {
			if c == 'l' {
				fmt.Println("lが見つかりました")
				continue outer
			}
		}
	}
}

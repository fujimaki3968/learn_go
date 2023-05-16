package main

import "fmt"

// このコードは正しく動作しません
func main() {
	a := []int{3, 6, 9, 12, 15, 18, 21}
	ch := make(chan int, len(a))
	ch2 := make(chan int, len(a))

	for _, v := range a {
		go func() {
			// このコードは正しく動作しません
			// なぜなら、vはforループの中で定義されているため、
			ch <- v * 2
		}()
	}

	fmt.Println("wrong:")

	for i := 0; i < len(a); i++ {
		fmt.Print(<-ch, " ")
	}

	for _, v := range a {
		go func(val int) {
			ch2 <- val * 2
		}(v)
	}

	fmt.Println("\nright:")

	for i := 0; i < len(a); i++ {
		// 入る順番は保証されないが、正しく動作する
		fmt.Print(<-ch2, " ")
	}
}

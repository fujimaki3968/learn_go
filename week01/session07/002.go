package main

type Inner struct {
	Val int
}

type Outer struct {
	Inner
	Val int
}

func main() {
	o := Outer{
		Inner: Inner{Val: 1},
		Val:   2,
	}
	println(o.Val)       // 2
	println(o.Inner.Val) // 1
}

package main

func failedUpdate(px *int) {
	x2 := 100
	px = &x2
}

func update(px *int) {
	*px = 100
}

func main() {
	x := 10
	px := &x
	failedUpdate(px)
	println(x) // 10 ポインタ自体を書き換えることはできない
	update(px)
	println(x) // 100
}

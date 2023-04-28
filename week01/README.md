# week 01

やったこと
「はじめてのGo言語入門」を読む

## memo

### 配列について

Go言語の性質上appendで配列のサイズを大きくしていくとキャパシティが無駄に大きくなるのである程度大きさがわかっているのならmake関数を使用する。

スライスのサブスライスは記憶を共有しているのでスライスの値を変えるとサブスライスも変わる。またappendも反映されてよくわからなくなるので注意が必要。

### シャドーイングについて

パッケージを誤ってシャドーイングすると途中で壊れるので注意が必要。（シャドーイング検知のGoツールを利用するのがおすすめ）

このような頭の悪いこともできるのでlintを入れておこう
```go
fmt.Println(true) // true
true := 10
fmt.Println(true) // 10 <- :P
```

### if

このように書けてすこし便利

```go
if m := rand.Intn(10); m == 5 {
    fmt.Println("m is 5")
} else if m < 5 {
    fmt.Println("m is less than 5")
} else {
    fmt.Println("m is greater than 5")
}
```

### map

mapの変数をfor-rangeすると毎回出力順序が変わる。これは，セキュリティを意識した仕様（順番が固定されていると仮定してコードを書く人間を消すため）

### ポインタ

nilポインタは更新できない

関数でポインタ自体を更新できない

```go
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
```

大きな構造体を関数の引数として使う場合はポインタで渡した方がよい！


### メソッド

これ面白い

```go
type Counter struct {
	total       int
	lastUpdated time.Time
}

// wrong: func (c Counter) Increment() {
// correct: func (c *Counter) Increment() {
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last_updated: %v", c.total, c.lastUpdated)
}

func main() {
	var c Counter
	fmt.Println(c) // total: 0, last_updated: 0001-01-01 00:00:00 +0000 UTC

	// (&c).Increment()
	c.Increment() // <- これで実行できる

	fmt.Println(c) // total: 1, last_updated: hogehoge
}
```

### 埋め込み

「クラス継承よりオブジェクト合成のほうがよい」という主張があり，Goではそれを採用しているらしい。(GoF Gang of Four)

たしかに使いやすいし個人的に好きかも

埋め込みのフィールドが重複しているときは明示的に指定する必要がある。

```go
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
```

### interface

型を定義せずに任意の型の値を記憶できてしまう

ただ関数群の扱いとかで楽になるところもある

```go
func main() {
	var i any
	i = 1
	fmt.Println(i) // 1

	i = "hello"
	fmt.Println(i) // hello

	i = struct {
		name string
	}{name: "world"}
	fmt.Println(i) // {world}
}
```
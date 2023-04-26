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
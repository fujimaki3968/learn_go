# week04

## test

### シンプルなテスト
`該当ファイル名_test.go`でテストを作成でき、`go test`で実行できる。

```go
package main

import (
    "testing"
)

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, world"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

#### 値の設定

testing.M型のオブジェクトを受け取る関数を作成することで、テスト実行前に値を設定できる。

```go
var hoge string
var fuga string

func TestMain(m *testing.M) {
    // do stuff before tests
    hoge := "hoge"
    fuga := "fuga"

    exitCode := m.Run()

    // do stuff after tests

    os.Exit(exitCode)
}

func FirstTest(t *testing.T) {
    // use hoge and fuga
    fmt.Println(hoge) // hoge
}
```

#### t.Cleanup

テスト実行後に実行したい処理を登録できる。

以下のように書くことにより、テスト実行後にファイルを削除することができる。

```go
func createFile(t *testing.T) (string, error) {
    f, err := os.Create("test.txt")
    if err != nil {
        return "", err
    }

    // file をほげほげする

    t.Cleanup(func() {
        f.Close()
        os.Remove(f.Name())
    })
    return f.Name(), nil
}
```

### cmpを使ったテスト

cmpを使うことで、構造体の比較をいい感じに行うことができる。

```go
cmp.Diff(got, want)
```


### coverage

`go test -cover`でカバレッジを確認できる。

`go test -coverprofile=coverage.out`でカバレッジをファイルに出力できる。

`go tool cover -html=coverage.out`でカバレッジをhtmlで確認できる。

### テーブルテスト

テストケースをテーブルで定義することで、テストを簡潔に書くことができる。

```go
func TestArea(t *testing.T) {
    areaTests := []struct {
        name string
        shape Shape
        want float64
    }{
        {name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
        {name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
        {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
    }

    for _, tt := range areaTests {
        // using tt.name from the case
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Area()
            if got != tt.want {
                t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
            }
        })
    }
}
```

### benchmark

`go test -bench=.`でベンチマークを実行できる。

`go test -bench=. -benchmem`でベンチマークのメモリ使用量も確認できる。


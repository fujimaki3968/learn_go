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
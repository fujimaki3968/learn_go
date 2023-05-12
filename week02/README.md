# week 02

## memo

### deferを使ったエラーのラップ

deferを使ってエラーをラップするとエラーの発生箇所がわかりやすくなる。

```go
func DoSomeThingsDefer(val1 int, val2 string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomeThingsDefer: %w", err)
		}
	}()
	val3, err := doThing1(val1)
	if err != nil {
		return "", err
	}
	val4, err := doThing2(val2)
	if err != nil {
		return "", err
	}
	return doThing3(val3, val4)
}
```

### recoverエラー

panicをrecoverでキャッチするとエラーを返すことができる。

```go
defer func() {
    if p := recover(); p != nil {
        fmt.Println(p)
    }
}()
```

### go module

go moduleを使うと依存関係を管理できる。

```bash
go mod init example.com/hogehoge

go mod tidy # 依存関係を整理する
```

#### パッケージの命名記法

`util.ExtractNames`, `util.FormatNames`ではなく`extract.Names`, `format.Names`のようにパッケージ名は機能に合わせて命名する方が良い。

#### パッケージ名のオーバーライド

`import`でパッケージ名をオーバーライドできる。

```go
import (
	"fmt"
	myfmt "example.com/hogehoge/fmt"
)
```
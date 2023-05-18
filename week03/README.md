# week 03

## 並行処理

### goroutineとfor

for内のgoroutineは、参照タイミングによって値が変化するので、意図した動作をしないことがある。

```go
for _, v := range a {
    go func() {
        // このコードは正しく動作しません
        // なぜなら、vはforループの中で定義されているため、
        ch <- v * 2
    }()
}
```

## 標準ライブラリ

### ファイルの読み書き

```go
// ファイルを開く
file, err := os.Open("test.txt")
```

### json

いい感じに構造体とjsonを相互変換できて気持ちがいい。標準で入っているから安心。型の拡張も可能。

個人的にポイント高い。

### net/http

簡単にサーバーを立てられる。Pythonの標準ライブラリよりも簡単かつ高機能（な気がする）

結構シンプルに描けるから好き。

こんな感じに二つのAPIを作れる。

ミドルウェアもあって有能。

```go
personMux := http.NewServeMux()
personMux.HandleFunc("/greet",
    func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world!"))
    })

dogMux := http.NewServeMux()
dogMux.HandleFunc("/greet",
    func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Bow, wow!"))
    })

mux := http.NewServeMux()
mux.Handle("/person/", http.StripPrefix("/person", personMux))
mux.Handle("/dog/", http.StripPrefix("/dog", dogMux))

s := http.Server{
    Addr:         ":8080",
    ReadTimeout:  30 * time.Second,
    WriteTimeout: 90 * time.Second,
    IdleTimeout:  120 * time.Second,
    Handler:      mux,
}
fmt.Println("ブラウザで http://localhost:8080 にアクセスしてください")
err := s.ListenAndServe()
if err != nil {
    if err != http.ErrServerClosed {
        panic(err)
    }
}
```
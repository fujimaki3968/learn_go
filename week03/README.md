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


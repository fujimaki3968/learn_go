package main

import (
	"context"
	"fmt"
	"net/http"
)

func logic(ctx context.Context, info string) (string, error) {
	// do something
	text := info + " is done"
	return text, nil
}

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		req = req.WithContext(ctx)
		handler.ServeHTTP(rw, req)
	})
}

func handler(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	err := req.ParseForm()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	data := req.FormValue("data")
	result, err := logic(ctx, data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write([]byte(result))
}

type ServiceCaller struct {
	clieent *http.Client
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", handler)
	s := http.Server{
		Addr:    ":8080",
		Handler: Middleware(mux),
	}
	fmt.Println("ブラウザで http://localhost:8080 にアクセスしてください")
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

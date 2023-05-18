package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
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

}

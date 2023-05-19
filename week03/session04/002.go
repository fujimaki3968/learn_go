package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

func slowServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.Write([]byte("Slow server"))
	}))
	return s
}

func fastServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("error") == "true" {
			w.Write([]byte("Error"))
			return
		}
		w.Write([]byte("ok"))
	}))
	return s
}

var client = http.Client{}

func callBoth(ctx context.Context, errVal string, slowURL string, fastURL string) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := callServer(ctx, "slow", slowURL)
		if err != nil {
			cancel()
		}
	}()

	go func() {
		defer wg.Done()
		err := callServer(ctx, "fast", fastURL+"?error="+errVal)
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()
	fmt.Println("done")
}

func callServer(ctx context.Context, label string, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(label, "error:", err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(label, "error:", err)
		return err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(label, "error:", err)
		return err
	}
	result := string(data)
	if result != "" {
		fmt.Println(label, "result:", result)
	}
	if result == "Error" {
		return fmt.Errorf("error occurred")
	}
	return nil
}

func main() {
	ss := slowServer()
	defer ss.Close()
	fs := fastServer()
	defer fs.Close()

	ctx := context.Background()
	// callBoth(ctx, "true", ss.URL, fs.URL)
	callBoth(ctx, "false", ss.URL, fs.URL)
}

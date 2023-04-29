package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func main() {
	err := fileChecker("non-existent-file.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("file does not exist")
		}
	}
}

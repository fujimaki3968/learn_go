package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("ready to test")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("all tests done")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	fmt.Println("TestFirst uses testTime here", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("TestSecond uses testTime here", testTime)
}

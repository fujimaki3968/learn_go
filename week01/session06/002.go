package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

// wrong: func (c Counter) Increment() {
// correct: func (c *Counter) Increment() {
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last_updated: %v", c.total, c.lastUpdated)
}

func main() {
	var c Counter
	fmt.Println(c) // total: 0, last_updated: 0001-01-01 00:00:00 +0000 UTC

	// (&c).Increment()
	c.Increment()

	fmt.Println(c) // total: 1, last_updated: hogehoge
}

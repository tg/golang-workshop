package main

import (
	"fmt"
	"time"
)

func count(name string, delay time.Duration) {
	n := 0
	for {
		time.Sleep(delay)
		fmt.Printf("%6s: %3d\n", name, n)
		n++
	}
}

func main() {
	go count("tiger", 2*time.Second)
	go count("turtle", 4*time.Second)

	// in the main goroutine
	count("gopher", time.Second)
}

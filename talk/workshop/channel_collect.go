package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func doSomeJob(n int, ch chan string) {
	time.Sleep(time.Second) // say it takes a second
	ch <- strings.Repeat(strconv.Itoa(n), n)
}

func main() {
	results := make(chan string)

	for n := 0; n < 20; n++ {
		go doSomeJob(n, results)
	}

	// Collect
	for n := 0; n < 20; n++ {
		fmt.Println(<-results)
	}
}

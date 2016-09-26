package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for {
			ch <- "ping"
			fmt.Println(<-ch)
		}
	}()

	for {
		fmt.Println(<-ch)
		time.Sleep(time.Second)
		ch <- "pong"
	}
}

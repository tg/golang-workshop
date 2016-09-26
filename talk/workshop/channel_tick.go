package main

import (
	"log"
	"time"
)

func main() {
	tick := time.Tick(time.Second)

	for n := 0; n < 10; n++ {
		go print(n, tick)
	}

	time.Sleep(time.Minute)
}

func print(n int, tick <-chan time.Time) {
	for {
		<-tick
		log.Println(n)
	}
}

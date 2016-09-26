package main

import (
	"fmt"
	"time"
)

func main() {
	for _, name := range []string{"gopher", "turtle", "tiger"} {
		go func() {
			n := 0
			for {
				fmt.Printf("%6s: %3d\n", name, n)
				n++
				time.Sleep(time.Second)
			}
		}()
	}

	time.Sleep(time.Minute)
}

package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Hello from outer space")
	}()
}

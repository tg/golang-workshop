package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("there is no place like", os.Getenv("HOME"))
	fmt.Println(os.ExpandEnv("$USER: there is no place like $HOME"))
}

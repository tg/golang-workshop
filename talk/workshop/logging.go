package main

import (
	"fmt" // for printing
	"log" // for logging with timestamp etc.
	"os"
	"runtime"
)

func main() {
	fmt.Print("go", 4, "it", "\n")
	fmt.Println("we", "are", "individual", "words")

	fmt.Printf("@%s: %s\n", "bob", "hello")
	fmt.Fprintf(os.Stderr, "booooo👻ooooo!\n")

	log.Printf("Detected %d CPU(s)", runtime.NumCPU())
	log.Fatal("Dying due to a fatal error...")

	log.Print("unreachable code here")
}

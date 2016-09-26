package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	names := []string{"bob", "john", "alfred"}
	err := json.NewEncoder(os.Stdout).Encode(names)

	if err != nil {
		// Problem with JSON encoding or writing to stream
		log.Fatal(err)
	}
}

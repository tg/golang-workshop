package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	names := []string{"bob", "john", "alfred"}

	data, err := json.Marshal(names)
	if err != nil {
		log.Fatal(err)
	}

	// data is of type []byte
	fmt.Println(data)
}

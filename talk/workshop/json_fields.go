package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	FullName string `json:"full_name"`
	Age      int

	password string // unexported field
}

func main() {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t") // for pretty printing

	enc.Encode(&Person{
		FullName: "Rob Pike",
		Age:      60,
		password: "golang2008",
	})
}

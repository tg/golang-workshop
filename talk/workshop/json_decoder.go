package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// code_begin OMIT

type User struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Nicknames []string `json:"nicknames"`
}

func main() {
	data := []byte(`{"name":"Tomasz","age":29,"nicknames":["genius","dumbass"]}`)

	var user User
	err := json.Unmarshal(data, &user)
	//err := json.NewDecoder(bytes.NewReader(data)).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", user)
}

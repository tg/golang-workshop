package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	hello := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, Internet!")
	}

	err := http.ListenAndServe(":8080", http.HandlerFunc(hello))

	log.Fatal(err) // it shouldn't happen!
}

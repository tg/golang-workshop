package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Very polite HTTP handler
	hello := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, Internet!")
	}

	// Run http server on port 8080
	err := http.ListenAndServe(":8080", http.HandlerFunc(hello))

	// Log and die, in case something go wrong
	log.Fatal(err)
}

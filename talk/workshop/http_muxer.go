package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, Internet!")
	})
	mux.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Bye bye, Internet!")
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}

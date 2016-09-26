package main

import (
	"bytes"
	"log"
	"net/http"
)

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data bytes.Buffer // []byte with IO

	n, err := data.ReadFrom(r.Body)
	if err != nil {
		return
	}

	log.Printf("Got %d bytes from %s", n, r.RemoteAddr)
}

func main() {
	http.Handle("/", &Handler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

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
	
	n, err := data.ReadFrom(r.Body) // read body to the buffer
	if err != nil {  
		panic(err) 
	}

	log.Printf("Got %d bytes from %s: %s\n", n, r.RemoteAddr, data.String())
}

func main() {
	http.Handle("/", &Handler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

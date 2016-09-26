package main

import (
	"io"
	"log"
	"net/http"
)

type CopyHandler struct {
}

func (h *CopyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.Copy(w, r.Body)
}

func main() {
	http.Handle("/", &CopyHandler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

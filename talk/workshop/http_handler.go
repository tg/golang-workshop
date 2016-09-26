package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyHandler struct {
	Greeting string
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, %s!", h.Greeting, r.RemoteAddr)
}

func main() {
	http.Handle("/", &MyHandler{
		Greeting: "Yo",
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

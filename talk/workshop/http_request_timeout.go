package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// main_start OMIT
	c := &http.Client{
		Timeout: 2 * time.Second,
	}

	log.Println("Fetching...")
	_, err := c.Get("http://workshop.x7467.com:1234")

	if err != nil {
		log.Fatal(err)
	}
	// main_end OMIT
}

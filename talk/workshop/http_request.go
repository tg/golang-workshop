package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// main_start OMIT
	// Fetch the URL
	resp, err := http.Get("http://localhost:3999")
	if err != nil {
		log.Fatal(err) // problem with connection or protocol
	}

	// Close the underlying connection when done
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("status: ", resp.Status) // unexpected status (not 200 OK)
	}

	// Print the response
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatal(err)
	}
	// main_end OMIT
}

// Simple base code for the bot.
// You can test it by sending the following data:
//
// 		text=booo&user_name=gopher
//

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Response structure serialisable to JSON
type Response struct {
	Text string `json:"text"`
}

type SlackHandler struct {
}

func replaceWord(word string) (string, error) {
	// Fetch the URL
	resp, err := http.Get("http://workshop.x7467.com:1080/" + word)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// If 404, not such word in the dictionary
	if resp.StatusCode == http.StatusNotFound {
		return "", nil
	}

	// Otherwise expecting 200 OK
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("dict status: %s", resp.Status)
	}

	// Decode response
	var parsed struct {
		Synonyms []string `json:"synonyms"`
	}
	err = json.NewDecoder(resp.Body).Decode(&parsed)

	if err != nil {
		return "", err
	}

	// If found anything, return the first one
	if s := parsed.Synonyms; len(s) > 0 {
		return s[0], nil
	}

	return "", nil
}

// ServeHTTP reads message from Slack and respond with greetings
func (h *SlackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data bytes.Buffer

	// Read request data
	_, err := data.ReadFrom(r.Body)
	if err != nil {
		return
	}

	// Parse incoming data
	q, err := url.ParseQuery(data.String())

	if err != nil {
		log.Println("wrong data:", err)
		return
	}

	if q.Get("token") != os.Getenv("BOT_TOKEN") {
		log.Println("invalid token")
		http.Error(w, "bad token", http.StatusForbidden)
		return
	}

	// Get text of incoming message
	text := q.Get("text")

	// Replace words, keep in []string
	var words []string
	for _, word := range strings.Split(text, " ") {
		newWord, err := replaceWord(word)
		if err != nil {
			log.Printf("%s: %s", newWord, err)
		}

		if newWord == "" {
			newWord = word
		}
		words = append(words, newWord)
	}

	// Send personalised response
	json.NewEncoder(w).Encode(&Response{
		Text: strings.Join(words, " "),
	})
}

func main() {
	// Listen on address specified in $BOT_ADDR, or :8080 if empty
	addr := os.Getenv("BOT_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	// Run the server
	log.Fatal(http.ListenAndServe(addr, &SlackHandler{}))
}

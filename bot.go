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
)

// Response structure serialisable to JSON
type Response struct {
	Text string `json:"text"`
}

type SlackHandler struct {
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

	// Send personalised response
	json.NewEncoder(w).Encode(&Response{
		Text: fmt.Sprintf("Hello %s, how are you?", q.Get("user_name")),
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

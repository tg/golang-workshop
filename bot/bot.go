// Simple base code for the bot.
// You can test it by sending the following data:
//
// 		text=booo&user_name=gopher
//

package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Response structure serialisable to JSON
type Response struct {
	Text string `json:"text"`
}

// SlackHandler responing to incoming messages
type SlackHandler struct {
	// Converter will be used for converting incoming messages and send as a response.
	converter TextConverter
}

// NewSlackHandler returns new SlackHandler which uses specific TextConverter.
func NewSlackHandler(c TextConverter) *SlackHandler {
	return &SlackHandler{c}
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

	// Get incoming text and convert to response
	resp, err := h.converter.ConvertText(q.Get("text"))

	if err != nil {
		log.Println("convert text:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Send JSON-encoded response
	json.NewEncoder(w).Encode(&Response{
		Text: resp,
	})
}

func main() {
	// Listen on address specified in $BOT_ADDR, or :8080 if empty
	addr := os.Getenv("BOT_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	// Text converter used by bot.
	var converter TextConverter

	// We'll use synonym converter using synonymizer over HTTP.
	converter = &SynonymTextConverter{
		Synonymizer: &WorkshopSynonymizer{
			URL: "",
		},
	}

	// And wrap it with parallelizer to make fast!
	converter = NewParallelTextConverter(converter, " ")

	// Run the server
	log.Fatal(http.ListenAndServe(addr, NewSlackHandler(converter)))
}

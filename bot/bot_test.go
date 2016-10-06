package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/tg/golang-workshop/bot/textconv"
)

func TestSlackHandler(t *testing.T) {
	// Create a request
	data := strings.NewReader("token=secret&user_name=gopher&text=my%20code%20is%20the%20best")
	r, _ := http.NewRequest("POST", "http://test.com", data)

	// Set matching BOT_TOKEN
	os.Setenv("BOT_TOKEN", "secret")

	// Handle request and store result in w
	w := httptest.NewRecorder()

	// Create handler converting text to upper case
	handler := NewSlackHandler(textconv.StringMapFunc(strings.ToUpper))
	handler.ServeHTTP(w, r)

	// Check code
	if w.Code != http.StatusOK {
		t.Fatal(w.Code, w.Body.String())
	}

	// Check response (after trimming new line which is appended to JSON)
	response := strings.TrimRight(w.Body.String(), "\n")
	if response != `{"text":"MY CODE IS THE BEST"}` {
		t.Fatal(response)
	}
}

func TestSlackHandler_bad_token(t *testing.T) {
	// Create a request
	data := strings.NewReader("token=secret&user_name=gopher")
	r, _ := http.NewRequest("POST", "http://test.com", data)

	// Set invalid BOT_TOKEN to make usre we drop it
	os.Setenv("BOT_TOKEN", "invalid")

	// Handle request and store result in w
	w := httptest.NewRecorder()

	// We pass nil converter as it shouldn't be called anyway (will panic if it does)
	handler := NewSlackHandler(nil)
	handler.ServeHTTP(w, r)

	// Check forbidden
	if w.Code != http.StatusForbidden {
		t.Fatal(w.Code, w.Body.String())
	}
}

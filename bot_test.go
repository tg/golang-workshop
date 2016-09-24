package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestSlackHandler(t *testing.T) {
	// Create a request
	data := strings.NewReader("token=secret&user_name=gopher&text=my%20code%20is%20the%20best")
	r, _ := http.NewRequest("POST", "http://test.com", data)

	// Set matching BOT_TOKEN
	os.Setenv("BOT_TOKEN", "secret")

	// Handle request and store result in w
	w := httptest.NewRecorder()

	handler := &SlackHandler{}
	handler.ServeHTTP(w, r)

	// Log response (so it can be viewed with "go test -v")
	response := w.Body.String()
	t.Log(response)

	// Check code
	if w.Code != http.StatusOK {
		t.Fatal(w.Code, w.Body.String())
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

	handler := &SlackHandler{}
	handler.ServeHTTP(w, r)

	// Log response (so it can be viewed with "go test -v")
	response := w.Body.String()
	t.Log(response)

	// Check forbidden
	if w.Code != http.StatusForbidden {
		t.Fatal(w.Code, w.Body.String())
	}
}

func TestReplaceWord(t *testing.T) {
	w, err := replaceWord("code")
	if err != nil {
		t.Fatal(err)
	}
	if w == "" || w == "code" {
		t.Error(w)
	}
}

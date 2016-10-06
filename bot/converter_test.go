package main

import (
	"strings"
	"testing"
)

// ToUpperTextConverter will convert text to upper-case (for testing purposes)
type ToUpperTextConverter struct {
}

func (c *ToUpperTextConverter) ConvertText(text string) (string, error) {
	return strings.ToUpper(text), nil
}

func TestParallelTextConverter(t *testing.T) {
	conv := NewParallelTextConverter(&ToUpperTextConverter{}, " ")

	got, err := conv.ConvertText("hello, my name is dr. greenthumb!")
	if err != nil {
		t.Fatal(err)
	}

	if got != "HELLO, MY NAME IS DR. GREENTHUMB!" {
		t.Error(got)
	}
}

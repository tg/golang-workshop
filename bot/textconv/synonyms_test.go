package textconv

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestSynonymizer implements Synonymizer interface and will use the underlying
// function to return a single synonym.
type TestSynonymizer struct {
	// Func will map words to synonym
	Func func(string) string

	// Err will be returned directly by Synonyms method
	Err error

	// Delay to sleep before returning Synonyms.
	// Can be used to imitate slow responses etc.
	Delay time.Duration
}

func (s *TestSynonymizer) Synonyms(word string) ([]string, error) {
	time.Sleep(s.Delay)

	if s.Func == nil {
		return nil, s.Err
	}
	nw := s.Func(word)
	return []string{nw}, s.Err
}

// SynonymizerHandler returns synonyms for a word using HTTPSynonymizer protocol.
type SynonymizerHandler struct {
	// Synonymizer will be used for fetching synonyms
	Synonymizer Synonymizer
}

func (h *SynonymizerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Path[1:] // word is the path without leading slash
	synonyms, err := h.Synonymizer.Synonyms(word)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Write JSON response
	json.NewEncoder(w).Encode(struct {
		Word     string
		Synonyms []string
	}{
		word,
		synonyms,
	})
}

func TestHTTPSynonymizer(t *testing.T) {
	// Run synonym server returning upper-cased words as synonyms
	server := httptest.NewServer(&SynonymizerHandler{
		Synonymizer: &TestSynonymizer{
			Func: strings.ToUpper,
		},
	})
	defer server.Close()

	s := &HTTPSynonymizer{URL: server.URL}

	got, err := s.Synonyms("gopher")
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != 1 || got[0] != "GOPHER" {
		t.Fatal(got)
	}
}

func TestSynonymTextConverter(t *testing.T) {
	conv := &SynonymTextConverter{
		Synonymizer: &TestSynonymizer{
			Func: strings.ToUpper,
		},
	}

	got, err := conv.ConvertText("hello, my name is dr. greenthumb!")
	if err != nil {
		t.Fatal(err)
	}

	if got != "HELLO, MY NAME IS DR. GREENTHUMB!" {
		t.Error(got)
	}
}

func TestSynonymTextConverter_error(t *testing.T) {
	n := 0
	expectedErr := errors.New("test error")

	// Create test synonymizer returning error at fourth call
	synon := &TestSynonymizer{}
	synon.Func = func(word string) string {
		n++
		synon.Err = nil
		if n == 4 {
			synon.Err = expectedErr
		}
		return strings.ToUpper(word)
	}

	conv := &SynonymTextConverter{
		Synonymizer: synon,
	}

	got, err := conv.ConvertText("hello, my name is dr. greenthumb!")
	if err != expectedErr {
		t.Error(err)
	}

	if got != "HELLO, MY NAME is dr. greenthumb!" {
		t.Error(got)
	}
}

// Dummy text for benchmarking
var loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Fusce auctor tempus elementum. Pellentesque vehicula neque vitae sapien interdum,
id pharetra quam porta. Donec dictum viverra lorem, cursus euismod arcu luctus vel.`

// Capitalised lorem ipsum
var loremIpsumUpper = strings.ToUpper(loremIpsum)

// Text converter for benchmarking, with 1ms delay.
var benchmarkConverter = &SynonymTextConverter{
	Synonymizer: &TestSynonymizer{
		Func:  strings.ToUpper,
		Delay: time.Millisecond,
	},
}

func BenchmarkSynonymTextConverter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		got, err := benchmarkConverter.ConvertText(loremIpsum)
		if err != nil {
			b.Fatal(err)
		}
		if got != loremIpsumUpper {
			b.Error(got)
		}
	}
}

func BenchmarkSynonymTextConverter_parallel(b *testing.B) {
	conv := NewParallelTextConverter(benchmarkConverter, " ")

	for n := 0; n < b.N; n++ {
		got, err := conv.ConvertText(loremIpsum)
		if err != nil {
			b.Fatal(err)
		}
		if got != loremIpsumUpper {
			b.Error(got)
		}
	}
}

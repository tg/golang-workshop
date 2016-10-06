package textconv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

// DefaultSynonymizerURL is a default URL for fetching synonyms over HTTP
const DefaultSynonymizerURL = "http://workshop.x7467.com:1080"

// Synonymizer interface requires function for getting synonyms for a word
type Synonymizer interface {
	Synonyms(word string) ([]string, error)
}

// WorkshopSynonymizer implements fetching synonyms via HTTP requests as described on the workshop.
// Synonyms are being fetched by querying URL/{word} and expects.
type WorkshopSynonymizer struct {
	// Base URL used for fetching synonyms
	URL string

	// HTTP client used for requests. Empty value is fine and usable.
	Client http.Client
}

// Synonyms will fetch synonyms from URL.
func (s *WorkshopSynonymizer) Synonyms(word string) ([]string, error) {
	url := s.URL
	// If URL not specified use default
	if url == "" {
		url = DefaultSynonymizerURL
	}

	resp, err := s.Client.Get(fmt.Sprintf("%s/%s", url, word))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// If 404, not such word in the dictionary
	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	// Otherwise expecting 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dict status: %s", resp.Status)
	}

	// Decode response
	var parsed struct {
		Synonyms []string `json:"synonyms"`
	}
	err = json.NewDecoder(resp.Body).Decode(&parsed)

	return parsed.Synonyms, err
}

// ReplaceWord will get synonyms for a word using Synonymizer and will return first one if any found.
// Otherwise the same word will be returned (including on errors).
func ReplaceWord(s Synonymizer, word string) (newWord string, err error) {
	newWord = word
	syns, err := s.Synonyms(word)

	if len(syns) > 0 {
		newWord = syns[0]
	}

	return
}

// DefaultWordMatcher is a regular expression matching words
var DefaultWordMatcher = regexp.MustCompile(`\w+`)

// SynonymTextConverter replaces words (one by one) in text with synonyms.
// This structure implements TextConverter interface to be used with a handler.
type SynonymTextConverter struct {
	// Synonymizer to be used for text conversion.
	Synonymizer Synonymizer

	// Matcher is a regular expression used for matching words.
	// If empty, DefaultWordMatcher will be used.
	WordMatcher *regexp.Regexp
}

// ConvertText will take words one by one and will replace them with synonyms.
func (c *SynonymTextConverter) ConvertText(text string) (string, error) {
	// Return text itself if synonymizer not set
	if c.Synonymizer == nil {
		return text, nil
	}

	// Use custom matcher or default if not defined
	wordMatcher := c.WordMatcher
	if wordMatcher == nil {
		wordMatcher = DefaultWordMatcher
	}

	// We'll keep the very first error from synonymizer here (if any)
	var err error

	// Now replace words with first synonym found
	newText := wordMatcher.ReplaceAllStringFunc(text, func(word string) string {
		newWord := word
		// Check synonyms only if we didn't encounter any errors before
		if err == nil {
			var candidates []string
			candidates, err = c.Synonymizer.Synonyms(word)
			if err == nil && len(candidates) > 0 {
				newWord = candidates[0]
			}
		}
		return newWord
	})

	return newText, err
}

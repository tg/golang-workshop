// Package textconv implements text conversions for the slack bot,
// including synonym fetcher and parallel executor.
package textconv

import "strings"

// TextConverter is the interface that wraps ReplaceText method.
// It's being used for converting string from one form to another.
type TextConverter interface {
	// ConvertText takes text and replaces it with something else.
	ConvertText(text string) (string, error)
}

// StringMapFunc is a helper type implementing TextConverter
// by calling the string->string function and returning a nil error.
// This is helpful if we want to use some functions from strings package,
// for example StringMapFunc(strings.ToUpper) will yield TextConverter
// converting text to its upper case version.
type StringMapFunc func(string) string

// ConvertText calls underlying function.
func (c StringMapFunc) ConvertText(text string) (string, error) {
	return c(text), nil
}

// ParallelTextConverter splits text into chunks and convert them in parallel using
// the underlying TextConverter.
type ParallelTextConverter struct {
	// actual converter to be used
	converter TextConverter

	// splitBy defines string by which text will be split up for parallel execution
	splitBy string
}

// NewParallelTextConverter returns text converter which splits text into chunks
// (using splitBy as a separator) and converts them in parallel using tc.
func NewParallelTextConverter(tc TextConverter, splitBy string) *ParallelTextConverter {
	return &ParallelTextConverter{tc, splitBy}
}

// ConvertText converts text or returns an empty string along with an error.
func (c *ParallelTextConverter) ConvertText(text string) (string, error) {
	chunks := strings.Split(text, c.splitBy)

	// convertedChunk will be created by the worker
	type convertedChunk struct {
		N    int    // chunk index
		Text string // chunk text
		Err  error  // error during conversion, if any
	}

	// create (empty) array of converted chunks
	converted := make([]string, len(chunks))

	// Channel for sending the finished jobs
	done := make(chan convertedChunk)

	// Spawn converters
	for n := range chunks {
		go func(n int) {
			// Convert using underlying converter
			newText, err := c.converter.ConvertText(chunks[n])
			// Send result
			done <- convertedChunk{
				N:    n,
				Text: newText,
				Err:  err,
			}
		}(n)
	}

	// Collect results and put into the right slot
	for range chunks {
		chunk := <-done
		if chunk.Err != nil {
			return "", chunk.Err
		}
		converted[chunk.N] = chunk.Text
	}

	return strings.Join(converted, c.splitBy), nil
}

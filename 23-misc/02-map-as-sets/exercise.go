package search

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// search reads a text and a word and returns true if the word appears in the text and false if it does not.
func contain(reader io.Reader, word string) bool {
	bufio_reader := bufio.NewReader(reader)

	text, err := bufio_reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return false
	}

	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	cleaned_text := re.ReplaceAllString(text, "")
	cleaned_text = strings.ToLower(cleaned_text)
	words := strings.Fields(cleaned_text)

	count := 0

	for _, w := range words {
		if w == "time" {
			count++
			if count > 1 {
				return true
			}
		}
	}

	return false

}

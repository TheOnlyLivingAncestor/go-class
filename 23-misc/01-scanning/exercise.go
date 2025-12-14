package scanning

import (
	"bufio"
	"io"
	"unicode"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// counter reads a text and returns the counted values.
func counter(reader io.Reader) int {
	bufioReader := bufio.NewReader(reader)
	var count int
	for {
		ch, _, err := bufioReader.ReadRune()
		if err != nil {
			break
		}
		if unicode.IsLower(ch) {
			count++
		}
	}
	return count
}

package printer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrinter(t *testing.T) {
	assert.Equal(t, `type: boolean value: true`, printBool(true), "bool")
	assert.Equal(t, `type: boolean value: false`, printBool(false), "bool")
	assert.Equal(t, `type: integer value: 12`, printInt(12), "int")
	assert.Equal(t, `type: integer value: 33`, printInt(33), "int")
	assert.Equal(t, `type: integer in hexadecimal form value: 10`, printHex(16), "hex")
	assert.Equal(t, `type: integer in hexadecimal form value: 9`, printHex(9), "hex")
	assert.Equal(t, `type: integer in hexadecimal form value: ff`, printHex(255), "hex")
	assert.Equal(t, `type: float value: 1.12`, printFloat(1.1234), "float 1")
	assert.Equal(t, `type: float value: 9.97`, printFloat(9.971), "float 2")
	assert.Equal(t, `type: float value: 0.01`, printFloat(0.0123), "float 3")

	assert.Equal(t, "ab", concatStrings("a", "b"), "concat 1")
	assert.Equal(t, "aaabb", concatStrings("aaa", "bb"), "concat 2")

	assert.Equal(t, `type: string value: "ab"`, printConcatStrings("a", "b"), "print-concat 1")
	assert.Equal(t, `type: string value: "aaabb"`, printConcatStrings("aaa", "bb"), "print concat 2")
}

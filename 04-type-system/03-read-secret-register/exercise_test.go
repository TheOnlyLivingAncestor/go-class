package readsecretregister

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseChannelControlRegister(t *testing.T) {
	a,b,c,d := parseChannelControlRegister(2868124196)
	assert.Equal(t, []byte{0x24, 0x1A, 0xAA, 0xF4}, []byte{a,b,c,d})

	a,b,c,d = parseChannelControlRegister(0xdeadbeef)
	assert.Equal(t, []byte{0xEF, 0xBE, 0xDE, 0xAD}, []byte{a,b,c,d})

	a,b,c,d = parseChannelControlRegister(0x01234567)
	assert.Equal(t, []byte{0x67, 0x45, 0x01, 0x23}, []byte{a,b,c,d})
}

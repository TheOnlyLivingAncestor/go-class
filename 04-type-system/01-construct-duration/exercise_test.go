package constructduration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructDuration(t *testing.T) {
	assert.Equal(t, `1.002s`, constructDuration(1, 2).String())
	assert.Equal(t, `4.007s`, constructDuration(4, 7).String())
	assert.Equal(t, `12.023s`, constructDuration(12, 23).String())
}

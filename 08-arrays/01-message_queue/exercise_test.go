package messagequeue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageQueue(t *testing.T) {
	x := messageQueue("Hello", "World", "Again")
	assert.Equal(t, "Again", x[0])
	assert.Equal(t, "Hello", x[1])
	assert.Equal(t, "World", x[2])
}

package sleepSort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSleepSort(t *testing.T) {
	input := []uint{47, 1, 10, 23, 42, 7}
	result := []uint{47, 42, 23, 10, 7, 1}
	assert.Equal(t, result, reverseSleepSort(input), "reverse sleep-sort")
}

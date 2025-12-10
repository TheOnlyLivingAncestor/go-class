package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	assert.Equal(t, 2, gmean(1.2, 2.4), "basic test 0")
	assert.Equal(t, 19, gmean(3.67, 100.0), "basic test 1")
	assert.Equal(t, 71, gmean(34.91, 144.02), "basic test 2")

	val, err := gmeanString("3.67", "100.0")
	assert.NoError(t, err, "string parse test 1")
	assert.Equal(t, 19, val, "string test 1")

	val, err = gmeanString("34.91", "144.02")
	assert.NoError(t, err, "string parse test 2")
	assert.Equal(t, 71, val, "string test 2")

	val, err = gmeanString("dummmy", "100.0")
	assert.Error(t, err, "string parse test x")
	assert.Equal(t, 0, val, "string test 1")

	val, err = gmeanString("1.0", "dummmy")
	assert.Error(t, err, "string parse test y")
	assert.Equal(t, 0, val, "string test 2")
}

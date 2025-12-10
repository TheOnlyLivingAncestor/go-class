package logicalops

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func dm(x, y bool) bool { return !(x && y) }

func TestInverse(t *testing.T) {
	assert.Equal(t, false, inverse(true), "inv 1")
	assert.Equal(t, true, inverse(false), "inv 1")
}

func TestOR(t *testing.T) {
	assert.Equal(t, true, or(true, true), "op1")
	assert.Equal(t, true || false, or(true, false), "op2")
	assert.Equal(t, false || true, or(false, true), "op3")
	assert.Equal(t, false, or(false, false), "op4")
}

func TestDeMorgan(t *testing.T) {
	assert.Equal(t, dm(true,  true),  deMorgan(true, true), "op1")
	assert.Equal(t, dm(true,  false), deMorgan(true, false), "op1")
	assert.Equal(t, dm(false, true),  deMorgan(false, true), "op1")
	assert.Equal(t, dm(false, false), deMorgan(false, false), "op1")
}

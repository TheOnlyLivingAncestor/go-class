package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcFactorialOrSumOrAbs(t *testing.T) {
	
	
	
	assert.Equal(t, 10, calcAbs(10))
	assert.Equal(t, 4, calcAbs(-4))
	assert.Equal(t, 1, calcAbs(-1))
	
}

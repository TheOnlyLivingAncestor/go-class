package pointerbasic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveValue(t *testing.T) {
	

	
	var x string
	x = "100"
	assert.Equal(t, x, retrieveValue(&x))
	x = "aaa"
	assert.Equal(t, x, retrieveValue(&x))
	x = "Joe"
	assert.Equal(t, x, retrieveValue(&x))
	x = ""
	assert.Equal(t, x, retrieveValue(&x))
	

	
}

package digits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigits(t *testing.T) {
	
	
	
		assert.Equal(t, 7, multiplyDigits(7))
	
		assert.Equal(t, 1, multiplyDigits(111))
	
		assert.Equal(t, 0, multiplyDigits(1000))
	
		assert.Equal(t, 0, multiplyDigits(1307674368))
	
		assert.Equal(t, 508032, multiplyDigits(137674368))
	
	
}

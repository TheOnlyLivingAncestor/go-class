package closures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProxy(t *testing.T) {
	
	
	
	
	onOffOne := proxy(func(inp string) int{return len(inp)})
	onOffTwo := proxy(func(inp string) int{return len(inp)})
	one, err1 := onOffOne("alma")
	two, err2 := onOffTwo("ban")
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Equal(t, 4, one)
	assert.Equal(t, 3, two)
	one, err1 = onOffOne("alma")
	two, err2 = onOffTwo("ban")
	assert.NotNil(t, err1)
	assert.NotNil(t, err2)
	assert.Equal(t, 0, one)
	assert.Equal(t, 0, two)
	
}
package repaint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepaintColor(t *testing.T) {
	c, err := repaintColor("chartreuse")
	assert.Equal(t, "magenta", c)
	assert.NoError(t, err)

	c, err = repaintColor("magenta")
	assert.Equal(t, "chartreuse", c)
	assert.NoError(t, err)

	c, err = repaintColor("xxx")
	assert.Equal(t, "", c)
	assert.Error(t, err)
}

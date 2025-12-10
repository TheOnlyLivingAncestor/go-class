package secretprotocolheader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePublishFixHeader(t *testing.T) {
	assert.Equal(t, byte(72), createPublishFixHeader(false, false, false))
	assert.Equal(t, byte(73), createPublishFixHeader(false, false, true))
	assert.Equal(t, byte(74), createPublishFixHeader(false, true, false))
	assert.Equal(t, byte(75), createPublishFixHeader(false, true, true))
	assert.Equal(t, byte(88), createPublishFixHeader(true, false, false))
	assert.Equal(t, byte(89), createPublishFixHeader(true, false, true))
	assert.Equal(t, byte(90), createPublishFixHeader(true, true, false))
	assert.Equal(t, byte(91), createPublishFixHeader(true, true, true))
}

package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"bytes"
)

func TestContain(t *testing.T) {
	var stdin bytes.Buffer
	stdin.WriteString("The great secret known to Apollonius of Tyana, Paul of Tarsus, Simon Magus, Asklepios, Paracelsus, Boehme and Bruno is that: we are moving backward in time. The universe in fact is contracting into a unitary entity which is completing itself. Decay and disorder are seen by us in reverse, as increasing. These healers learned to move forward in time, which is retrograde to us.")

	assert.Equal(t, true, contain(&stdin, "TIME"), "The result should be: true")
}

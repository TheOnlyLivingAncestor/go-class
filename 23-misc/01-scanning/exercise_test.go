package scanning

import (
	"testing"
	"bytes"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	var stdin bytes.Buffer
	stdin.WriteString("Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts. Separated they live in Bookmarksgrove right at the coast of the Semantics, a large language ocean. A small river named Duden flows by their place and supplies it with the necessary regelialia. It is a paradisematic country, in which roasted parts of sentences fly into your mouth.")
	assert.Equal(t, 314, counter(&stdin), "The two numbers should be equal.")
}

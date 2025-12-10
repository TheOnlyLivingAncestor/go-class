package helloworld

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, "Sawubona Mhlaba!", helloWorld(), "hello world in Zulu")
}

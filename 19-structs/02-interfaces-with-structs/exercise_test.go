package structsinterfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testInfo(e Printable) string {
	return e.Info()
}

func testPageNum(e Printable) int {
	return e.PageNum()
}

func TestStructsInterfaces(t *testing.T) {
	s1 := NewBook("Matthew A. Titmus", "Cloud Native Go", 433)

	assert.Equal(t, s1.Info(), testInfo(s1))
	assert.Equal(t, s1.PageNum(), testPageNum(s1))

	s2 := NewMagazine("Communications of the ACM", "Volume 67, Issue 9", 92)
	assert.Equal(t, s2.Info(), testInfo(s2))
	assert.Equal(t, s2.PageNum(), testPageNum(s2))
}

package pathsplit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitPath(t *testing.T) {
	assert.Equal(t, "static/js/", splitPath("static/js/jquery.js"))
	assert.Equal(t, "", splitPath("multi_langs.js"))
	assert.Equal(t, "css/", splitPath("css/style.css"))
	assert.Equal(t, "css/", splitPath("css/flags.css"))
	assert.Equal(t, "image/", splitPath("image/favicon.ico"))
}

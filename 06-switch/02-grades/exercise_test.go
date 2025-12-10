package grades

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGradeExam(t *testing.T) {
	assert.Equal(t, 5, gradeExam(100))
	assert.Equal(t, 5, gradeExam(90))
	assert.Equal(t, 4, gradeExam(75))
	assert.Equal(t, 3, gradeExam(50))
	assert.Equal(t, 2, gradeExam(30))
	assert.Equal(t, 0, gradeExam(0))
}

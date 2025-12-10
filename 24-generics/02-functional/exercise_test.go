package functional

import (

	"strconv"

	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenericFunction(t *testing.T) {
	
	
	
	valuesInt := []int{2,5,8,88,12}
	valuesStr := []string{"22","3","4","56","alma","korte","banan"}
	filteredInt := filter(valuesInt, func(i int) bool { return i%2 == 0 })
	filteredStr := filter(valuesStr, func(s string) bool {
		_,err := strconv.Atoi(s)
		return err != nil
	})
	assert.Equal(t, []int{2,8,88,12},filteredInt)
	assert.Equal(t, []string{"alma","korte","banan"},filteredStr)
	
	
}

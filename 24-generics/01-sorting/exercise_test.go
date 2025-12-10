package sorting

import (
	
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSorting(t *testing.T) {
	

	

	
	arrTypeOne := []uint8{1,2,7,5,4,6,3,9,8}
	arrTypeTwo := []int32{3,-1,22,-11,8080,5432,123,53,179,-66}
	assert.Equal(t, []uint8{1,2,3,4,5,6,7,8,9}, sortSlice(arrTypeOne))
	assert.Equal(t, []int32{-66,-11,-1,3,22,53,123,179,5432,8080},sortSlice(arrTypeTwo))
	

	

	
}

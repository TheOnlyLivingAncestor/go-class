package richterscale

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDescribeEarthquake(t *testing.T) {

	assert.Equal(t, "micro", describeEarthquake(0.5))

	assert.Equal(t, "very minor", describeEarthquake(2.5))

	assert.Equal(t, "minor", describeEarthquake(3))

	assert.Equal(t, "light", describeEarthquake(4.5))

	assert.Equal(t, "moderate", describeEarthquake(5))

	assert.Equal(t, "strong", describeEarthquake(6))

	assert.Equal(t, "major", describeEarthquake(7))

	assert.Equal(t, "great", describeEarthquake(8))

	assert.Equal(t, "massive", describeEarthquake(11))

}

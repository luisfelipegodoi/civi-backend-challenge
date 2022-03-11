package models

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestClassifyDistance(t *testing.T) {
	distances := []*Points{
		{Distance: 2},
		{Distance: 1},
		{Distance: 2},
		{Distance: 4},
	}

	sort.Sort(DistanceClassify(distances))

	assert.Equal(t, distances[0].Distance, 1)
	assert.Equal(t, distances[1].Distance, 2)
	assert.Equal(t, distances[2].Distance, 2)
	assert.Equal(t, distances[3].Distance, 4)
}

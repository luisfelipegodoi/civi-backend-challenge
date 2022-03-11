package services

import (
	"civi-backend-challenge/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPoints(t *testing.T) {
	data := []*models.Points{
		{X: 5, Y: 2},
		{X: 4, Y: 2},
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 13, Y: 11},
	}
	
	s := NewPointService(data)

	list := s.GetPoints(3, 1, 5)

	assert.Len(t, list, 4)

	assert.Equal(t, list[0].Distance, 0)
	assert.Equal(t, list[1].Distance, 1)
	assert.Equal(t, list[2].Distance, 2)
	assert.Equal(t, list[3].Distance, 3)
}

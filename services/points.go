package services

import (
	"civi-backend-challenge/models"
	"math"
	"sort"
)

type PointsService interface {
	GetPoints(x, y, distance int) []*models.Points
}

type pointsService struct {
	points []*models.Points
}

func NewPointService(p []*models.Points) *pointsService {
	return &pointsService{
		points: p,
	}
}

func (s *pointsService) GetPoints(x, y, distance int) []*models.Points {
	var list []*models.Points

	param := &models.Points{
		X:        x,
		Y:        y,
		Distance: distance,
	}

	for _, v := range s.points {
		distance := calculate(param, v)
		if param.Distance >= distance {
			list = append(list, &models.Points{
				X:        v.X,
				Y:        v.Y,
				Distance: distance,
			})
		}
	}

	sort.Sort(models.DistanceClassify(list))

	return list
}

func calculate(a, b *models.Points) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

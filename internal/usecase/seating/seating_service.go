package seatingservice

import (
	"cinema/internal/entity"
)

type SeatingService struct{}

func NewSeatingService() *SeatingService {
	return &SeatingService{}
}

func (st *SeatingService) IsValidSeatGroup(group []entity.Seat, reserved []entity.Seat, minDistance int) bool {

	return isValidSeatGroup(group, reserved, minDistance)
}

func isValidSeatGroup(group []entity.Seat, reserved []entity.Seat, minDistance int) bool {
	for _, g := range group {
		for _, r := range reserved {
			if manhattanDistance(g, r) < minDistance {
				return false
			}
		}
	}
	return true
}

func manhattanDistance(a, b entity.Seat) int {
	return abs(a.Row-b.Row) + abs(a.Column-b.Column)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

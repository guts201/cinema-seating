package seatingservice

import (
	"cinema/internal/entity"
)

type SeatingService struct{}

func NewSeatingService() *SeatingService {
	return &SeatingService{}
}

func (st *SeatingService) IsValidSeatGroup(numRows, numCols int, group []entity.Seat, reserved []entity.Seat, minDistance int) bool {

	return isValidSeatGroup(numRows, numCols, group, reserved, minDistance)
}

func (st *SeatingService) GetAvailableSeats(numRows, numCols int, reserved []entity.Seat, minDistance int) []entity.Seat {
	return getAvailableSeats(numRows, numCols, reserved, minDistance)
}

func isValidSeatGroup(numRows, numCols int, group []entity.Seat, reserved []entity.Seat, minDistance int) bool {
	for _, g := range group {
		if g.Row < 0 || g.Row >= numRows || g.Column < 0 || g.Column >= numCols {
			return false
		}
	}

	// check each seat
	for _, g := range group {
		for _, r := range reserved {
			if distance(g, r) <= minDistance {
				return false
			}
		}
	}
	return true
}

func getAvailableSeats(numRows, numCols int, reserved []entity.Seat, minDistance int) []entity.Seat {
	var result []entity.Seat

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			current := entity.Seat{Row: row, Column: col}
			isReserved := false

			// Skip if this seat is reserved
			for _, r := range reserved {
				if r.Row == current.Row && r.Column == current.Column {
					isReserved = true
					break
				}
			}
			if isReserved {
				continue
			}

			// Check Manhattan distance to all reserved seats
			valid := true
			for _, r := range reserved {

				if distance(current, r) <= minDistance {
					valid = false
					break
				}
			}

			if valid {
				result = append(result, current)
			}
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(a, b entity.Seat) int {
	return abs(a.Row-b.Row) + abs(a.Column-b.Column) - 1
}

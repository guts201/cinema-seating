package seatingservice

import (
	"cinema/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidSeatGroup(t *testing.T) {
	type Seat = entity.Seat

	tests := []struct {
		name        string
		group       []Seat
		reserved    []Seat
		minDistance int
		wantValid   bool
	}{
		{
			name:        "valid - no reserved seats",
			group:       []Seat{{Row: 0, Column: 0}, {Row: 0, Column: 1}},
			reserved:    []Seat{},
			minDistance: 2,
			wantValid:   true,
		},
		{
			name:        "valid - far enough",
			group:       []Seat{{Row: 0, Column: 0}, {Row: 0, Column: 1}},
			reserved:    []Seat{{Row: 5, Column: 5}},
			minDistance: 3,
			wantValid:   true,
		},
		{
			name:        "invalid - too close",
			group:       []Seat{{Row: 1, Column: 1}},
			reserved:    []Seat{{Row: 2, Column: 1}},
			minDistance: 2,
			wantValid:   false,
		},
		{
			name:        "invalid - exactly equal to min distance",
			group:       []Seat{{Row: 0, Column: 0}},
			reserved:    []Seat{{Row: 0, Column: 2}},
			minDistance: 3,
			wantValid:   false, // because 2 < 3
		},
		{
			name:        "valid - just enough distance",
			group:       []Seat{{Row: 0, Column: 0}},
			reserved:    []Seat{{Row: 0, Column: 3}},
			minDistance: 3,
			wantValid:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok := isValidSeatGroup(tt.group, tt.reserved, tt.minDistance)
			assert.Equal(t, tt.wantValid, ok)
		})
	}
}

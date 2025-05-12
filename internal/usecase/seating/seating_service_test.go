package seatingservice

import (
	"cinema/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSeatGroup(t *testing.T) {
	type Seat = entity.Seat

	tests := []struct {
		name        string
		row         int
		col         int
		group       []entity.Seat
		reserved    []entity.Seat
		minDistance int
		wantValid   bool
	}{
		{
			name:        "invalid test case",
			row:         5,
			col:         5,
			group:       []entity.Seat{{Row: 3, Column: 2}},
			reserved:    []entity.Seat{{1, 1}},
			minDistance: 2,
			wantValid:   false,
		},
		{
			name:        "valid - far enough",
			row:         6,
			col:         6,
			group:       []entity.Seat{{Row: 0, Column: 0}, {Row: 0, Column: 1}},
			reserved:    []entity.Seat{{Row: 5, Column: 5}},
			minDistance: 3,
			wantValid:   true,
		},
		{
			name:        "invalid - too close",
			row:         3,
			col:         3,
			group:       []entity.Seat{{Row: 1, Column: 1}},
			reserved:    []entity.Seat{{Row: 2, Column: 1}},
			minDistance: 2,
			wantValid:   false,
		},
		{
			name:        "invalid - exactly equal to min distance",
			row:         3,
			col:         3,
			group:       []entity.Seat{{Row: 0, Column: 0}},
			reserved:    []entity.Seat{{Row: 0, Column: 2}},
			minDistance: 3,
			wantValid:   false, // because 1 < 3
		},
		{
			name:        "invalid - exactly equal to min distance 2",
			row:         5,
			col:         5,
			group:       []entity.Seat{{Row: 0, Column: 0}},
			reserved:    []entity.Seat{{Row: 0, Column: 3}},
			minDistance: 2,
			wantValid:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok := isValidSeatGroup(tt.row, tt.col, tt.group, tt.reserved, tt.minDistance)
			assert.Equal(t, tt.wantValid, ok)
		})
	}
}
func TestGetAvailableSeats(t *testing.T) {
	type args struct {
		numRows     int
		numCols     int
		reserved    []entity.Seat
		minDistance int
	}
	tests := []struct {
		name     string
		args     args
		expected []entity.Seat
	}{
		{
			name: "No reserved seats (minDistance=2)",
			args: args{
				numRows:     3,
				numCols:     3,
				reserved:    []entity.Seat{},
				minDistance: 2,
			},
			expected: []entity.Seat{
				{0, 0}, {0, 1}, {0, 2},
				{1, 0}, {1, 1}, {1, 2},
				{2, 0}, {2, 1}, {2, 2},
			},
		},
		{
			name: "1 reserved seat center, minDistance=2",
			args: args{
				numRows: 3,
				numCols: 3,
				reserved: []entity.Seat{
					{1, 1},
				},
				minDistance: 2,
			},
			expected: []entity.Seat{},
		},
		{
			name: "1 reserved corner, minDistance=3",
			args: args{
				numRows: 5,
				numCols: 5,
				reserved: []entity.Seat{
					{0, 0},
				},
				minDistance: 3,
			},
			expected: []entity.Seat{
				// All seats where Manhattan distance to (0,0) >= 3

				{1, 4},
				{2, 3}, {2, 4},
				{3, 2}, {3, 3}, {3, 4},
				{4, 1}, {4, 2}, {4, 3}, {4, 4},
			},
		},
		{
			name: "minDistance=1 (only self blocked)",
			args: args{
				numRows: 2,
				numCols: 2,
				reserved: []entity.Seat{
					{0, 0},
				},
				minDistance: 1,
			},
			expected: []entity.Seat{},
		},
		{
			name: "All seats blocked (small room, big distance)",
			args: args{
				numRows: 3,
				numCols: 3,
				reserved: []entity.Seat{
					{1, 1},
				},
				minDistance: 5,
			},
			expected: []entity.Seat{},
		},
		{
			name: "Complex case with multiple reserved seats",
			args: args{
				numRows: 5,
				numCols: 5,
				reserved: []entity.Seat{
					{1, 1},
				},
				minDistance: 3,
			},
			expected: []entity.Seat{
				{3, 4},
				{4, 3},
				{4, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getAvailableSeats(tt.args.numRows, tt.args.numCols, tt.args.reserved, tt.args.minDistance)
			if len(got) != len(tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, got)
				return
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("Mismatch at index %d: expected %v, got %v", i, tt.expected[i], got[i])
					return
				}
			}
		})
	}
}

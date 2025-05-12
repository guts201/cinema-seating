package entity

import "time"

type Cinema struct {
	ID          int
	Rows        int
	Columns     int
	MinDistance int
	Name        string
}
type Movie struct {
	ID              int
	Title           string
	DurationMinutes int
}

type SeatGroup struct {
	Seats []Seat
}

type Screening struct {
	ID          int
	MovieID     int
	StartTime   time.Time
	MinDistance int
	Row         int
	Column      int
}
type Seat struct {
	Row    int
	Column int
}

type SeatWithId struct {
	Row    int
	Column int
	ID     int
}

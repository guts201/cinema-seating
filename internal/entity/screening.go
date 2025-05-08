package entity

import "time"

type Screening struct {
	ID          int
	MovieID     int
	StartTime   time.Time
	Rows        int
	Columns     int
	MinDistance int
}

package schedule

import (
	"time"
)

type Year struct {
	days []Day
}

type Day struct {
	date time.Time
	busy bool
}

type Booking struct {
	startDate   time.Time
	endDate     time.Time
	daysTaken   int
	breakLength int
}

func Schedule(calendar Year, availableDays int) ([]Booking, error) {
	// TODO: Implement function to determine optimal holiday days
	return []Booking{
		{
			startDate:   time.Now(),
			endDate:     time.Now().AddDate(0, 0, 2),
			daysTaken:   4,
			breakLength: int(time.Since(time.Now().AddDate(0, 0, 2)).Hours()) / 24,
		},
	}, nil
}

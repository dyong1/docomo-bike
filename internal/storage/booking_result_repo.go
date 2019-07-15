package storage

import (
	"time"
)

type BookingResult struct {
	ID        int64
	BikeID    string
	StationID string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookingResultRepository interface {
	Tx() (Tx, error)
	Get(id int64) (*BookingResult, error)
	Add(tx Tx, one *BookingResult) (*BookingResult, error)
	Update(tx Tx, one *BookingResult) (*BookingResult, error)
}

package storage

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type BookingDetail struct {
	ID              int64
	BookerID        string
	StationID       string
	ProgressStatus  int
	BookingResultID null.Int
	StartedAt       time.Time
	CanceledAt      null.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type BookingDetailRepository interface {
	Tx() (Tx, error)
	Get(id int64) (*BookingDetail, error)
	Add(tx Tx, one *BookingDetail) (*BookingDetail, error)
	Update(tx Tx, one *BookingDetail) (*BookingDetail, error)

	GetAllByBookerID(bookerID string) ([]*BookingDetail, error)
}

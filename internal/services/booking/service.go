package booking

import (
	"docomo-bike/internal/libs/docomo/bookbike"
	"time"
)

type Service interface {
	BookBike(bike *Bike) (*BookingResult, error)
}

type Bike = bookbike.Bike

type BookingResult struct {
	BikeID   string
	Passcode string
	BookedAt time.Time
}

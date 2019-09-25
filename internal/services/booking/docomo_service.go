package booking

import (
	"docomo-bike/internal/libs/docomo/bookbike"
	"time"
)

type DocomoService struct {
	BookBikeClient bookbike.Client
}

func (s *DocomoService) BookBike(bike *Bike) (*BookingResult, error) {
	booked, err := s.BookBikeClient.BookBike(bike)
	if err != nil {
		return nil, err
	}
	return &BookingResult{
		BikeID:   booked.BikeID,
		Passcode: booked.Passcode,
		BookedAt: time.Now(),
	}, nil
}

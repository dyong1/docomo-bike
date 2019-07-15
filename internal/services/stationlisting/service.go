package stationlisting

import (
	"docomo-bike/internal/libs/docomo/getstation"
	"docomo-bike/internal/services/auth"
)

type Station struct {
	ID    string
	Name  string
	Bikes []*Bike
}
type Bike struct {
	ID string
}

type Service interface {
	GetStation(auth *auth.Auth, stationID string) (*Station, error)
}

func NewService(getStationClient getstation.Client) Service {
	return &DocomoService{
		GetStationClient: getStationClient,
	}
}

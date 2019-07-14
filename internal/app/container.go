package app

import (
	"context"
	"docomo-bike/internal/libs/docomo/getstation"
	"docomo-bike/internal/libs/docomo/login"
	"docomo-bike/internal/libs/logging"
	"docomo-bike/internal/services/auth"
	"docomo-bike/internal/services/bikebooking"
	"docomo-bike/internal/services/stationlisting"
)

func NewContainer() *Container {
	return &Container{}
}

type Container struct {
	AppLogger        logging.Logger
	HTTPClientLogger logging.Logger

	DocomoClients

	JWTAuthService        auth.JWTService
	BikeBookingService    bikebooking.Service
	StationListingService stationlisting.Service
}
type DocomoClients struct {
	Login      login.Client
	GetStation getstation.Client
}

func (c *Container) Shutdown(ctx context.Context) error {
	return nil
}

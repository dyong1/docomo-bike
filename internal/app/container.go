package app

import (
	"context"
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

	JWTAuthService        auth.JWTService
	BikeBookingService    bikebooking.Service
	StationListingService stationlisting.Service
}

func (c *Container) Shutdown(ctx context.Context) error {
	return nil
}

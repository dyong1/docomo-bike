package app

import (
	"docomo-bike/internal/libs/docomo/getstation"
	"docomo-bike/internal/libs/docomo/login"
	"docomo-bike/internal/libs/logging"
	"docomo-bike/internal/services/auth"
	"docomo-bike/internal/services/booking"
	"docomo-bike/internal/services/listing"
)

func NewContainer() *Container {
	return &Container{}
}

type Container struct {
	AppLogger        logging.Logger
	HTTPClientLogger logging.Logger

	DocomoClients

	JWTAuthService        auth.JWTService
	BikeBookingService    booking.Service
	StationListingService listing.Service
}
type DocomoClients struct {
	Login      login.Client
	GetStation getstation.Client
}

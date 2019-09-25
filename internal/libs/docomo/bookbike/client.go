package bookbike

// EventNo: 25901
// SessionID: ********
// UserID: TYO
// MemberID: *******
// CenterLat:
// CenterLon:
// CycLat: 35.650845
// CycLon: 139.722487
// CycleID: 7960
// AttachID: 18923
// CycleTypeNo: 6
// CycleEntID: TYO

// EventNo: 25901
// SessionID: **********
// UserID: TYO
// MemberID: *******
// CenterLat: 35.650732
// CenterLon: 139.772650
// CycLat: 35.650853
// CycLon: 139.722592
// CycleID: 7287
// AttachID: 18427
// CycleTypeNo: 6
// CycleEntID: TYO

import (
	"docomo-bike/internal/libs/logging"

	"github.com/gojektech/heimdall/httpclient"
)

type Client interface {
	BookBike(bike *Bike) (*BookingResult, error)
}

type Bike struct {
	CenterLat   string
	CenterLon   string
	CycLat      string
	CycLon      string
	CycleID     string
	AttachID    string
	CycleTypeNo string
	CycleEntID  string
}

type BookingResult struct {
	BikeID   string
	BikeNo   string
	Passcode string
}

type ScrappingClient struct {
	HTTPClient *httpclient.Client
	Logger     logging.Logger
}

func (c *ScrappingClient) BookBike(bike *Bike) (*BookingResult, error) {
	panic("Not implemented")
}

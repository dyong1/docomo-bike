package getstation

import (
	"docomo-bike/internal/libs/logging"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gojektech/heimdall/httpclient"

	"github.com/pkg/errors"
)

const (
	getStationEventNo     = "25701"
	getStationAPIEndpoint = "https://tcc.docomo-cycle.jp/cycle/TYO/cs_web_main.php"
	getStationPageSize    = "40"
)
const (
	ToranomonSotoboriStreet    = 10069
	StationRoppongiHills       = 10082
	StationNishiSimbashi1Chome = 10070
)

var (
	bikeIDRegex = regexp.MustCompile(`.*tab_([A-Z]+_\d+).submit`)
)

type Client interface {
	GetStation(userID string, sessionKey string, stationID string) (*Station, error)
}

type ScrappingClient struct {
	HTTPClient *httpclient.Client
	Logger     logging.Logger
}

// EventNo: 25701
// SessionID: ******
// UserID: TYO
// MemberID: ******
// GetInfoNum: 20
// GetInfoTopNum: 1
// ParkingEntID: TYO
// ParkingID: 10131
// ParkingLat: 35.656991
// ParkingLon: 139.740739
type Station struct {
	Name  string
	Bikes []*Bike
}
type Bike struct {
	ID string
}

// EventNo: 25901
// SessionID: ******
// UserID: TYO
// MemberID: ******
// CenterLat: 35.653608
// CenterLon: 139.731380
// CycLat: 35.653637
// CycLon: 139.731303
// CycleID: 13141
// AttachID: 23209
// CycleTypeNo: 14
// CycleEntID: TYO

func (c *ScrappingClient) GetStation(userID string, sessionKey string, stationID string) (*Station, error) {
	data := url.Values{}
	data.Add("EventNo", getStationEventNo)
	data.Add("MemberID", userID)
	data.Add("SessionID", sessionKey)
	data.Add("ParkingID", stationID)
	data.Add("GetInfoNum", getStationPageSize)
	data.Add("GetInfoTopNum", "1")  // First page
	data.Add("UserID", "TYO")       // Required, don't know why TYO is okay
	data.Add("ParkingEntID", "TYO") // Required, don't know why TYO is okay

	dataEncoded := data.Encode()
	headers := http.Header{}
	headers.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.HTTPClient.Post(getStationAPIEndpoint, strings.NewReader(dataEncoded), headers)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	htmlBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	lines := strings.Split(string(htmlBytes), "\n")
	c.Logger.Debugf("Get station response html: %s", spew.Sdump(lines))

	bikes := []*Bike{}
	stationNameLineAt := -1
	for idx, l := range lines {
		if strings.Contains(l, "Port name") {
			stationNameLineAt = idx
		}
		if bid := extractBikeID(l); bid != "" {
			bikes = append(bikes, &Bike{
				ID: bid,
			})
		}
	}
	if stationNameLineAt < 0 {
		return nil, nil
	}

	return &Station{
		Name:  lines[stationNameLineAt+2],
		Bikes: bikes,
	}, nil
}

func extractBikeID(l string) string {
	matches := bikeIDRegex.FindStringSubmatch(l)
	if len(matches) < 2 {
		return ""
	}
	return matches[1]
}

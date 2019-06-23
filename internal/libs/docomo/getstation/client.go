package getstation

import (
	"docomo-bike/internal/auth"
	"docomo-bike/internal/libs/logger"
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
	getStationPageSize    = "100"
)
const (
	StationRoppongiHills       = 10082
	StationNishiSimbashi1Chome = 10070
)

var (
	bikeLineRegex = regexp.MustCompile("<a.*class=\".*cycle_list_btn.*\".*")
)

type Client interface {
	GetStation(auth *auth.Auth, stationID string) (*Station, error)
}

type ScrappingClient struct {
	HTTPClient *httpclient.Client
	Logger     *logger.Logger
}

type Station struct {
	Name       string
	TotalBikes int
}

func (c *ScrappingClient) GetStation(auth *auth.Auth, stationID string) (*Station, error) {
	data := url.Values{}
	data.Add("EventNo", getStationEventNo)
	data.Add("MemberID", auth.UserID)
	data.Add("SessionID", auth.SessionKey)
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

	bikeLines := []string{}
	var stationNameLineAt int
	for idx, l := range lines {
		if bikeLineRegex.MatchString(l) {
			bikeLines = append(bikeLines, l)
		}
		if strings.Contains(l, "Port name") {
			stationNameLineAt = idx
		}
	}

	return &Station{
		Name:       lines[stationNameLineAt+2],
		TotalBikes: len(bikeLines),
	}, nil
}

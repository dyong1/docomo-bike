package login

import (
	"docomo-bike/internal/libs/logger"
	"fmt"
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
	loginEventNo     = "21401"
	loginAPIEndpoint = "https://tcc.docomo-cycle.jp/cycle/TYO/cs_web_main.php"
)

var (
	sessionIDRegex = regexp.MustCompile("value=\"(.+)\"")
)

type Client interface {
	Login(id string, password string) (string, error)
}

type ScrappingClient struct {
	HTTPClient *httpclient.Client
	Logger     *logger.Logger
}

func (c *ScrappingClient) Login(userID string, password string) (string, error) {
	data := url.Values{}
	data.Add("EventNo", loginEventNo)
	data.Add("MemberID", userID)
	data.Add("Password", password)
	dataEncoded := data.Encode()
	headers := http.Header{}
	headers.Add("Content-Type", "application/x-www-form-urlencoded")

	c.Logger.Debugf("Login request header: %s", spew.Sdump(headers))
	c.Logger.Debugf("Login request body: %s", spew.Sdump(dataEncoded))

	res, err := c.HTTPClient.Post(loginAPIEndpoint, strings.NewReader(dataEncoded), headers)
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	htmlBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "")
	}
	lines := strings.Split(string(htmlBytes), "\n")
	c.Logger.Debugf("Login response html: %s", spew.Sdump(lines))

	var sessionIDLine string
	for _, l := range lines {
		if strings.Contains(l, "SessionID") {
			sessionIDLine = l
			break
		}
	}
	if sessionIDLine == "" {
		return "", fmt.Errorf("SessionID is not found in the HTML")
	}
	matches := sessionIDRegex.FindStringSubmatch(sessionIDLine)
	if len(matches) == 0 {
		return "", fmt.Errorf("SessionID is not found in the HTML")
	}

	return matches[1], nil
}

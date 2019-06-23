package docomo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/davecgh/go-spew/spew"

	"github.com/pkg/errors"
)

const (
	LOGIN_EVENT_NO     = "21401"
	LOGIN_API_ENDPOINT = "https://tcc.docomo-cycle.jp/cycle/TYO/cs_web_main.php"
)

func (c *ScrappingClient) Login(userID string, password string) (string, error) {
	data := url.Values{}
	data.Add("EventNo", LOGIN_EVENT_NO)
	data.Add("MemberID", userID)
	data.Add("Password", password)
	dataEncoded := data.Encode()
	headers := http.Header{}
	headers.Add("Content-Type", "application/x-www-form-urlencoded")

	c.Logger.Debugf("Login request header: %s", spew.Sdump(headers))
	c.Logger.Debugf("Login request body: %s", spew.Sdump(dataEncoded))

	res, err := c.HTTPClient.Post(LOGIN_API_ENDPOINT, strings.NewReader(dataEncoded), headers)
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
		}
		break
	}
	if sessionIDLine == "" {
		return "", fmt.Errorf("SessionID is not found in the HTML")
	}
	r := regexp.MustCompile("/value=\"(.+)\"/")
	matches := r.FindStringSubmatch(sessionIDLine)
	if len(matches) == 0 {
		return "", fmt.Errorf("SessionID is not found in the HTML")
	}
	return matches[1], nil
}

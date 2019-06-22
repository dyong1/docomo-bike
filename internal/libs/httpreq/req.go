package httpreq

import (
	"encoding/json"
	"net/http"
)

func JSONBody(r *http.Request, reqBody interface{}) error {
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&reqBody); err != nil {
		return err
	}
	return nil
}

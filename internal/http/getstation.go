package http

import (
	"docomo-bike/internal/services/stationlisting"
	"fmt"
	"net/http"
)

func HandleGetStation(serv stationlisting.Service) http.HandlerFunc {
	var urlParams struct {
		StationID string `urlParam:"stationId"`
	}
	type bike struct {
		ID string `json:"id"`
	}
	type responseBody struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Bikes []bike `json:"bikes"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if err := bindURLParams(&urlParams, r); err != nil {
			badRequest(w, err.Error())
			return
		}
		s, err := serv.GetStation(authFromContext(r.Context()), urlParams.StationID)
		if err != nil {
			internalServerError(w, err.Error())
			return
		}
		if s == nil {
			notFound(w, fmt.Sprintf("Station details are not found [stationID=%s]", urlParams.StationID))
			return
		}

		bb := []bike{}
		for _, b := range s.Bikes {
			bb = append(bb, bike{
				ID: b.ID,
			})
		}
		jsonres(w, responseBody{
			ID:    s.ID,
			Name:  s.Name,
			Bikes: bb,
		})
	}
}

package http

import (
	"docomo-bike/internal/services/bikebooking"
	"encoding/json"
	"net/http"
	"time"
)

func HandleNewBikeBooking(serv bikebooking.Service) http.HandlerFunc {
	var requestBody struct {
		StationID string `json:"stationId"`
	}
	type responseBody struct {
		BookingID        int64     `json:"bookingId"`
		StationID        string    `json:"stationId"`
		Progress         string    `json:"progress"`
		BookingStartedAt time.Time `json:"bookingStartedAt"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&requestBody); err != nil {
			badRequest(w, err.Error())
			return
		}
		b, err := serv.BeginBooking(userIDFromContext(r.Context()), requestBody.StationID)
		if err != nil {
			internalServerError(w, err.Error())
			return
		}
		jsonres(w, responseBody{
			BookingID:        b.ID,
			StationID:        b.StationID,
			Progress:         b.ProgressStatus.String(),
			BookingStartedAt: b.StartedAt,
		})
	}
}

func HandleGetCurrentBikeBooking(serv bikebooking.Service) http.HandlerFunc {
	type responseBody struct {
		BookingID        int64     `json:"bookingId"`
		StationID        string    `json:"stationId"`
		Progress         string    `json:"progress"`
		BookingStartedAt time.Time `json:"bookingStartedAt"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := serv.GetCurrentBookingOfBooker(userIDFromContext(r.Context()))
		if err != nil {
			internalServerError(w, err.Error())
			return
		}
		if b == nil {
			notFound(w, "No current bike booking found")
			return
		}
		jsonres(w, responseBody{
			BookingID:        b.ID,
			StationID:        b.StationID,
			Progress:         b.ProgressStatus.String(),
			BookingStartedAt: b.StartedAt,
		})
	}
}

func HandleCancelBikeBooking(serv bikebooking.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := serv.GetCurrentBookingOfBooker(userIDFromContext(r.Context()))
		if err != nil {
			internalServerError(w, err.Error())
			return
		}
		if b == nil {
			notFound(w, "No current bike booking found")
			return
		}
		if _, err := serv.CancelBooking(b.ID); err != nil {
			internalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func HandleCompleteBookingBike(serv bikebooking.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := serv.GetCurrentBookingOfBooker(userIDFromContext(r.Context()))
		if err != nil {
			internalServerError(w, err.Error())
			return
		}
		if b == nil {
			notFound(w, "No current bike booking found")
			return
		}
		if _, err := serv.CompleteBooking(b.ID); err != nil {
			internalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

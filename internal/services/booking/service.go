package booking

import "time"

type Service interface {
	BookAnyBikeInStation(stationID string) (*BookingDetail, error)
}

type BookingDetail struct {
	ID             int64
	BookerID       string
	StationID      string
	ProgressStatus BookingProgressStatus
	BookingResult  *BookingResult
	StartedAt      time.Time
	CanceledAt     *time.Time
}
type BookingProgressStatus int

func (s BookingProgressStatus) Int() int {
	return int(s)
}
func (s BookingProgressStatus) String() string {
	switch s {
	case ProgressStatusPending:
		return "Pending"
	case ProgressStatusFindingBike:
		return "Finding a bike"
	case ProgressStatusBooked:
		return "Booked"
	case ProgressStatusCanceled:
		return "Canceled"
	}
	return ""
}

var (
	ProgressStatusPending     BookingProgressStatus = 1
	ProgressStatusFindingBike BookingProgressStatus = 2
	ProgressStatusBooked      BookingProgressStatus = 11
	ProgressStatusCanceled    BookingProgressStatus = 102
)

type BookingResult struct {
	ID        int64
	BikeID    string
	StationID string
	BookedAt  time.Time
}

func NewService() Service {
	return nil
}

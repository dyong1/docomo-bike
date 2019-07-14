package bikebooking

import "time"

type Service interface {
	BeginBooking(bookerID string, stationID string) (*BookingDetail, error)
	CancelBooking(bookingDetailID int64) (*BookingDetail, error)
	CompleteBooking(bookingDetailID int64) (*BookingDetail, error)
	GetCurrentBookingOfBooker(bookerID string) (*BookingDetail, error)
}

type BookingDetail struct {
	ID             int64
	StationID      string
	ProgressStatus BookingProgressStatus
	BookingResult  *BookingResult
	StartedAt      time.Time
	CanceledAt     *time.Time
}
type BookingProgressStatus int

func (s BookingProgressStatus) String() string {
	switch s {
	case ProgressStatusPending:
		return "Pending"
	case ProgressStatusFindingBike:
		return "Finding a bike"
	case ProgressStatusBooked:
		return "Booked"
	case ProgressStatusCheckedOutAlready:
		return "Checked out already"
	case ProgressStatusCanceled:
		return "Canceled"
	}
	return ""
}

var (
	ProgressStatusPending           BookingProgressStatus = 1
	ProgressStatusFindingBike       BookingProgressStatus = 2
	ProgressStatusBooked            BookingProgressStatus = 11
	ProgressStatusCheckedOutAlready BookingProgressStatus = 101
	ProgressStatusCanceled          BookingProgressStatus = 102
)

type BookingResult struct {
	BikeID   string
	BookedAt time.Time
}

func NewService() Service {
	return nil
}

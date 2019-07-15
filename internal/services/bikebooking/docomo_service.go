package bikebooking

import (
	"docomo-bike/internal/storage"
	"time"

	"gopkg.in/guregu/null.v3"

	"github.com/pkg/errors"
)

type DocomoService struct {
	BookingDetailRepo storage.BookingDetailRepository
	BookingResultRepo storage.BookingResultRepository
	BookingQueue      storage.BookingQueue
}

func (serv *DocomoService) BeginBooking(bookerID string, stationID string) (*BookingDetail, error) {
	current, err := serv.GetCurrentBookingOfBooker(bookerID)
	if err != nil {
		return nil, err
	}
	if current != nil {
		return nil, errors.Errorf("Unable to begin a new booking, existing booking was found [bookingDetailID=%d] [stationID=%s] [bookerID=%s]", current.ID, current.StationID, current.BookerID)
	}

	rtx, err := serv.BookingDetailRepo.Tx()
	if err != nil {
		return nil, err
	}
	qtx, err := serv.BookingQueue.Tx()
	if err != nil {
		return nil, err
	}
	var detail *storage.BookingDetail
	err = storage.WithEveryTx([]storage.Tx{rtx, qtx}, func(tx storage.Tx) error {
		detail, err = serv.BookingDetailRepo.Add(rtx, &storage.BookingDetail{
			BookerID:       bookerID,
			StationID:      stationID,
			ProgressStatus: ProgressStatusPending.Int(),
		})
		if err != nil {
			return err
		}
		_, err = serv.BookingQueue.Add(qtx, &storage.BookingJob{
			BookerID:        bookerID,
			StationID:       stationID,
			BookingDetailID: detail.ID,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return storageBookingDetail(*detail).BookingDetail(), nil
}

func (serv *DocomoService) CancelBooking(bookingDetailID int64) (*BookingDetail, error) {
	tx, err := serv.BookingDetailRepo.Tx()
	if err != nil {
		return nil, err
	}
	var detail *storage.BookingDetail
	err = storage.WithEveryTx([]storage.Tx{tx}, func(tx storage.Tx) error {
		detail, err = serv.getBookingDetailFromStorage(bookingDetailID)
		if err != nil {
			return err
		}

		detail.ProgressStatus = ProgressStatusCanceled.Int()
		detail.CanceledAt = null.TimeFrom(time.Now())
		if _, err := serv.BookingDetailRepo.Update(tx, detail); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return storageBookingDetail(*detail).BookingDetail(), nil
}

func (serv *DocomoService) CompleteBooking(bookingDetailID int64, bikeID string) (*BookingDetail, error) {
	dtx, err := serv.BookingDetailRepo.Tx()
	if err != nil {
		return nil, err
	}
	rtx, err := serv.BookingResultRepo.Tx()
	if err != nil {
		return nil, err
	}

	var detail *storage.BookingDetail
	var result *storage.BookingResult
	err = storage.WithEveryTx([]storage.Tx{dtx, rtx}, func(tx storage.Tx) error {
		detail, err = serv.getBookingDetailFromStorage(bookingDetailID)
		if err != nil {
			return err
		}

		result, err = serv.BookingResultRepo.Add(rtx, &storage.BookingResult{
			BikeID: bikeID,
		})
		if err != nil {
			return err
		}
		detail.ProgressStatus = ProgressStatusBooked.Int()
		detail.CanceledAt = null.TimeFrom(time.Now())
		detail.BookingResultID = null.IntFrom(result.ID)
		if _, err := serv.BookingDetailRepo.Update(dtx, detail); err != nil {
			return err
		}
		return nil
	})
	completed := storageBookingDetail(*detail).BookingDetail()
	completed.BookingResult = storageBookingResult(*result).BookingResult()
	return completed, nil
}

func (serv *DocomoService) GetCurrentBookingOfBooker(bookerID string) (*BookingDetail, error) {
	dd, err := serv.BookingDetailRepo.GetAllByBookerID(bookerID)
	if err != nil {
		return nil, err
	}
	for _, d := range dd {
		if d.ProgressStatus == ProgressStatusPending.Int() {
			return storageBookingDetail(*d).BookingDetail(), nil
		}
		if d.ProgressStatus == ProgressStatusFindingBike.Int() {
			return storageBookingDetail(*d).BookingDetail(), nil
		}
	}
	return nil, nil
}

func (serv *DocomoService) getBookingDetailFromStorage(id int64) (*storage.BookingDetail, error) {
	found, err := serv.BookingDetailRepo.Get(id)
	if err != nil {
		return nil, err
	}
	if found == nil {
		return nil, errors.Errorf("No such booking detail [bookingDetailID=%d]", id)
	}
	return found, err
}

type storageBookingDetail storage.BookingDetail

func (s storageBookingDetail) BookingDetail() *BookingDetail {
	return &BookingDetail{
		ID:             s.ID,
		BookerID:       s.BookerID,
		StationID:      s.StationID,
		ProgressStatus: BookingProgressStatus(s.ProgressStatus),
		BookingResult:  nil,
		StartedAt:      s.CreatedAt,
		CanceledAt:     s.CanceledAt.Ptr(),
	}
}

type storageBookingResult storage.BookingResult

func (s storageBookingResult) BookingResult() *BookingResult {
	return &BookingResult{
		ID:        s.ID,
		BikeID:    s.BikeID,
		StationID: s.StationID,
		BookedAt:  s.CreatedAt,
	}
}

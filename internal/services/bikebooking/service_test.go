package bikebooking

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBooking(t *testing.T) {
	serv := NewService()

	beginAt := time.Now()
	d, err := serv.BeginBooking("bookerID", "stationID")
	assert.NoError(t, err)
	assert.NotNil(t, d)
	verifyBeginBooking(t, beginAt, d, err)

	current, err := serv.GetCurrentBookingOfBooker("bookerID")
	assert.NoError(t, err)
	assert.NotNil(t, current)

	completed, err := serv.CompleteBooking(current.ID)
	assert.NoError(t, err)
	assert.NotNil(t, completed)
	assert.Equal(t, ProgressStatusBooked, completed.ProgressStatus)
	assert.NotNil(t, completed.BookingResult)
	assert.NotNil(t, completed.BookingResult.BookedAt)
}
func verifyBeginBooking(t *testing.T, beginAt time.Time, d *BookingDetail, err error) {
	assert.Nil(t, d.CanceledAt)
	assert.Nil(t, d.BookingResult)
	assert.Equal(t, ProgressStatusPending, d.ProgressStatus)
	assert.True(t, d.StartedAt.After(beginAt))
	sleepFinished := make(chan bool)
	go func() {
		time.Sleep(time.Millisecond)
		sleepFinished <- true
	}()
	<-sleepFinished
	assert.True(t, d.StartedAt.Before(time.Now()))
}

func TestCanceling(t *testing.T) {
	serv := NewService()
	d, err := serv.BeginBooking("stationID", "bookerID")
	assert.NoError(t, err)
	assert.NotNil(t, d)

	current, err := serv.GetCurrentBookingOfBooker("bookerID")
	assert.NoError(t, err)
	assert.NotNil(t, current)

	canceled, err := serv.CancelBooking(current.ID)
	assert.NoError(t, err)
	assert.NotNil(t, canceled)
	assert.Equal(t, ProgressStatusCanceled, canceled.ProgressStatus)
}

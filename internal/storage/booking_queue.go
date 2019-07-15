package storage

type BookingJob struct {
	BookerID        string
	StationID       string
	BookingDetailID int64
}

type BookingQueue interface {
	Tx() (Tx, error)
	Add(tx Tx, job *BookingJob) (*BookingJob, error)
}

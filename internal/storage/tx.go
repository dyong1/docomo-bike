package storage

type Tx interface {
	Rollback() error
	Commit() error
}

func CommitAll(tt ...Tx) error {
	for _, t := range tt {
		if err := t.Commit(); err != nil {
			return err
		}
	}
	return nil
}

func WithEveryTx(tt []Tx, txFunc TxFunc) (err error) {
	tx := everyTx{tt: tt}
	defer func() {
		err = tx.Rollback()
	}()
	if err := txFunc(tx); err != nil {
		return err
	}
	return tx.Commit()
}

type TxFunc func(tx Tx) error

type everyTx struct {
	tt []Tx
}

func (e everyTx) Rollback() error {
	var err error
	for _, t := range e.tt {
		if e := t.Rollback(); e != nil {
			err = e
		}
	}
	return err
}

func (e everyTx) Commit() error {
	return CommitAll(e.tt...)
}

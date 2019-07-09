package common

import (
	"database/sql"
)

type ICloser interface {
	Close() error
}

type Transaction struct {
	Tx     *sql.Tx
	Closer ICloser
}

func (t *Transaction) Commit() error {
	defer func() {
		t.Closer.Close()
	}()

	err := t.Tx.Commit()

	return err
}

func (t *Transaction) Rollback() error {
	defer func() {
		t.Closer.Close()
	}()

	err := t.Tx.Rollback()

	return err
}

func (t *Transaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	rs, err := t.Tx.Exec(query, args...)
	return rs, err
}

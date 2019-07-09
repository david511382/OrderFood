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
	err := t.Tx.Commit()
	if err != nil {
		return err
	}

	return t.Closer.Close()
}

func (t *Transaction) Rollback() error {
	err := t.Tx.Rollback()
	if err != nil {
		return err
	}

	return t.Closer.Close()
}

func (t *Transaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	rs, err := t.Tx.Exec(query, args...)
	return rs, err
}

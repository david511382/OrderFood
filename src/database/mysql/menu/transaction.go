package menu

import "database/sql"

func (d *MenuDb) Begin() (*sql.Tx, error) {
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	return tx, err
}

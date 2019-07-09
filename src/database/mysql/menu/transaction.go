package menu

import (
	"orderfood/src/database/common"
	mysqlCommon "orderfood/src/database/mysql/common"
)

func (d *MenuDb) Begin() (common.ITransaction, error) {
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()

	t := &mysqlCommon.Transaction{
		Tx:     tx,
		Closer: db,
	}
	return t, err
}

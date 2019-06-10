package menu

import (
	"orderfood/src/database/common"
	"orderfood/src/database/models"

	"github.com/jmoiron/sqlx"
)

func (d *MenuDb) GetItem(item *models.Item) ([]*models.Item, error) {
	condictionCols := make([]string, 0)
	if item != nil {
		condictionCols = itemCondiction(item)
	}

	sqlStr := common.ItemDt.SelectSQL(nil, condictionCols)

	args := make([]interface{}, 0)
	var err error
	if item != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, item)
		if err != nil {
			return nil, err
		}
	}

	items := make([]*models.Item, 0)
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Select(&items, sqlStr, args...)

	return items, err
}

func (d *MenuDb) AddItem(item *models.Item) error {
	if item == nil {
		return common.DbDataError
	}
	sqlStr := common.ItemDt.InsertSQL([]string{"id", "name", "price", "group_id"})

	sqlStr, args, err := sqlx.Named(sqlStr, item)
	if err != nil {
		return err
	}

	db, err := d.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	dbRes, err := db.Exec(sqlStr, args...)
	if err != nil {
		return err
	}

	id, err := dbRes.LastInsertId()
	if err != nil {
		return err
	}

	item.ID = int32(id)
	return nil
}

func (d *MenuDb) UpdateItem(item *models.Item) (int64, error) {
	if item == nil {
		return 0, common.DbDataError
	}

	cols := make([]string, 0)
	if item.GetName() != "" {
		cols = append(cols, "name")
	}
	if item.GetPrice() != 0 {
		cols = append(cols, "price")
	}
	if item.GetGroup_ID() != 0 {
		cols = append(cols, "group_id")
	}

	sqlStr := common.ItemDt.UpdateSQL(cols)

	args := make([]interface{}, 0)
	var err error
	if item != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, item)
		if err != nil {
			return 0, err
		}
	}

	db, err := d.Connect()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	r, err := db.Exec(sqlStr, args...)
	if err != nil {
		return 0, err
	}

	count, err := r.RowsAffected()
	return count, err
}

func (d *MenuDb) DeleteItem(item *models.Item) (int64, error) {
	if item == nil {
		return 0, common.DbDataError
	}

	condictionCols := itemCondiction(item)
	sqlStr := common.ItemDt.DeleteSQL(condictionCols)

	args := make([]interface{}, 0)
	var err error
	if item != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, item)
		if err != nil {
			return 0, err
		}
	}

	db, err := d.Connect()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	r, err := db.Exec(sqlStr, args...)
	if err != nil {
		return 0, err
	}

	count, err := r.RowsAffected()
	return count, err
}

func itemCondiction(item *models.Item) []string {
	condictionCols := make([]string, 0)
	if item.GetID() != 0 {
		condictionCols = append(condictionCols, "id")
	}
	if item.GetName() != "" {
		condictionCols = append(condictionCols, "name")
	}
	if item.GetPrice() != 0 {
		condictionCols = append(condictionCols, "price")
	}
	if item.GetGroup_ID() != 0 {
		condictionCols = append(condictionCols, "group_id")
	}
	return condictionCols
}

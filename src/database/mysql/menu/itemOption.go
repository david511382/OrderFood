package menu

import (
	"orderfood/src/database/common"
	"orderfood/src/database/models"

	"github.com/jmoiron/sqlx"
)

func (d *MenuDb) GetItemOption(itemOption *models.ItemOption) ([]*models.ItemOption, error) {
	condictionCols := make([]string, 0)
	if itemOption != nil {
		condictionCols = itemOptionCondiction(itemOption)
	}

	sqlStr := common.ItemOptionDt.SelectSQL(nil, condictionCols)

	args := make([]interface{}, 0)
	var err error
	if itemOption != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, itemOption)
		if err != nil {
			return nil, err
		}
	}

	itemOptions := make([]*models.ItemOption, 0)
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Select(&itemOptions, sqlStr, args...)

	return itemOptions, err
}

func (d *MenuDb) AddItemOption(itemOption *models.ItemOption) error {
	if itemOption == nil {
		return common.DbDataError
	}
	sqlStr := common.ItemOptionDt.InsertSQL([]string{"id", "option_id", "item_id"})

	sqlStr, args, err := sqlx.Named(sqlStr, itemOption)
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

	itemOption.ID = int32(id)
	return nil
}

func (d *MenuDb) UpdateItemOption(itemOption *models.ItemOption) (int64, error) {
	if itemOption == nil {
		return 0, common.DbDataError
	}

	cols := make([]string, 0)
	if itemOption.GetOption_ID() != 0 {
		cols = append(cols, "option_id")
	}
	if itemOption.GetItem_ID() != 0 {
		cols = append(cols, "item_id")
	}

	sqlStr := common.ItemOptionDt.UpdateSQL(cols)

	args := make([]interface{}, 0)
	var err error
	if itemOption != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, itemOption)
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

func (d *MenuDb) DeleteItemOption(itemOption *models.ItemOption) (int64, error) {
	if itemOption == nil {
		return 0, common.DbDataError
	}

	condictionCols := itemOptionCondiction(itemOption)
	sqlStr := common.ItemOptionDt.DeleteSQL(condictionCols)

	args := make([]interface{}, 0)
	var err error
	if itemOption != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, itemOption)
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

func itemOptionCondiction(itemOption *models.ItemOption) []string {
	condictionCols := make([]string, 0)
	if itemOption.GetID() != 0 {
		condictionCols = append(condictionCols, "id")
	}
	if itemOption.GetOption_ID() != 0 {
		condictionCols = append(condictionCols, "option_id")
	}
	if itemOption.GetItem_ID() != 0 {
		condictionCols = append(condictionCols, "item_id")
	}
	return condictionCols
}

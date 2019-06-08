package menu

import (
	"orderfood/src/database/common"
	"orderfood/src/database/models"

	"github.com/jmoiron/sqlx"
)

func (d *MenuDb) GetSelection(selection *models.Selection) ([]*models.Selection, error) {
	condictionCols := make([]string, 0)
	if selection != nil {
		condictionCols = selectionCondiction(selection)
	}

	sqlStr := common.SelectionDt.SelectSQL(nil, condictionCols)

	args := make([]interface{}, 0)
	var err error
	if selection != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, selection)
		if err != nil {
			return nil, err
		}
	}

	selections := make([]*models.Selection, 0)
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Select(&selections, sqlStr, args...)

	return selections, err
}

func (d *MenuDb) AddSelection(selection *models.Selection) error {
	if selection == nil {
		return common.DbDataError
	}
	sqlStr := common.SelectionDt.InsertSQL([]string{"id", "name"})

	sqlStr, args, err := sqlx.Named(sqlStr, selection)
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

	selection.ID = int32(id)
	return nil
}

func (d *MenuDb) UpdateSelection(selection *models.Selection) (int64, error) {
	if selection == nil {
		return 0, common.DbDataError
	}

	cols := make([]string, 0)
	if selection.GetName() != "" {
		cols = append(cols, "name")
	}

	sqlStr := common.SelectionDt.UpdateSQL(cols)

	args := make([]interface{}, 0)
	var err error
	if selection != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, selection)
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

func (d *MenuDb) DeleteSelection(selection *models.Selection) (int64, error) {
	if selection == nil {
		return 0, common.DbDataError
	}

	condictionCols := selectionCondiction(selection)
	sqlStr := common.SelectionDt.DeleteSQL(condictionCols)

	args := make([]interface{}, 0)
	var err error
	if selection != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, selection)
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

func selectionCondiction(selection *models.Selection) []string {
	condictionCols := make([]string, 0)
	if selection.GetID() != 0 {
		condictionCols = append(condictionCols, "id")
	}
	if selection.GetName() != "" {
		condictionCols = append(condictionCols, "name")
	}
	return condictionCols
}

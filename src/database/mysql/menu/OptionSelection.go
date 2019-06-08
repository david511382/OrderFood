package menu

import (
	"orderfood/src/database/common"
	"orderfood/src/database/models"

	"github.com/jmoiron/sqlx"
)

func (d *MenuDb) GetOptionSelection(optionSelection *models.OptionSelection) ([]*models.OptionSelection, error) {
	condictionCols := make([]string, 0)
	if optionSelection != nil {
		condictionCols = optionSelectionCondiction(optionSelection)
	}

	sqlStr := common.OptionSelectionDt.SelectSQL(nil, condictionCols)

	args := make([]interface{}, 0)
	var err error
	if optionSelection != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, optionSelection)
		if err != nil {
			return nil, err
		}
	}

	optionSelections := make([]*models.OptionSelection, 0)
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Select(&optionSelections, sqlStr, args...)

	return optionSelections, err
}

func (d *MenuDb) AddOptionSelection(optionSelection *models.OptionSelection) error {
	if optionSelection == nil {
		return common.DbDataError
	}
	sqlStr := common.OptionSelectionDt.InsertSQL([]string{"id", "option_id", "price", "selection_id"})

	sqlStr, args, err := sqlx.Named(sqlStr, optionSelection)
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

	optionSelection.ID = int32(id)
	return nil
}

func (d *MenuDb) UpdateOptionSelection(optionSelection *models.OptionSelection) (int64, error) {
	if optionSelection == nil {
		return 0, common.DbDataError
	}

	cols := make([]string, 0)
	if optionSelection.GetOption_ID() != 0 {
		cols = append(cols, "option_id")
	}
	if optionSelection.GetPrice() != 0 {
		cols = append(cols, "price")
	}
	if optionSelection.GetSelection_ID() != 0 {
		cols = append(cols, "selection_id")
	}

	sqlStr := common.OptionSelectionDt.UpdateSQL(cols)

	args := make([]interface{}, 0)
	var err error
	if optionSelection != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, optionSelection)
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

func (d *MenuDb) DeleteOptionSelection(optionSelection *models.OptionSelection) (int64, error) {
	if optionSelection == nil {
		return 0, common.DbDataError
	}

	condictionCols := optionSelectionCondiction(optionSelection)
	sqlStr := common.OptionSelectionDt.DeleteSQL(condictionCols)

	args := make([]interface{}, 0)
	var err error
	if optionSelection != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, optionSelection)
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

func optionSelectionCondiction(optionSelection *models.OptionSelection) []string {
	condictionCols := make([]string, 0)
	if optionSelection.GetID() != 0 {
		condictionCols = append(condictionCols, "id")
	}
	if optionSelection.GetOption_ID() != 0 {
		condictionCols = append(condictionCols, "option_id")
	}
	if optionSelection.GetPrice() != 0 {
		condictionCols = append(condictionCols, "price")
	}
	if optionSelection.GetSelection_ID() != 0 {
		condictionCols = append(condictionCols, "selection_id")
	}
	return condictionCols
}

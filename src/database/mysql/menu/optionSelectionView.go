package menu

import (
	"orderfood/src/database/common"
	"orderfood/src/database/models"

	"github.com/jmoiron/sqlx"
)

func (d *MenuDb) GetOptionSelectionView(optionSelectionView *models.OptionSelectionView) ([]*models.OptionSelectionView, error) {
	condictionCols := make([]string, 0)
	if optionSelectionView != nil {
		if optionSelectionView.GetOption_ID() != 0 {
			condictionCols = append(condictionCols, "option_id")
		}
		if optionSelectionView.GetSelect_Num() > -1 {
			condictionCols = append(condictionCols, "select_num")
		}
		if optionSelectionView.GetSelection_ID() != nil {
			condictionCols = append(condictionCols, "selection_id")
		}
		if optionSelectionView.GetName() != nil {
			condictionCols = append(condictionCols, "name")
		}
		if optionSelectionView.GetPrice() != nil {
			condictionCols = append(condictionCols, "price")
		}
	}

	sqlStr := common.OptionSelectionViewDt.SelectSQL(nil, condictionCols)

	args := make([]interface{}, 0)
	var err error
	if optionSelectionView != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, optionSelectionView)
		if err != nil {
			return nil, err
		}
	}

	optionSelectionViews := make([]*models.OptionSelectionView, 0)
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Select(&optionSelectionViews, sqlStr, args...)

	return optionSelectionViews, err
}

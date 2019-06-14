package menu

import (
	"orderfood/src/database/common"
	"orderfood/src/database/models"

	"github.com/jmoiron/sqlx"
)

func (d *MenuDb) GetItemOptionView(itemOptionView *models.ItemOptionView) ([]*models.ItemOptionView, error) {
	condictionCols := make([]string, 0)
	if itemOptionView != nil {
		if itemOptionView.GetShop_ID() != 0 {
			condictionCols = append(condictionCols, "shop_id")
		}
		if itemOptionView.GetOption_ID() != nil {
			condictionCols = append(condictionCols, "option_id")
		}
		if itemOptionView.GetItem_ID() != 0 {
			condictionCols = append(condictionCols, "item_id")
		}
		if itemOptionView.GetName() != "" {
			condictionCols = append(condictionCols, "name")
		}
		if itemOptionView.GetPrice() > -1 {
			condictionCols = append(condictionCols, "price")
		}
	}

	sqlStr := common.ItemOptionViewDt.SelectSQL(nil, condictionCols)

	args := make([]interface{}, 0)
	var err error
	if itemOptionView != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, itemOptionView)
		if err != nil {
			return nil, err
		}
	}

	itemOptionViews := make([]*models.ItemOptionView, 0)
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Select(&itemOptionViews, sqlStr, args...)

	return itemOptionViews, err
}

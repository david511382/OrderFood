package menu

import (
	"orderfood/src/database/common"
	"orderfood/src/database/models"

	"github.com/jmoiron/sqlx"
)

func (d *MenuDb) GetShop(shop *models.Shop) ([]*models.Shop, error) {
	condictionCols := make([]string, 0)
	if shop != nil {
		condictionCols = shopCondiction(shop)
	}

	sqlStr := common.ShopDt.SelectSQL(nil, condictionCols)

	args := make([]interface{}, 0)
	var err error
	if shop != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, shop)
		if err != nil {
			return nil, err
		}
	}

	shops := make([]*models.Shop, 0)
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Select(&shops, sqlStr, args...)

	return shops, err
}

func (d *MenuDb) AddShop(shop *models.Shop) error {
	if shop == nil {
		return common.DbDataError
	}
	sqlStr := common.ShopDt.InsertSQL([]string{"id", "name"})

	sqlStr, args, err := sqlx.Named(sqlStr, shop)
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

	shop.ID = int(id)
	return nil
}

func (d *MenuDb) UpdateShop(shop *models.Shop) (int64, error) {
	if shop == nil {
		return 0, common.DbDataError
	}

	cols := make([]string, 0)
	if shop.GetName() != "" {
		cols = append(cols, "name")
	}

	sqlStr := common.ShopDt.UpdateSQL(cols)

	args := make([]interface{}, 0)
	var err error
	if shop != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, shop)
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

func (d *MenuDb) DeleteShop(shop *models.Shop) (int64, error) {
	if shop == nil {
		return 0, common.DbDataError
	}

	condictionCols := shopCondiction(shop)
	sqlStr := common.ShopDt.DeleteSQL(condictionCols)

	args := make([]interface{}, 0)
	var err error
	if shop != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, shop)
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

func shopCondiction(shop *models.Shop) []string {
	condictionCols := make([]string, 0)
	if shop.GetID() != 0 {
		condictionCols = append(condictionCols, "id")
	}
	if shop.GetName() != "" {
		condictionCols = append(condictionCols, "name")
	}
	return condictionCols
}

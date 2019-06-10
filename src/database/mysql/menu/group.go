package menu

import (
	"orderfood/src/database/common"
	"orderfood/src/database/models"

	"github.com/jmoiron/sqlx"
)

func (d *MenuDb) GetGroup(group *models.Group) ([]*models.Group, error) {
	condictionCols := make([]string, 0)
	if group != nil {
		condictionCols = groupCondiction(group)
	}

	sqlStr := common.GroupDt.SelectSQL(nil, condictionCols)

	args := make([]interface{}, 0)
	var err error
	if group != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, group)
		if err != nil {
			return nil, err
		}
	}

	groups := make([]*models.Group, 0)
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Select(&groups, sqlStr, args...)

	return groups, err
}

func (d *MenuDb) AddGroup(group *models.Group) error {
	if group == nil {
		return common.DbDataError
	}
	sqlStr := common.GroupDt.InsertSQL([]string{"id", "shop_id", "least_select_num"})

	sqlStr, args, err := sqlx.Named(sqlStr, group)
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

	group.ID = int32(id)
	return nil
}

func (d *MenuDb) UpdateGroup(group *models.Group) (int64, error) {
	if group == nil {
		return 0, common.DbDataError
	}

	cols := make([]string, 0)
	if group.GetShop_ID() != 0 {
		cols = append(cols, "shop_id")
	}
	if group.GetLeast_Select_Num() != -1 {
		cols = append(cols, "least_select_num")
	}

	sqlStr := common.GroupDt.UpdateSQL(cols)

	args := make([]interface{}, 0)
	var err error
	if group != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, group)
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

func (d *MenuDb) DeleteGroup(group *models.Group) (int64, error) {
	if group == nil {
		return 0, common.DbDataError
	}

	condictionCols := groupCondiction(group)
	sqlStr := common.GroupDt.DeleteSQL(condictionCols)

	args := make([]interface{}, 0)
	var err error
	if group != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, group)
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

func groupCondiction(group *models.Group) []string {
	condictionCols := make([]string, 0)
	if group.GetID() != 0 {
		condictionCols = append(condictionCols, "id")
	}
	if group.GetShop_ID() != 0 {
		condictionCols = append(condictionCols, "shop_id")
	}
	if group.GetLeast_Select_Num() > -1 {
		condictionCols = append(condictionCols, "least_select_num")
	}
	return condictionCols
}

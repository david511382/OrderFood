package menu

import (
	"orderfood/src/database/common"
	"orderfood/src/database/models"

	"github.com/jmoiron/sqlx"
)

func (d *MenuDb) GetOption(option *models.Option) ([]*models.Option, error) {
	condictionCols := make([]string, 0)
	if option != nil {
		condictionCols = optionCondiction(option)
	}

	sqlStr := common.OptionDt.SelectSQL(nil, condictionCols)

	args := make([]interface{}, 0)
	var err error
	if option != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, option)
		if err != nil {
			return nil, err
		}
	}

	options := make([]*models.Option, 0)
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Select(&options, sqlStr, args...)

	return options, err
}

func (d *MenuDb) AddOption(option *models.Option, tx common.IExecer) error {
	if option == nil {
		return common.DbDataError
	}
	sqlStr := common.OptionDt.InsertSQL([]string{"id", "select_num"})

	sqlStr, args, err := sqlx.Named(sqlStr, option)
	if err != nil {
		return err
	}

	var execer sqlx.Execer
	if tx != nil {
		execer = tx
	} else {
		db, err := d.Connect()
		if err != nil {
			return err
		}
		defer db.Close()

		execer = db.DB
	}

	dbRes, err := execer.Exec(sqlStr, args...)
	if err != nil {
		return err
	}

	id, err := dbRes.LastInsertId()
	if err != nil {
		return err
	}

	option.ID = int(id)
	return nil
}

func (d *MenuDb) UpdateOption(option *models.Option, tx common.IExecer) (int64, error) {
	if option == nil {
		return 0, common.DbDataError
	}

	cols := make([]string, 0)
	if option.GetSelect_Num() != -1 {
		cols = append(cols, "select_num")
	}

	sqlStr := common.OptionDt.UpdateSQL(cols)

	args := make([]interface{}, 0)
	var err error
	if option != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, option)
		if err != nil {
			return 0, err
		}
	}

	var execer sqlx.Execer
	if tx != nil {
		execer = tx
	} else {
		db, err := d.Connect()
		if err != nil {
			return 0, err
		}
		defer db.Close()

		execer = db.DB
	}

	r, err := execer.Exec(sqlStr, args...)
	if err != nil {
		return 0, err
	}

	count, err := r.RowsAffected()
	return count, err
}

func (d *MenuDb) DeleteOption(option *models.Option, tx common.IExecer) (int64, error) {
	if option == nil {
		return 0, common.DbDataError
	}

	condictionCols := optionCondiction(option)
	sqlStr := common.OptionDt.DeleteSQL(condictionCols)

	args := make([]interface{}, 0)
	var err error
	if option != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, option)
		if err != nil {
			return 0, err
		}
	}

	var execer sqlx.Execer
	if tx != nil {
		execer = tx
	} else {
		db, err := d.Connect()
		if err != nil {
			return 0, err
		}
		defer db.Close()

		execer = db.DB
	}
	r, err := execer.Exec(sqlStr, args...)
	if err != nil {
		return 0, err
	}

	count, err := r.RowsAffected()
	return count, err
}

func optionCondiction(option *models.Option) []string {
	condictionCols := make([]string, 0)
	if option.GetID() != 0 {
		condictionCols = append(condictionCols, "id")
	}
	if option.GetSelect_Num() > -1 {
		condictionCols = append(condictionCols, "select_num")
	}
	return condictionCols
}

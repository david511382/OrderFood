package member

import (
	"orderfood/src/database/common"
	"orderfood/src/database/models"

	"github.com/jmoiron/sqlx"
)

func (d *MemberDb) GetMember(member *models.Member) ([]models.Member, error) {
	condictionCols := make([]string, 0)
	if member != nil {
		if member.GetID() != 0 {
			condictionCols = append(condictionCols, "id")
		}
		if member.GetName() != "" {
			condictionCols = append(condictionCols, "name")
		}
		if member.GetUsername() != "" {
			condictionCols = append(condictionCols, "username")
		}
		if member.GetPassword() != "" {
			condictionCols = append(condictionCols, "password")
		}
	}

	sqlStr := common.MemberDt.SelectSQL(nil, condictionCols)

	args := make([]interface{}, 0)
	var err error
	if member != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, member)
		if err != nil {
			return nil, err
		}
	}

	members := make([]models.Member, 0)
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Select(&members, sqlStr, args...)

	return members, err
}

func (d *MemberDb) AddMember(member *models.Member) (*models.Member, error) {
	if member == nil {
		return nil, common.DbDataError
	}
	sqlStr := common.MemberDt.InsertSQL([]string{"id", "name", "username", "password"})

	sqlStr, args, err := sqlx.Named(sqlStr, member)
	if err != nil {
		return nil, err
	}

	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	dbRes, err := db.Exec(sqlStr, args...)
	if err != nil {
		return nil, err
	}

	id, err := dbRes.LastInsertId()
	if err != nil {
		return nil, err
	}

	member.ID = int32(id)
	return member, err
}

func (d *MemberDb) UpdateMember(member *models.Member) (*models.Member, error) {
	if member == nil {
		return nil, common.DbDataError
	}

	cols := []string{"name", "username", "password"}
	
	condictionCols := make([]string, 0)
	if member.GetID() != 0 {
		condictionCols = append(condictionCols, "id")
	}
	if member.GetName() != "" {
		condictionCols = append(condictionCols, "name")
	}
	if member.GetUsername() != "" {
		condictionCols = append(condictionCols, "username")
	}
	if member.GetPassword() != "" {
		condictionCols = append(condictionCols, "password")
	}

	sqlStr := common.MemberDt.UpdateSQL(cols, condictionCols)

	args := make([]interface{}, 0)
	var err error
	if member != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, member)
		if err != nil {
			return nil, err
		}
	}

	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	_, err = db.Exec(sqlStr, args...)

	return member, err
}

func (d *MemberDb) DeleteMember(member *models.Member) error {
	if member == nil {
		return common.DbDataError
	}

	condictionCols := make([]string, 0)
	if member.GetID() != 0 {
		condictionCols = append(condictionCols, "id")
	}
	if member.GetName() != "" {
		condictionCols = append(condictionCols, "name")
	}
	if member.GetUsername() != "" {
		condictionCols = append(condictionCols, "username")
	}
	if member.GetPassword() != "" {
		condictionCols = append(condictionCols, "password")
	}

	sqlStr := common.MemberDt.DeleteSQL(condictionCols)

	args := make([]interface{}, 0)
	var err error
	if member != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, member)
		if err != nil {
			return err
		}
	}

	db, err := d.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	r, err := db.Exec(sqlStr, args...)
	if err != nil {
		return err
	}
	if id, err := r.RowsAffected(); id != 1 {
		return common.DbDataError
	} else if err != nil {
		return err
	}

	return nil
}

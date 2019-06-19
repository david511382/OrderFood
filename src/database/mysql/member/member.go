package member

import (
	"orderfood/src/database/common"
	"orderfood/src/database/models"

	"github.com/jmoiron/sqlx"
)

func (d *MemberDb) GetMember(member *models.Member) ([]*models.Member, error) {
	condictionCols := make([]string, 0)
	if member != nil {
		condictionCols = memberCondiction(member)
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

	members := make([]*models.Member, 0)
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Select(&members, sqlStr, args...)

	return members, err
}

func (d *MemberDb) AddMember(member *models.Member) error {
	if member == nil {
		return common.DbDataError
	}
	sqlStr := common.MemberDt.InsertSQL([]string{"id", "name", "username", "password"})

	sqlStr, args, err := sqlx.Named(sqlStr, member)
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

	member.ID = int32(id)
	return nil
}

func (d *MemberDb) UpdateMember(member *models.Member) (int64, error) {
	if member == nil {
		return 0, common.DbDataError
	}

	cols := make([]string, 0)
	if member.GetName() != "" {
		cols = append(cols, "name")
	}
	if member.GetUsername() != "" {
		cols = append(cols, "username")
	}
	if member.GetPassword() != "" {
		cols = append(cols, "password")
	}

	sqlStr := common.MemberDt.UpdateSQL(cols)

	args := make([]interface{}, 0)
	var err error
	if member != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, member)
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

func (d *MemberDb) DeleteMember(member *models.Member) (int64, error) {
	if member == nil {
		return 0, common.DbDataError
	}

	condictionCols := memberCondiction(member)
	sqlStr := common.MemberDt.DeleteSQL(condictionCols)

	args := make([]interface{}, 0)
	var err error
	if member != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, member)
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

func memberCondiction(member *models.Member) []string {
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
	return condictionCols
}

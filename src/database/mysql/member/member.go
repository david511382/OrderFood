package member

import (
	"fmt"
	"orderfood/src/database/common"
	"orderfood/src/database/models"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

func (d *MemberDb) GetMember(member *models.Member) ([]models.Member, error) {
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStr := `
		SELECT
			*
		FROM
			%s
		%%s`
	sqlStr = fmt.Sprintf(sqlStr, common.MemberDt.Name())

	wheres := make([]string, 0)
	if member != nil {
		if member.GetID() != 0 {
			wheres = append(wheres, strconv.Itoa(int(member.GetID())))
		}
		if member.GetName() != "" {
			wheres = append(wheres, member.GetName())
		}
		if member.GetUsername() != "" {
			wheres = append(wheres, member.GetUsername())
		}
		if member.GetPassword() != "" {
			wheres = append(wheres, member.GetPassword())
		}
	}

	whereStr := strings.Join(wheres, " AND ")
	sqlStr = fmt.Sprintf(sqlStr, whereStr)

	args := make([]interface{}, 0)
	if member != nil {
		sqlStr, args, err = sqlx.Named(sqlStr, member)
		if err != nil {
			return nil, err
		}
	}

	members := make([]models.Member, 0)
	err = db.Select(&members, sqlStr, args...)

	return members, err
}

func (d *MemberDb) AddMember(member *models.Member) (*models.Member, error) {
	if member == nil {
		return nil, common.DbDataError
	}

	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStr := `
		INSERT INTO
			%s
			(id, name, username, password)
		VALUES
			(:id, :name, :username, :password)`
	sqlStr = fmt.Sprintf(sqlStr, common.MemberDt.Name())

	sqlStr, args, err := sqlx.Named(sqlStr, member)
	if err != nil {
		return nil, err
	}

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

func (db *MemberDb) UpdateMember(member *models.Member) (*models.Member, error) {
	return nil, nil
}

func (db *MemberDb) DeleteMember(member *models.Member) error {
	return nil
}

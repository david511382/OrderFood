package member

import (
	"fmt"
	"orderfood/src/database/models"

	"orderfood/src/database/common"
)

func (d *MemberDb) GetMembers() ([]models.Member, error) {
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStr := `
		SELECT
			*
		FROM
			%s`
	sqlStr = fmt.Sprintf(sqlStr, common.MemberDt.Name())

	members := make([]models.Member, 0)
	err = db.Select(&members, sqlStr)

	return members, err
}

func (d *MemberDb) AddMembers(*models.Member) error {
	return nil
}

func (db *MemberDb) UpdateMembers(member models.Member) error {
	return nil
}

package member

import (
	"orderfood/src/database/models"
)

func (d *MemberDb) GetMembers() ([]models.Member, error) {
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	members := make([]models.Member, 0)
	err = db.Select(&members, "SELECT * FROM user_info")

	return members, err
}

func (d *MemberDb) AddMembers(*models.Member) error {
	return nil
}

func (db *MemberDb) UpdateMembers(member models.Member) error {
	return nil
}

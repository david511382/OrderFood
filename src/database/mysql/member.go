package mysql

import (
	"orderfood/src/database/models"
)

func (d *mysqlDb) GetMembers() ([]models.Member, error) {
	db, err := d.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	members := make([]models.Member, 0)
	err = db.Select(&members, "SELECT * FROM user_info")

	return members, err
}

func (d *mysqlDb) AddMembers(models.Member) error {
	return nil
}

func (db *mysqlDb) UpdateMembers(member models.Member) error {
	return nil
}

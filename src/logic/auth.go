package logic

import (
	"errors"
	"orderfood/src/database"
	"orderfood/src/database/models"
)

var (
	ExisitErr = errors.New("exist")
)

func Register(member *models.Member) error {
	dbMember := GetMember(member.GetUsername())
	if dbMember != nil {
		*member = *dbMember
		return ExisitErr
	}

	_, err := database.Db.Member().AddMember(member)

	return err
}

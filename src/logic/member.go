package logic

import (
	"orderfood/src/database"
	"orderfood/src/database/models"
)

var (
	Members = make([]models.Member, 0)
)

func LoadMembers() {
	members, err := database.Db.Member().GetMember(nil)
	if err != nil {
		panic(err)
	}

	Members = members
}

func GetMember(username string) *models.Member {
	for _, v := range Members {
		if v.GetUsername() == username {
			return &v
		}
	}
	return nil
}

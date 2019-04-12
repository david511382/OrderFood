package logic

import (
	"orderfood/src/database"
	"orderfood/src/handler/models/resp"
)

func GetMenu() ([]resp.MenuKind, error) {
	shop := GetView()

	_, err := database.Db.GetMenus(shop)
	if err != nil {
		return nil, err
	}

	resp := make([]resp.MenuKind, 0)
	return resp, nil
}

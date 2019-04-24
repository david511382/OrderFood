package logic

import (
	"orderfood/src/database"
	"orderfood/src/database/models"
	"orderfood/src/handler/models/resp"
)

func GetMenu() ([]resp.MenuKind, error) {
	shop := GetView()

	_, err := database.Db.Menu().GetMenus(shop)
	if err != nil {
		return nil, err
	}

	resp := make([]resp.MenuKind, 0)
	return resp, nil
}

func AddItem(name string) (*models.Item, error) {
	item := &models.Item{
		Name: name,
	}
	item, err := database.Db.Menu().AddItem(item)
	return item, err
}

func GetItem() ([]*models.Item, error) {
	items, err := database.Db.Menu().GetItems()
	return items, err
}

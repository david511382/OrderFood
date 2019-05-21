package logic

import (
	"orderfood/src/database"
	"orderfood/src/database/models"
	"orderfood/src/handler/models/resp"
)

func GetMenu(shop string) ([]resp.MenuKind, error) {
	_, err := database.Db.Menu().GetMenus(shop)
	if err != nil {
		return nil, err
	}

	resp := make([]resp.MenuKind, 0)
	return resp, nil
}

func AddShopItem(shopID int32, name string) (*models.Item, error) {
	item := &models.Item{
		Name: name,
	}
	item, err := database.Db.Menu().AddItem(item)
	if err != nil {
		return item, err
	}

	shopitem := &models.ShopItem{
		ShopID:     shopID,
		MenuItemID: item.GetID(),
	}
	shopitem, err = database.Db.Menu().AddShopItem(shopitem)
	if err != nil {
		return nil, err
	}

	return item, err
}

func GetShopItem(shopID int32) ([]*models.Item, error) {
	items, err := database.Db.Menu().GetItems(shopID)
	return items, err
}

func AddSize(name string) (*models.Size, error) {
	size := &models.Size{
		Name: name,
	}
	size, err := database.Db.Menu().AddSize(size)
	return size, err
}

func GetSize() ([]*models.Size, error) {
	sizes, err := database.Db.Menu().GetSizes()
	return sizes, err
}

func AddKind(name string) (*models.Kind, error) {
	kind := &models.Kind{
		Name: name,
	}
	kind, err := database.Db.Menu().AddKind(kind)
	return kind, err
}

func GetKind() ([]*models.Kind, error) {
	kinds, err := database.Db.Menu().GetKinds()
	return kinds, err
}

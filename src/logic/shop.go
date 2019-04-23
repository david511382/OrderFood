package logic

import (
	"orderfood/src/database"
	"orderfood/src/database/models"
)

func AddShop(name string) (*models.Shop, error) {
	shop := &models.Shop{
		Name: name,
	}
	shop, err := database.Db.AddShop(shop)
	if err != nil {
		return nil, err
	}

	return shop, nil
}

func GetShop() ([]*models.Shop, error) {
	shop, err := database.Db.GetShop()
	return shop, err
}
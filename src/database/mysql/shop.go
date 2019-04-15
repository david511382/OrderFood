package mysql

import (
	"orderfood/src/database/models"
)

func (d *mysqlDb) GetMenus(shop string) ([]models.MenuItem, error) {
	return nil, nil
}

func (d *mysqlDb) AddShop(shop *models.Shop) (*models.Shop, error) {
	return shop, nil
}

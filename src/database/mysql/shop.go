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

func (d *mysqlDb) GetShops() ([]*models.Shop, error) {
	return nil, nil
}

func (d *mysqlDb) AddItem(item *models.Item) (*models.Item, error) {
	return nil, nil
}

func (d *mysqlDb) GetItems(shopID int32) ([]*models.Item, error) {
	return nil, nil
}

func (d *mysqlDb) AddItemSize(itemSize *models.ItemSize) (*models.ItemSize, error) {
	return nil, nil
}

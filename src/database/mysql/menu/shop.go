package menu

import (
	"orderfood/src/database/models"
)

// Shop 。
func (d *MenuDb) GetShop(*models.Shop) ([]*models.Shop, error) {
	return nil, nil
}
func (d *MenuDb) AddShop(*models.Shop) error {
	return nil
}
func (d *MenuDb) DeleteShop(*models.Shop) (int64, error) {
	return 0, nil
}
func (d *MenuDb) UpdateShop(*models.Shop) (int64, error) {
	return 0, nil
}

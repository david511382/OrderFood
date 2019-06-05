package menu

import (
	"orderfood/src/database/models"
)

// ItemOption 。
func (d *MenuDb) GetItemOption(*models.ItemOption) ([]*models.ItemOption, error) {
	return nil, nil
}
func (d *MenuDb) AddItemOption(*models.ItemOption) error {
	return nil
}
func (d *MenuDb) DeleteItemOption(*models.ItemOption) (int64, error) {
	return 0, nil
}
func (d *MenuDb) UpdateItemOption(*models.ItemOption) (int64, error) {
	return 0, nil
}

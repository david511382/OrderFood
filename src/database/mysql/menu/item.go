package menu

import (
	"orderfood/src/database/models"
)

// Item 。
func (d *MenuDb) GetItem(*models.Item) ([]*models.Item, error) {
	return nil, nil
}
func (d *MenuDb) AddItem(*models.Item) error {
	return  nil
}
func (d *MenuDb) DeleteItem(*models.Item) error {
	return nil
}
func (d *MenuDb) UpdateItem(*models.Item) error {
	return  nil
}

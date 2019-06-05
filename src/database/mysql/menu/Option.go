package menu

import (
	"orderfood/src/database/models"
)

// Option 。
func (d *MenuDb) GetOption(*models.Option) ([]*models.Option, error) {
	return nil, nil
}
func (d *MenuDb) AddOption(*models.Option) error {
	return nil
}
func (d *MenuDb) DeleteOption(*models.Option) (int64, error) {
	return 0, nil
}
func (d *MenuDb) UpdateOption(*models.Option) (int64, error) {
	return 0, nil
}

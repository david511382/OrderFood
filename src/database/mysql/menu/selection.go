package menu

import (
	"orderfood/src/database/models"
)

// selection
func (d *MenuDb) GetSelection(*models.Selection) ([]*models.Selection, error) {
	return nil, nil
}
func (d *MenuDb) AddSelection(*models.Selection) error {
	return nil
}
func (d *MenuDb) DeleteSelection(*models.Selection) (int64, error) {
	return 0, nil
}
func (d *MenuDb) UpdateSelection(*models.Selection) (int64, error) {
	return 0, nil
}

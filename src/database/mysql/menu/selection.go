package menu

import (
	"orderfood/src/database/models"
)

// selection
func (d *MenuDb) GetSelection(*models.Selection) ([]*models.Selection, error) {
	return nil, nil
}
func (d *MenuDb) AddSelection(*models.Selection) (*models.Selection, error) {
	return nil, nil
}
func (d *MenuDb) DeleteSelection(*models.Selection) error {
	return nil
}
func (d *MenuDb) UpdateSelection(*models.Selection) (*models.Selection, error) {
	return nil, nil
}

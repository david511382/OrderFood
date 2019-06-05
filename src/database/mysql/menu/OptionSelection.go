package menu

import (
	"orderfood/src/database/models"
)

// OptionSelection 。
func (d *MenuDb) GetOptionSelection(*models.OptionSelection) ([]*models.OptionSelection, error) {
	return nil, nil
}
func (d *MenuDb) AddOptionSelection(*models.OptionSelection) error {
	return nil
}
func (d *MenuDb) DeleteOptionSelection(*models.OptionSelection) (int64, error) {
	return 0, nil
}
func (d *MenuDb) UpdateOptionSelection(*models.OptionSelection) (int64, error) {
	return 0, nil
}

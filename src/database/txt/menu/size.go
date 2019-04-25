package menu

import (
	"orderfood/src/database/models"
	"orderfood/src/database/txt/orm"
)

func (d *MenuDb) AddSize(size *models.Size) (*models.Size, error) {
	err := orm.SizeDT.Insert(size)
	
	return size, err
}

func (d *MenuDb) GetSizes() ([]*models.Size, error) {
	isizes, err := orm.SizeDT.Select(nil)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Size, 0)
	for _, v := range isizes {
		result = append(result, v.(*models.Size))
	}

	return result, nil
}

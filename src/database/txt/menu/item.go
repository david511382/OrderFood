package menu

import (
	"orderfood/src/database/models"
	"orderfood/src/database/txt/orm"
)

func (d *MenuDb) AddItem(item *models.Item) (*models.Item, error) {
	err := orm.ItemDT.Insert(item)
	
	return item, err
}

func (d *MenuDb) GetItems() ([]*models.Item, error) {
	iitems, err := orm.ItemDT.Select().Exec()
	if err != nil {
		return nil, err
	}

	result := make([]*models.Item, 0)
	for _, v := range iitems {
		result = append(result, v.(*models.Item))
	}

	return result, nil
}

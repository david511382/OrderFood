package menu

import (
	"orderfood/src/database/models"
	"orderfood/src/database/txt/orm"
)

func (d *MenuDb) AddKind(kind *models.Kind) (*models.Kind, error) {
	err := orm.KindDT.Insert(kind)

	return kind, err
}

func (d *MenuDb) GetKinds() ([]*models.Kind, error) {
	ikinds, err := orm.KindDT.Select(nil)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Kind, 0)
	for _, v := range ikinds {
		result = append(result, v.(*models.Kind))
	}

	return result, nil
}

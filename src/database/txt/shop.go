package txt

import (
	"orderfood/src/database/models"
)

func (d *txtDb) GetMenus(shopName string) ([]models.MenuItem, error) {
	// get shop id
	ishops, err := shopDT.Select().Where(
		func(model interface{}) bool {
			shop := model.(models.Shop)
			if shop.Name == shopName {
				return true
			}

			return false
		}).Exec()
	if err != nil {
		return nil, err
	}

	shops := make([]models.Shop, 0)
	for _, v := range ishops {
		shop := v.(models.Shop)
		shops = append(shops, shop)
	}

	// get item ids

	// make every item
	// get item size and price
	// get item kind and price
	return nil, nil
}

func (d *txtDb) AddShop(shop *models.Shop) (*models.Shop, error) {
	return shop, nil
}

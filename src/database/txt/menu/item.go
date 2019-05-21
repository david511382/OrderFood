package menu

import (
	"orderfood/src/database/models"
	"orderfood/src/database/txt/orm"
)

func (d *MenuDb) AddItem(item *models.Item) (*models.Item, error) {
	err := orm.ItemDT.Insert(item)

	return item, err
}

func (d *MenuDb) GetItems(shopID int32) ([]*models.Item, error) {
	iitemIDs, err := orm.ShopItemDT.Select(func(model interface{}) bool {
		item := model.(*models.ShopItem)
		if item.GetShopID() == shopID {
			return true
		}
		return false
	})
	if err != nil {
		return nil, err
	}

	iitems, err := orm.ItemDT.Select(func(model interface{}) bool {
		item := model.(*models.Item)

		for i, v := range iitemIDs {
			si := v.(*models.ShopItem)
			if item.GetID() == si.GetMenuItemID() {
				pre := iitemIDs[:i]
				aft := iitemIDs[i+1:]
				iitemIDs = append(iitemIDs[:0], pre...)
				iitemIDs = append(iitemIDs, aft...)

				return true
			}
		}

		return false
	})
	if err != nil {
		return nil, err
	}

	result := make([]*models.Item, 0)
	for _, v := range iitems {
		result = append(result, v.(*models.Item))
	}

	return result, nil
}

func (d *MenuDb) AddShopItem(item *models.ShopItem) (*models.ShopItem, error) {
	err := orm.ShopItemDT.Insert(item)

	return item, err
}

package txt

import (
	"orderfood/src/database/models"

	linq "github.com/ahmetb/go-linq"
	proto "github.com/golang/protobuf/proto"
)

func (d *txtDb) GetMenus(shopName string) ([]models.MenuItem, error) {
	// get shop id
	ishops, err := shopDT.Select().Where(
		func(model interface{}) bool {
			shop := model.(*models.Shop)
			if shop.Name == shopName {
				return true
			}

			return false
		}).Exec()
	if err != nil {
		return nil, err
	}

	if len(ishops) != 1 {
		return nil, dbDataError
	}
	shop := (ishops[0]).(*models.Shop)

	// get item ids
	ishopItems, err := shopItemDT.Select().Where(
		func(model interface{}) bool {
			shopItem := model.(*models.ShopItem)
			if shopItem.GetShopID() == shop.GetID() {
				return true
			}

			return false
		}).Exec()
	if err != nil {
		return nil, err
	}

	shopItemIDs := make([]int32, 0)
	for _, v := range ishopItems {
		shopItemIDs = append(shopItemIDs, v.(*models.ShopItem).GetMenuItemID())
	}

	// make every item
	iItems, err := itemDT.Select().Where(
		func(model interface{}) bool {
			item := model.(*models.Item)
			for i, id := range shopItemIDs {
				if item.GetID() == id {
					shopItemIDs = append(shopItemIDs[:0], shopItemIDs[:i]...)
					shopItemIDs = append(shopItemIDs, shopItemIDs[i+1:]...)
					return true
				}
			}

			return false
		}).Exec()
	if err != nil {
		return nil, err
	}

	result := make([]models.MenuItem, 0)
	for _, v := range iItems {
		result = append(result, models.MenuItem{
			Item:      v.(*models.Item),
			KindPrice: make([]*models.KindPrice, 0),
			SizePrice: make([]*models.SizePrice, 0),
		})
	}

	// get item size and price
	iAllItemSizes, err := itemSizeDT.Select().Exec()
	if err != nil {
		return nil, err
	}

	iAllSizes, err := sizeDT.Select().Exec()
	if err != nil {
		return nil, err
	}

	// get item kind and price
	iAllItemKinds, err := itemKindDT.Select().Exec()
	if err != nil {
		return nil, err
	}

	iAllKinds, err := kindDT.Select().Exec()
	if err != nil {
		return nil, err
	}

	// combine
	for i := 0; i < len(result); i++ {
		itemID := result[i].GetItem().GetID()

		linq.From(iAllItemSizes).Where(func(c interface{}) bool {
			return c.(*models.ItemSize).GetItemID() == itemID
		}).Select(func(c interface{}) interface{} {
			result := &models.SizePrice{}
			itemSize := c.(*models.ItemSize)

			result.Size = linq.From(iAllSizes).Where(func(c interface{}) bool {
				return c.(*models.Size).GetID() == itemSize.GetSizeID()
			}).Select(func(c interface{}) interface{} {
				return c.(*models.Size).GetName()
			}).First().(string)

			result.Price = itemSize.GetPrice()

			return result
		}).ToSlice(&(result[i].SizePrice))

		linq.From(iAllItemKinds).Where(func(c interface{}) bool {
			return c.(*models.ItemKind).GetItemID() == itemID
		}).Select(func(c interface{}) interface{} {
			result := &models.KindPrice{}
			itemKind := c.(*models.ItemKind)

			result.Kind = linq.From(iAllKinds).Where(func(c interface{}) bool {
				return c.(*models.Kind).GetID() == itemKind.GetKindID()
			}).Select(func(c interface{}) interface{} {
				return c.(*models.Kind).GetName()
			}).First().(string)

			result.Price = itemKind.GetPrice()

			return result
		}).ToSlice(&(result[i].KindPrice))
	}

	return result, nil
}

func (d *txtDb) AddShop(shop *models.Shop) (*models.Shop, error) {
	f, err := d.Connect(shopDT.TableName())
	if err != nil {
		return nil, err
	}
	defer f.Close()

	out, err := proto.Marshal(shop)
	if err != nil {
		return nil, err
	}

	_, err = f.Write(out)
	if err != nil {
		return nil, err
	}

	_, err = f.WriteString("\n")
	return shop, nil
}

package logic

import (
	"orderfood/src/database"
	"orderfood/src/database/models"
	"orderfood/src/handler/models/resp"

	linq "github.com/ahmetb/go-linq"
)

func GetMenu(shopName string) (menu *resp.ShopMenu, err error) {
	menu = nil
	db := database.Db.Menu()

	shop := &models.Shop{
		Name: shopName,
	}
	shops, err := db.GetShop(shop)
	if err != nil {
		return
	} else if len(shops) == 0 {
		err = NoDataError
		return
	}
	resShop := &resp.Shop{
		ID:   shops[0].GetID(),
		Name: shops[0].GetName(),
	}

	items, err := db.GetItem(&models.Item{
		Shop_ID: resShop.GetID(),
	})
	if err != nil {
		return
	} else if len(items) == 0 {
		menu = &resp.ShopMenu{
			Shop:  resShop,
			Items: make([]*resp.MenuItem, 0),
		}
		return
	}

	itemOptions, err := db.GetItemOption(nil)
	if err != nil {
		return
	}

	itemOptionSlice := make([]*resp.MenuItem, 0)
	linq.From(items).Join(linq.From(itemOptions),
		func(m interface{}) interface{} {
			o := m.(models.Item)
			return o.GetID()
		},
		func(m interface{}) interface{} {
			o := m.(models.ItemOption)
			return o.GetItem_ID()
		}, func(IItem interface{}, IItemOption interface{}) interface{} {
			item := IItem.(models.Item)
			itemOption := IItemOption.(models.ItemOption)

			options := make([]*resp.MenuOption, 0)
			options = append(options, &resp.MenuOption{
				OptionID: itemOption.GetOption_ID(),
			})

			return &resp.MenuItem{
				ItemID:  item.GetID(),
				Name:    item.GetName(),
				Price:   item.GetPrice(),
				Options: options,
			}
		}).ToSlice(&itemOptionSlice)

	// options, err := db.GetOption(nil)
	// if err != nil {
	// 	return
	// }

	return
}

func AddShop(name string) (*models.Shop, error) {
	db := database.Db.Menu()
	shop := &models.Shop{
		Name: name,
	}

	err := db.AddShop(shop)
	if err != nil {
		return nil, err
	}

	return shop, nil
}

func GetShop() ([]*models.Shop, error) {
	db := database.Db.Menu()
	shops, err :=db.GetShop(nil)
	return shops, err
}

func AddItem(shopID int32, name string) (*models.Item, error) {
	db := database.Db.Menu()
	item := &models.Item{
		Name: name,
		Shop_ID:shopID,
	}
	err := db.AddItem(item)
	 return item, err
}

func GetItem(shopID int32) ([]*models.Item, error) {
	db := database.Db.Menu()
	item := &models.Item{
		Shop_ID:shopID,
	}

	items, err :=db.GetItem(item)
	return items, err
}

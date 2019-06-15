package logic

import (
	"orderfood/src/database"
	"orderfood/src/database/models"
	"orderfood/src/handler/models/resp"
	"strconv"

	"strings"

	linq "github.com/ahmetb/go-linq"
)

func GetMenu(shopName string) (menu *resp.ShopMenu, err error) {
	// menu = nil
	// db := database.Db.Menu()

	// shop := &models.Shop{
	// 	Name: shopName,
	// }
	// shops, err := db.GetShop(shop)
	// if err != nil {
	// 	return
	// } else if len(shops) == 0 {
	// 	err = NoDataError
	// 	return
	// }
	// resShop := &resp.Shop{
	// 	ID:   shops[0].GetID(),
	// 	Name: shops[0].GetName(),
	// }

	// items, err := db.GetItem(&models.Item{
	// 	Option_ID: resShop.GetID(),
	// })
	// if err != nil {
	// 	return
	// } else if len(items) == 0 {
	// 	menu = &resp.ShopMenu{
	// 		Shop:  resShop,
	// 		Items: make([]*resp.MenuItem, 0),
	// 	}
	// 	return
	// }

	// itemOptions, err := db.GetItemOption(nil)
	// if err != nil {
	// 	return
	// }

	// itemOptionSlice := make([]*resp.MenuItem, 0)
	// linq.From(items).Join(linq.From(itemOptions),
	// 	func(m interface{}) interface{} {
	// 		o := m.(models.Item)
	// 		return o.GetID()
	// 	},
	// 	func(m interface{}) interface{} {
	// 		o := m.(models.ItemOption)
	// 		return o.GetItem_ID()
	// 	}, func(IItem interface{}, IItemOption interface{}) interface{} {
	// 		item := IItem.(models.Item)
	// 		itemOption := IItemOption.(models.ItemOption)

	// 		options := make([]*resp.MenuOption, 0)
	// 		options = append(options, &resp.MenuOption{
	// 			ID: itemOption.GetOption_ID(),
	// 		})

	// 		return &resp.MenuItem{
	// 			ID:      item.GetID(),
	// 			Name:    item.GetName(),
	// 			Price:   item.GetPrice(),
	// 			Options: options,
	// 		}
	// 	}).ToSlice(&itemOptionSlice)

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

func GetShop(id int, name string) ([]*models.Shop, error) {
	db := database.Db.Menu()
	shop := &models.Shop{
		ID:   id,
		Name: name,
	}
	shops, err := db.GetShop(shop)
	return shops, err
}

func UpdateShop(id int, name string) (bool, error) {
	db := database.Db.Menu()
	shop := &models.Shop{
		ID:   id,
		Name: name,
	}
	count, err := db.UpdateShop(shop)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func DeleteShop(id int) (bool, error) {
	db := database.Db.Menu()
	shop := &models.Shop{
		ID: id,
	}
	count, err := db.DeleteShop(shop)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func AddItem(shopID int, name string, price int) (*models.Item, error) {
	db := database.Db.Menu()
	item := &models.Item{
		Name:    name,
		Shop_ID: shopID,
		Price:   price,
	}
	err := db.AddItem(item)
	return item, err
}

func GetItem(shopID, optionID int) ([]*resp.Item, error) {
	db := database.Db.Menu()
	itemOptionView := &models.ItemOptionView{
		Shop_ID: shopID,
		Price:   -1,
	}
	itemOptionViews, err := db.GetItemOptionView(itemOptionView)
	if err != nil {
		return nil, err
	}

	selections, err := db.GetSelection(nil)
	if err != nil {
		return nil, err
	}

	// make option selection
	linq.From(selections).OrderBy(func(m interface{}) interface{} {
		selection := m.(*models.Selection)
		return selection.GetName()
	}).ToSlice(&selections)

	optionNameMap := make(map[int][]string)
	for _, selection := range selections {
		oi := selection.GetOption_ID()
		arr := make([]string, 0)
		if optionNameMap[oi] != nil {
			arr = optionNameMap[oi]
		}
		optionNameMap[oi] = append(arr, selection.GetName())
	}

	optionNames := make(map[int]string)
	for optionID, names := range optionNameMap {
		optionName := strings.Join(names, ",")
		optionNames[optionID] = optionName
	}

	linq.From(itemOptionViews).OrderBy(func(m interface{}) interface{} {
		itemOption := m.(*models.ItemOptionView)
		v := 0
		if itemOption.GetOption_ID() != nil {
			v = *itemOption.GetOption_ID()
		}
		return v
	}).ToSlice(&itemOptionViews)

	// find item ids
	itemIDs := make([]int, 0)
	linq.From(itemOptionViews).Where(
		func(m interface{}) bool {
			itemOption := m.(*models.ItemOptionView)
			if optionID == 0 {
				return true
			}
			if itemOption.GetOption_ID() != nil && optionID == *itemOption.GetOption_ID() {
				return true
			}
			return false
		}).Select(func(m interface{}) interface{} {
		itemOption := m.(*models.ItemOptionView)
		return itemOption.GetItem_ID()
	}).Distinct().ToSlice(&itemIDs)

	// make item option
	itemOptionIDMap := make(map[int][]int)
	for _, itemID := range itemIDs {
		for _, itemOption := range itemOptionViews {
			iID := itemOption.GetItem_ID()
			if iID != itemID {
				continue
			}

			arr := make([]int, 0)
			if itemOption.GetOption_ID() != nil {
				if itemOptionIDMap[itemID] != nil {
					arr = itemOptionIDMap[itemID]
				}
				arr = append(arr, *itemOption.GetOption_ID())
			}

			itemOptionIDMap[itemID] = arr
		}
	}

	items := make([]*resp.Item, 0)
	for itemID, optionIDs := range itemOptionIDMap {
		optionNameArr := make([]string, 0)
		for _, optionID := range optionIDs {
			name := optionNames[optionID]
			if name == "" {
				name = "no selection OptionID:" + strconv.Itoa(optionID)
			}
			optionNameArr = append(optionNameArr, name)
		}
		optionName := strings.Join(optionNameArr, "|")
		if optionName == "" {
			optionName = "無"
		}

		for _, itemOption := range itemOptionViews {
			if itemOption.GetItem_ID() == itemID {
				item := &resp.Item{
					ID:      int32(itemID),
					Options: optionName,
					Name:    itemOption.GetName(),
					Price:   int32(itemOption.GetPrice()),
				}
				items = append(items, item)
				break
			}
		}
	}

	return items, err
}

func UpdateItem(id int, name string, price int) (bool, error) {
	db := database.Db.Menu()
	item := &models.Item{
		ID:    id,
		Name:  name,
		Price: price,
	}
	count, err := db.UpdateItem(item)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func DeleteItem(id int) (bool, error) {
	db := database.Db.Menu()
	item := &models.Item{
		ID:    id,
		Price: -1,
	}
	count, err := db.DeleteItem(item)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, NoDataError
	} else {
		return true, nil
	}
}

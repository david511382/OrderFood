package manager

import (
	"orderfood/src/database"
	"orderfood/src/database/models"
	"orderfood/src/handler/models/resp"
	"orderfood/src/logic"
	"strconv"

	"strings"

	linq "github.com/ahmetb/go-linq"
)

func AddShop(name string) (*models.Shop, error) {
	db := database.Db.MenuShop()
	shop := &models.Shop{
		Name: name,
	}

	err := db.AddShop(shop, nil)
	if err != nil {
		return nil, err
	}

	return shop, nil
}

func GetShop(id int32, name string) ([]*models.Shop, error) {
	db := database.Db.MenuShop()
	shop := &models.Shop{
		ID:   id,
		Name: name,
	}
	shops, err := db.GetShop(shop)
	return shops, err
}

func UpdateShop(id int32, name string) (bool, error) {
	db := database.Db.MenuShop()
	shop := &models.Shop{
		ID:   id,
		Name: name,
	}
	count, err := db.UpdateShop(shop, nil)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func DeleteShop(id int32) (bool, error) {
	db := database.Db.MenuShop()
	shop := &models.Shop{
		ID: id,
	}
	count, err := db.DeleteShop(shop, nil)
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
	err := db.AddItem(item, nil)
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
	count, err := db.UpdateItem(item, nil)
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
	count, err := db.DeleteItem(item, nil)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, logic.NoDataError
	} else {
		return true, nil
	}
}

func AddItemOption(itemID, optionID int) (*models.ItemOption, error) {
	db := database.Db.Menu()
	itemOption := &models.ItemOption{
		Item_ID:   itemID,
		Option_ID: optionID,
	}
	err := db.AddItemOption(itemOption, nil)
	return itemOption, err
}

func DeleteItemOption(id int) (bool, error) {
	db := database.Db.Menu()
	itemOption := &models.ItemOption{
		ID: id,
	}
	count, err := db.DeleteItemOption(itemOption, nil)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func AddOption(selectNum int) (*models.Option, error) {
	db := database.Db.Menu()
	option := &models.Option{
		Select_Num: selectNum,
	}
	err := db.AddOption(option, nil)
	if err != nil {
		return nil, err
	}

	// _,err= AddSelection( option.GetID(),0,selectionName)
	// if err != nil {
	// 	db.DeleteOption(option)
	// 	return nil, err
	// }

	return option, nil
}

func UpdateOption(id, selectNum int) (bool, error) {
	db := database.Db.Menu()
	option := &models.Option{
		ID:         id,
		Select_Num: selectNum,
	}
	count, err := db.UpdateOption(option, nil)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func DeleteOption(id int) (bool, error) {
	db := database.Db.Menu()
	option := &models.Option{
		ID: id,
	}
	count, err := db.DeleteOption(option, nil)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func AddSelection(optionID, price int, name string) (*models.Selection, error) {
	db := database.Db.Menu()
	selection := &models.Selection{
		Name:      name,
		Option_ID: optionID,
		Price:     price,
	}
	err := db.AddSelection(selection, nil)
	if err != nil {
		return nil, err
	}

	return selection, nil
}

func UpdateSelection(id, price int, name string) (bool, error) {
	db := database.Db.Menu()
	selection := &models.Selection{
		ID:    id,
		Name:  name,
		Price: price,
	}
	count, err := db.UpdateSelection(selection, nil)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func DeleteSelection(id int) (bool, error) {
	db := database.Db.Menu()
	selection := &models.Selection{
		ID: id,
	}
	count, err := db.DeleteSelection(selection, nil)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

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
	// shops, err := GetShop(shop)
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


func GetShopMenu(shopID int) (*resp.ShopMenu, error) {
	shops,err:= GetShop(int32(shopID),"")
	if err != nil {
		return nil, err
	}
	if len(shops) == 0{
		return nil, NoDataError
	}
	shopMenu:=&resp.ShopMenu{
		Shop:&resp.Shop{
			ID:shops[0].GetID(),
			Name :shops[0].GetName(), 
		},
		Options:make([]*resp.MenuOption,0),
	}

	db := database.Db.Menu()
	itemOptionView := &models.ItemOptionView{
		Shop_ID: shopID,
		Price:   -1,
	}
	itemOptionViews, err := db.GetItemOptionView(itemOptionView)
	if err != nil {
		return nil, err
	}
	linq.From(itemOptionViews).OrderBy(func (m interface{})interface{}{
		v:= m.(*models.ItemOptionView)
		if  v.GetOption_ID()==nil{
			return 0
		}
		 return *v.GetOption_ID()
	}).ToSlice(&itemOptionViews)

	menuOptions,err:= getMenuOptions()
	linq.From(menuOptions).OrderBy(func (m interface{})interface{}{
		v:= m.(*resp.MenuOption)
		return v.GetOption().GetID()
	}).ToSlice(&menuOptions)

	// item id to option name map
	itemOptionNamesMap := make(map[int][]string)
	for _, itemOptionView := range itemOptionViews {
		itemID := itemOptionView.GetItem_ID()

		arr := make([]string, 0)
		if itemOptionView.GetOption_ID() != nil {
			if itemOptionNamesMap[itemID] != nil {
				arr = itemOptionNamesMap[itemID]
			}

			optionID32:=int32(*itemOptionView.GetOption_ID())
			for si,ei:=0,len(menuOptions)-1;;{
				k :=  int((ei-si)/2)+si
				id := menuOptions[k].GetOption().GetID()
				if optionID32 < id{
					ei = k
				} else if optionID32 > id {
					si = k+1
				}else{
					arr = append(arr, menuOptions[k].GetName())
					break
				}
			}
		}

		itemOptionNamesMap[itemID] = arr
	}
	itemOptionNameMap := make(map[int]string)
	for itemID,names:=range itemOptionNamesMap{
		name:= strings.Join(names,"|")
		itemOptionNameMap[itemID] = name
	}

	// combine
	lastOptionID:=-1
	ShopMenuOptionIndex:=-1
	for _,itemOptionView:=range itemOptionViews{
		if itemOptionView.GetShop_ID()!=shopID{
			continue
		}
		
		optionID:=0
		if itemOptionView.GetOption_ID() !=nil{
			optionID=*itemOptionView.GetOption_ID()
		}
		
		if optionID!=lastOptionID{
			lastOptionID = optionID	
			optionID32:=int32(optionID)
			
			newMenuOption:=resp.MenuOption{
				Option :nil,
				Name :"無",
				Items :make([]*resp.Item,0),
				Selections:nil,
			}
			if optionID32!=0{
				newMenuOption.Option = &resp.Option{
					ID:optionID32,
				}

				for _,menuOption:=range menuOptions{
					menuOptionID:=menuOption.GetOption().GetID()
					if menuOptionID==optionID32 {
						newMenuOption.Name = menuOption.Name
						newMenuOption.Option.SelectNum = menuOption.GetOption().GetSelectNum()
						newMenuOption.Selections =menuOption.GetSelections()
						break
					}
				}
			}

			shopMenu.Options = append(shopMenu.Options,&newMenuOption)
			ShopMenuOptionIndex++
		}
		
		itemID:=itemOptionView.GetItem_ID()
		shopMenu.Options[ShopMenuOptionIndex].Items = append(shopMenu.Options[ShopMenuOptionIndex].Items,
			&resp.Item{
				ID:int32(itemID),
				Name:itemOptionView.GetName(),
				Price:int32(itemOptionView.GetPrice()),
				Options:itemOptionNameMap[itemID],
			},
		)
	}

	return shopMenu,nil
}

func getMenuOptions()([]*resp.MenuOption,  error){
	db := database.Db.Menu()
	
	options,err:=db.GetOption(nil)
	if err != nil {
		return nil, err
	}

	selections, err := db.GetSelection(nil)
	if err != nil {
		return nil, err
	}
	linq.From(selections).OrderBy(func(m interface{}) interface{} {
		selection := m.(*models.Selection)
		return selection.GetName()
	}).ToSlice(&selections)

	result:=make([]*resp.MenuOption,0)
	for _,option:=range options{
		optionID:=option.GetID()
		menuSelections := make([]*resp.MenuSelection,0)
		selectionNames := make([]string,0)
		for _,selection := range selections{
			if selection.GetOption_ID() == optionID{
				selectionName:=selection.GetName()
				menuSelections = append(menuSelections,&resp.MenuSelection{
					ID :int32(selection.GetID()),
					Name :selectionName,
					Price :int32(selection.GetPrice()),
				})
				selectionNames = append(selectionNames,selectionName)
			}
		}
		optionName := strings.Join(selectionNames, ",")
			
		result= append(result,&resp.MenuOption{
			Option:&resp.Option{
				ID:int32(optionID),
				SelectNum:int32(option.GetSelect_Num()),
			} , 
			Name:optionName,
			Selections:menuSelections,
		})
	}

	return result,nil
}

func AddShop(name string) (*models.Shop, error) {
	db := database.Db.MenuShop()
	shop := &models.Shop{
		Name: name,
	}

	err := db.AddShop(shop)
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
	count, err := db.UpdateShop(shop)
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

func AddItemOption(itemID, optionID int) (*models.ItemOption, error) {
	db := database.Db.Menu()
	itemOption := &models.ItemOption{
		Item_ID:   itemID,
		Option_ID: optionID,
	}
	err := db.AddItemOption(itemOption)
	return itemOption, err
}

func DeleteItemOption(id int) (bool, error) {
	db := database.Db.Menu()
	itemOption := &models.ItemOption{
		ID: id,
	}
	count, err := db.DeleteItemOption(itemOption)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func AddOption(selectNum int, selectionName string) (*models.Option, error) {
	db := database.Db.Menu()
	option := &models.Option{
		Select_Num: selectNum,
	}
	err := db.AddOption(option)
	if err != nil {
		return nil, err
	}

	_,err= AddSelection( option.GetID(),0,selectionName)
	if err != nil {
		db.DeleteOption(option)
		return nil, err
	}

	return option, nil
}

func UpdateOption(id, selectNum int) (bool, error) {
	db := database.Db.Menu()
	option := &models.Option{
		ID:         id,
		Select_Num: selectNum,
	}
	count, err := db.UpdateOption(option)
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
	count, err := db.DeleteOption(option)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func AddSelection(optionID,price int, name string) (*models.Selection, error) {
	db := database.Db.Menu()
	selection := &models.Selection{
		Name:      name,
		Option_ID:optionID,
		Price:     price,
	}
	err := db.AddSelection(selection)
	if err != nil {
		return nil, err
	}

	return selection, nil
}

func UpdateSelection(id, price int,name string) (bool, error) {
	db := database.Db.Menu()
	selection := &models.Selection{
		ID:id,
		Name:      name,		
		Price:     price,
	}
	count, err := db.UpdateSelection(selection)
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
	count, err := db.DeleteSelection(selection)
	if err != nil {
		return false, err
	} else if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

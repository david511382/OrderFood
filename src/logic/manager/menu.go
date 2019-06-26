package manager

import (
	"orderfood/src/database"
	"orderfood/src/database/models"
	"orderfood/src/handler/models/reqs"
	"orderfood/src/handler/models/resp"
	"orderfood/src/logic"

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
	const (
		noneOptionName = "無"
		allOptionName  = "所有"
	)

	shops, err := GetShop(int32(shopID), "")
	if err != nil {
		return nil, err
	}
	if len(shops) == 0 {
		return nil, logic.NoDataError
	}
	shopMenu := &resp.ShopMenu{
		Shop: &resp.Shop{
			ID:   shops[0].GetID(),
			Name: shops[0].GetName(),
		},
		Options: make([]*resp.MenuOption, 0),
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
	linq.From(itemOptionViews).OrderBy(func(m interface{}) interface{} {
		v := m.(*models.ItemOptionView)
		if v.GetOption_ID() == nil {
			return 0
		}
		return *v.GetOption_ID()
	}).ToSlice(&itemOptionViews)

	menuOptions, err := getMenuOptions()
	linq.From(menuOptions).OrderBy(func(m interface{}) interface{} {
		v := m.(*resp.MenuOption)
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

			optionID32 := int32(*itemOptionView.GetOption_ID())
			for si, ei := 0, len(menuOptions)-1; ; {
				k := int((ei-si)/2) + si
				id := menuOptions[k].GetOption().GetID()
				if optionID32 < id {
					ei = k
				} else if optionID32 > id {
					si = k + 1
				} else {
					arr = append(arr, menuOptions[k].GetName())
					break
				}
			}
		}

		itemOptionNamesMap[itemID] = arr
	}
	itemOptionNameMap := make(map[int]string)
	for itemID, names := range itemOptionNamesMap {
		name := strings.Join(names, "|")
		if name == "" {
			name = noneOptionName
		}
		itemOptionNameMap[itemID] = name
	}

	// add 所有
	firstMenuOption := resp.MenuOption{
		Option:     nil,
		Name:       allOptionName,
		Items:      make([]*resp.Item, 0),
		Selections: nil,
	}
	for _, itemOptionView := range itemOptionViews {
		itemID := itemOptionView.GetItem_ID()
		firstMenuOption.Items = append(firstMenuOption.Items, &resp.Item{
			ID:      int32(itemID),
			Name:    itemOptionView.GetName(),
			Price:   int32(itemOptionView.GetPrice()),
			Options: itemOptionNameMap[itemID],
		})
	}
	shopMenu.Options = append(shopMenu.Options, &firstMenuOption)

	// combine
	lastOptionID := -1
	ShopMenuOptionIndex := 0
	for _, itemOptionView := range itemOptionViews {
		if itemOptionView.GetShop_ID() != shopID {
			continue
		}

		itemHasOption := itemOptionView.GetOption_ID() != nil
		optionID := 0
		if itemHasOption {
			optionID = *itemOptionView.GetOption_ID()
		}

		if optionID != lastOptionID {
			lastOptionID = optionID
			optionID32 := int32(optionID)

			newMenuOption := resp.MenuOption{
				Option:     nil,
				Name:       noneOptionName,
				Items:      make([]*resp.Item, 0),
				Selections: nil,
			}
			if itemHasOption {
				newMenuOption.Option = &resp.Option{
					ID: optionID32,
				}

				for _, menuOption := range menuOptions {
					menuOptionID := menuOption.GetOption().GetID()
					if menuOptionID == optionID32 {
						newMenuOption.Name = menuOption.Name
						newMenuOption.Option.SelectNum = menuOption.GetOption().GetSelectNum()
						newMenuOption.Selections = menuOption.GetSelections()
						break
					}
				}
			}

			shopMenu.Options = append(shopMenu.Options, &newMenuOption)
			ShopMenuOptionIndex++
		}

		itemID := itemOptionView.GetItem_ID()
		shopMenu.Options[ShopMenuOptionIndex].Items = append(shopMenu.Options[ShopMenuOptionIndex].Items,
			&resp.Item{
				ID:      int32(itemID),
				Name:    itemOptionView.GetName(),
				Price:   int32(itemOptionView.GetPrice()),
				Options: itemOptionNameMap[itemID],
			},
		)
	}

	return shopMenu, nil
}

func getMenuOptions() ([]*resp.MenuOption, error) {
	db := database.Db.Menu()

	options, err := db.GetOption(nil)
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

	result := make([]*resp.MenuOption, 0)
	for _, option := range options {
		optionID := option.GetID()
		menuSelections := make([]*resp.MenuSelection, 0)
		selectionNames := make([]string, 0)
		for _, selection := range selections {
			if selection.GetOption_ID() == optionID {
				selectionName := selection.GetName()
				menuSelections = append(menuSelections, &resp.MenuSelection{
					ID:    int32(selection.GetID()),
					Name:  selectionName,
					Price: int32(selection.GetPrice()),
				})
				selectionNames = append(selectionNames, selectionName)
			}
		}
		optionName := getOptionName(selectionNames)

		result = append(result, &resp.MenuOption{
			Option: &resp.Option{
				ID:        int32(optionID),
				SelectNum: int32(option.GetSelect_Num()),
			},
			Name:       optionName,
			Selections: menuSelections,
		})
	}

	return result, nil
}

func getOptionName(selectionNames []string) string {
	return strings.Join(selectionNames, ",")
}

func CreateOption(menuOption *reqs.MenuOption) (*resp.OptionMenu, error) {
	shopID32 := menuOption.ShopID
	shopID := int(shopID32)
	shops, err := GetShop(shopID32, "")
	if err != nil {
		return nil, err
	} else if len(shops) == 0 {
		return nil, logic.ParamError
	}

	selectionLen := len(menuOption.Selections)
	if int(menuOption.SelectNum) > selectionLen {
		return nil, logic.ParamError
	}

	itemLen := len(menuOption.Items)
	if selectionLen == 0 || itemLen == 0 {
		return nil, logic.ParamError
	}

	// insert option
	newOption, err := AddOption(int(menuOption.SelectNum))
	if err != nil {
		return nil, err
	}
	optionID := int(newOption.GetID())

	result := &resp.OptionMenu{
		ShopID: shopID32,
		MenuOption: &resp.MenuOption{
			Option: &resp.Option{
				ID:        int32(optionID),
				SelectNum: menuOption.SelectNum,
			},
			Name:       "",
			Items:      make([]*resp.Item, 0),
			Selections: make([]*resp.MenuSelection, 0),
		},
	}

	// insert selection
	selectionNames := make([]string, 0)
	for _, selection := range menuOption.Selections {
		newSelection, err := AddSelection(optionID, int(selection.GetPrice()), selection.GetName())
		if err != nil {
			return nil, err
		}

		selectionName := newSelection.GetName()
		selectionNames = append(selectionNames, selectionName)

		result.MenuOption.Selections = append(result.MenuOption.Selections, &resp.MenuSelection{
			ID:    int32(newSelection.GetID()),
			Name:  selectionName,
			Price: int32(newSelection.GetPrice()),
		})
	}

	// insert item
	for _, item := range menuOption.Items {
		newItem, err := AddItem(shopID, item.GetName(), int(item.GetPrice()))
		if err != nil {
			// maybe name duplicate
			return nil, err
		}

		_, err = AddItemOption(newItem.GetID(), optionID)
		if err != nil {
			return nil, err
		}

		result.MenuOption.Items = append(result.MenuOption.Items, &resp.Item{
			ID:    int32(newItem.GetID()),
			Name:  newItem.GetName(),
			Price: int32(newItem.GetPrice()),
		})
	}

	// make option name
	result.MenuOption.Name = getOptionName(selectionNames)

	return result, nil
}

package orm

import (
	"errors"
	"orderfood/src/database/models"
	"os"
)

var (
	dataPath string

	undefinedError error = errors.New("Undefined Error")
	typeError error = errors.New("Type Error")
	
	ShopDT *DbTable = &DbTable{
		Name:  "orderfood_menu.shop.txt",
		Model: &models.Shop{},
	}
	ItemDT *DbTable = &DbTable{
		Name:  "orderfood_menu.item.txt",
		Model: &models.Item{},
	}
	ItemOptionDT *DbTable = &DbTable{
		Name:  "orderfood_menu.item_option.txt",
		Model: &models.ItemOption{},
	}
	OptionDT *DbTable = &DbTable{
		Name:  "orderfood_menu.option.txt",
		Model: &models.Option{},
	}
	OptionSelectionDT *DbTable = &DbTable{
		Name:  "orderfood_menu.option_selection.txt",
		Model: &models.OptionSelection{},
	}
	SelectionDT *DbTable = &DbTable{
		Name:  "orderfood_menu.selection.txt",
		Model: &models.Selection{},
	}		
	
	MemberDT *DbTable = &DbTable{
		Name:  "order_member.user_info.txt",
		Model: &models.Member{},
	}
			
	allFileNames []string = []string{
		MemberDT.TableName(),
		ShopDT.TableName(),
		ItemDT.TableName(),
		ItemOptionDT.TableName(),
		OptionDT.TableName(),
		OptionSelectionDT.TableName(),
		SelectionDT.TableName(),
	}
)

func Init(root string) {
	dataPath = root
}

func CheckDb() error {
	for _, filename := range allFileNames {
		f, _, err := Connect(filename)
		if err != nil {
			return err
		}

		err = f.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func GetAllFilePaths() []string {
	result := make([]string, 0)
	for _, filename := range allFileNames {
		result = append(result, dataPath+string(os.PathSeparator)+filename)
	}
	return result
}

func Connect(filename string) (*os.File, string, error) {
	file := filepath(filename)

	f, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND, 0660)

	return f, file, err
}

func filepath(filename string) string {
	return dataPath + string(os.PathSeparator) + filename
}

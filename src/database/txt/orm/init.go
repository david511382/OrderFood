package orm

import (
	"errors"
	"orderfood/src/database/models"
	"os"
)

var (
	dataPath string

	undefinedError error = errors.New("Undefined Error")

	ShopDT *DbTable = &DbTable{
		Name:  "order_shop.shop.txt",
		Model: models.Shop{},
	}
	ShopItemDT *DbTable = &DbTable{
		Name:  "order_shop.shop_item.txt",
		Model: models.ShopItem{},
	}
	ItemDT *DbTable = &DbTable{
		Name:  "order_shop.item.txt",
		Model: models.Item{},
	}
	ItemSizeDT *DbTable = &DbTable{
		Name:  "order_shop.item_size.txt",
		Model: models.ItemSize{},
	}
	SizeDT *DbTable = &DbTable{
		Name:  "order_shop.size.txt",
		Model: models.Size{},
	}
	ItemKindDT *DbTable = &DbTable{
		Name:  "order_shop.item_kind.txt",
		Model: models.ItemKind{},
	}
	KindDT *DbTable = &DbTable{
		Name:  "order_shop.kind.txt",
		Model: models.Kind{},
	}

	MemberDT *DbTable = &DbTable{
		Name:  "order_member.user_info.txt",
		Model: models.Member{},
	}

	allFileNames []string = []string{
		MemberDT.TableName(),
		ShopDT.TableName(),
		ShopItemDT.TableName(),
		ItemDT.TableName(),
		ItemSizeDT.TableName(),
		SizeDT.TableName(),
		ItemKindDT.TableName(),
		KindDT.TableName(),
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

package txt

import (
	"errors"
	"orderfood/src/config"
	"orderfood/src/database/models"
	"orderfood/src/util"
	"os"
)

var (
	dataPath string

	undefinedError error = errors.New("Undefined Error")

	shopDT *dbTable = &dbTable{
		name:  "order_shop.shop.txt",
		model: models.Shop{},
	}
	shopItemDT *dbTable = &dbTable{
		name:  "order_shop.shop_item.txt",
		model: models.ShopItem{},
	}
	itemDT *dbTable = &dbTable{
		name:  "order_shop.item.txt",
		model: models.Item{},
	}
	itemSizeDT *dbTable = &dbTable{
		name:  "order_shop.item_size.txt",
		model: models.ItemSize{},
	}
	sizeDT *dbTable = &dbTable{
		name:  "order_shop.size.txt",
		model: models.Size{},
	}
	itemKindDT *dbTable = &dbTable{
		name:  "order_shop.item_kind.txt",
		model: models.ItemKind{},
	}
	kindDT *dbTable = &dbTable{
		name:  "order_shop.kind.txt",
		model: models.Kind{},
	}

	memberDT *dbTable = &dbTable{
		name:  "order_member.user_info.txt",
		model: models.Member{},
	}

	allFileNames []string = []string{
		memberDT.TableName(),
		shopDT.TableName(),
		shopItemDT.TableName(),
		itemDT.TableName(),
		itemSizeDT.TableName(),
		sizeDT.TableName(),
		itemKindDT.TableName(),
		kindDT.TableName(),
	}
)

type txtDb struct {
	Filepath string
}

func (db *txtDb) Connect(filename string) (*os.File, error) {
	file := filepath(filename)

	f, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND, 0660)

	if err == nil {
		db.Filepath = file
	}

	return f, err
}

func NewDb(dbCfg config.DbConfig) (*txtDb, error) {
	path, err := util.GetFilePath(dbCfg.Domain)
	if err != nil {
		return nil, err
	}
	dataPath = path

	d := &txtDb{}

	//check db
	for _, filename := range allFileNames {
		f, err := d.Connect(filename)
		if err != nil {
			return d, err
		}

		err = f.Close()
		if err != nil {
			return d, err
		}
	}

	return d, err
}

func (d *txtDb) RebuildDb(dbCfg config.DbConfig) error {
	for _, filename := range allFileNames {
		file := dataPath + string(os.PathSeparator) + filename

		f, err := os.Create(file)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	return nil
}

func filepath(filename string) string {
	return dataPath + string(os.PathSeparator) + filename
}

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
		Name:  "order_shop.shop.txt",
		Model: models.Shop{},
	}
	shopItemDT *dbTable = &dbTable{
		Name:  "order_shop.shop_item.txt",
		Model: models.ShopItem{},
	}
	itemDT *dbTable = &dbTable{
		Name:  "order_shop.item.txt",
		Model: models.Item{},
	}
	itemSizeDT *dbTable = &dbTable{
		Name:  "order_shop.item_size.txt",
		Model: models.ItemSize{},
	}
	sizeDT *dbTable = &dbTable{
		Name:  "order_shop.size.txt",
		Model: models.Size{},
	}
	itemKindDT *dbTable = &dbTable{
		Name:  "order_shop.item_kind.txt",
		Model: models.ItemKind{},
	}
	kindDT *dbTable = &dbTable{
		Name:  "order_shop.kind.txt",
		Model: models.Kind{},
	}

	memberDT *dbTable = &dbTable{
		Name:  "order_member.user_info.txt",
		Model: models.Member{},
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
	f, err := d.Connect("order_member.user_info.txt")
	defer f.Close()

	return d, err
}

func (d *txtDb) RebuildDb(dbCfg config.DbConfig) error {
	filename := "order_member.user_info.txt"
	file := dataPath + string(os.PathSeparator) + filename

	f, err := os.Create(file)
	defer f.Close()

	return err
}

func filepath(filename string) string {
	return dataPath + string(os.PathSeparator) + filename
}

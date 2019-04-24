package database

import (
	"orderfood/src/config"
	"orderfood/src/database/models"
	"orderfood/src/database/mysql"
	"orderfood/src/database/txt"

	_ "github.com/go-sql-driver/mysql"
)

var Db IDb

type IDBM interface {
	RebuildDb(dbCfg config.DbConfig) error
}

type IDb interface {
	GetMembers() ([]models.Member, error)
	GetMenus(shop string) ([]models.MenuItem, error)
	GetShops() ([]*models.Shop, error)
	GetItems() ([]*models.Item, error) 

	AddMembers(models.Member) error
	AddShop(*models.Shop) (*models.Shop, error)
	AddItem(*models.Item) (*models.Item, error)

	UpdateMembers(models.Member) error

	IDBM
}

func InitMysql(dbCfg config.DbConfig) error {
	d, err := mysql.NewDb(dbCfg)

	Db = d

	return err
}

func InitTxt(dbCfg config.DbConfig) error {
	d, err := txt.NewDb(dbCfg)

	Db = d

	return err
}

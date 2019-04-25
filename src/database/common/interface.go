package common

import (
	"orderfood/src/config"
	"orderfood/src/database/models"
)

type IDb interface {
	Member() IMember
	Menu() IMenu
	DBM() IDBM
}

type IDBM interface {
	RebuildDb(dbCfg config.DbConfig) error
}

type IMember interface {
	GetMembers() ([]models.Member, error)
	AddMembers(*models.Member) error
	UpdateMembers(models.Member) error
}

type IMenu interface {
	GetMenus(shop string) ([]models.MenuItem, error)
	GetShops() ([]*models.Shop, error)
	GetItems() ([]*models.Item, error)

	AddShop(*models.Shop) (*models.Shop, error)
	AddItem(*models.Item) (*models.Item, error)
}
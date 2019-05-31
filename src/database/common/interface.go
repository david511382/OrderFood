package common

import (
	"orderfood/src/database/models"
)

type IDb interface {
	Member() IMember
	Menu() IMenu
	DBM() IDBM
}

type IDBM interface {
	RebuildDb() error
}

type IMember interface {
	GetMembers() ([]models.Member, error)
	AddMembers(*models.Member) error
	UpdateMembers(models.Member) error
}

type IMenu interface {
	// Shop 。
	GetShop() ([]*models.Shop, error)
	AddShop() (*models.Shop, error)
	DeleteShop() error
	UpdateShop() (*models.Shop, error)

	// Item 。
	GetItem() ([]*models.Item, error)
	AddItem() (*models.Item, error)
	DeleteItem() error
	UpdateItem() (*models.Item, error)

	// ItemOption 。
	GetItemOption() ([]*models.ItemOption, error)
	AddItemOption() (*models.ItemOption, error)
	DeleteItemOption() error
	UpdateItemOption() (*models.ItemOption, error)

	// Option 。
	GetOption() ([]*models.Option, error)
	AddOption() (*models.Option, error)
	DeleteOption() error
	UpdateOption() (*models.Option, error)

	// OptionSelection 。
	GetOptionSelection() ([]*models.OptionSelection, error)
	AddOptionSelection() (*models.OptionSelection, error)
	DeleteOptionSelection() error
	UpdateOptionSelection() (*models.OptionSelection, error)

	// selection
	GetSelection() ([]*models.Selection, error)
	AddSelection() (*models.Selection, error)
	DeleteSelection() error
	UpdateSelection() (*models.Selection, error)
}

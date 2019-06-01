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
	GetMember(*models.Member) ([]models.Member, error)
	AddMember(*models.Member) (*models.Member, error)
	UpdateMember(*models.Member) (*models.Member, error)
	DeleteMember(*models.Member) error
}

type IMenu interface {
	// Shop 。
	GetShop(*models.Shop) ([]*models.Shop, error)
	AddShop(*models.Shop) (*models.Shop, error)
	DeleteShop(*models.Shop) error
	UpdateShop(*models.Shop) (*models.Shop, error)

	// Item 。
	GetItem(*models.Item) ([]*models.Item, error)
	AddItem(*models.Item) (*models.Item, error)
	DeleteItem(*models.Item) error
	UpdateItem(*models.Item) (*models.Item, error)

	// ItemOption 。
	GetItemOption(*models.ItemOption) ([]*models.ItemOption, error)
	AddItemOption(*models.ItemOption) (*models.ItemOption, error)
	DeleteItemOption(*models.ItemOption) error
	UpdateItemOption(*models.ItemOption) (*models.ItemOption, error)

	// Option 。
	GetOption(*models.Option) ([]*models.Option, error)
	AddOption(*models.Option) (*models.Option, error)
	DeleteOption(*models.Option) error
	UpdateOption(*models.Option) (*models.Option, error)

	// OptionSelection 。
	GetOptionSelection(*models.OptionSelection) ([]*models.OptionSelection, error)
	AddOptionSelection(*models.OptionSelection) (*models.OptionSelection, error)
	DeleteOptionSelection(*models.OptionSelection) error
	UpdateOptionSelection(*models.OptionSelection) (*models.OptionSelection, error)

	// selection
	GetSelection(*models.Selection) ([]*models.Selection, error)
	AddSelection(*models.Selection) (*models.Selection, error)
	DeleteSelection(*models.Selection) error
	UpdateSelection(*models.Selection) (*models.Selection, error)
}

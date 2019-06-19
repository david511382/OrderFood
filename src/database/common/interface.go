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
	CheckDb() error
	RebuildDb() error
}

type IMember interface {
	GetMember(*models.Member) ([]*models.Member, error)
	AddMember(*models.Member) error
	UpdateMember(*models.Member) (int64, error)
	DeleteMember(*models.Member) (int64, error)
}

type ishop interface {
	GetShop(*models.Shop) ([]*models.Shop, error)
	AddShop(*models.Shop) error
	DeleteShop(*models.Shop) (int64, error)
	UpdateShop(*models.Shop) (int64, error)
}

type IMenu interface {
	// Shop 。
	ishop

	// Item 。
	GetItem(*models.Item) ([]*models.Item, error)
	AddItem(*models.Item) error
	DeleteItem(*models.Item) (int64, error)
	UpdateItem(*models.Item) (int64, error)

	// Option 。
	GetOption(*models.Option) ([]*models.Option, error)
	AddOption(*models.Option) error
	DeleteOption(*models.Option) (int64, error)
	UpdateOption(*models.Option) (int64, error)

	// ItemOption 。
	GetItemOption(*models.ItemOption) ([]*models.ItemOption, error)
	AddItemOption(*models.ItemOption) error
	DeleteItemOption(*models.ItemOption) (int64, error)
	UpdateItemOption(*models.ItemOption) (int64, error)

	// selection
	GetSelection(*models.Selection) ([]*models.Selection, error)
	AddSelection(*models.Selection) error
	DeleteSelection(*models.Selection) (int64, error)
	UpdateSelection(*models.Selection) (int64, error)

	// item option view
	GetItemOptionView(*models.ItemOptionView) ([]*models.ItemOptionView, error)

	// option selection view
	GetOptionSelectionView(*models.OptionSelectionView) ([]*models.OptionSelectionView, error)
}

type IRedis interface {
	IMember

	// Shop 。
	ishop
}

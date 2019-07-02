package common

import (
	"database/sql"
	"orderfood/src/database/models"
)

type IDb interface {
	Member() IMember
	Menu() IMenu
	MenuShop() IShop
	DBM() IDBM
}

type IDBM interface {
	CheckDb() error
	RebuildDb() error
}

type IMember interface {
	GetMember(*models.Member) ([]*models.Member, error)
	AddMember(*models.Member, *sql.Tx) error
	UpdateMember(*models.Member, *sql.Tx) (int64, error)
	DeleteMember(*models.Member, *sql.Tx) (int64, error)
}

type IShop interface {
	GetShop(*models.Shop) ([]*models.Shop, error)
	AddShop(*models.Shop, *sql.Tx) error
	DeleteShop(*models.Shop, *sql.Tx) (int64, error)
	UpdateShop(*models.Shop, *sql.Tx) (int64, error)
}

type IMenu interface {
	ITransaction

	// Shop 。
	IShop

	// Item 。
	GetItem(*models.Item) ([]*models.Item, error)
	AddItem(*models.Item, *sql.Tx) error
	DeleteItem(*models.Item, *sql.Tx) (int64, error)
	UpdateItem(*models.Item, *sql.Tx) (int64, error)

	// Option 。
	GetOption(*models.Option) ([]*models.Option, error)
	AddOption(*models.Option, *sql.Tx) error
	DeleteOption(*models.Option, *sql.Tx) (int64, error)
	UpdateOption(*models.Option, *sql.Tx) (int64, error)

	// ItemOption 。
	GetItemOption(*models.ItemOption) ([]*models.ItemOption, error)
	AddItemOption(*models.ItemOption, *sql.Tx) error
	DeleteItemOption(*models.ItemOption, *sql.Tx) (int64, error)
	UpdateItemOption(*models.ItemOption, *sql.Tx) (int64, error)

	// selection
	GetSelection(*models.Selection) ([]*models.Selection, error)
	AddSelection(*models.Selection, *sql.Tx) error
	DeleteSelection(*models.Selection, *sql.Tx) (int64, error)
	UpdateSelection(*models.Selection, *sql.Tx) (int64, error)

	// item option view
	GetItemOptionView(*models.ItemOptionView) ([]*models.ItemOptionView, error)

	// option selection view
	GetOptionSelectionView(*models.OptionSelectionView) ([]*models.OptionSelectionView, error)
}

type ITransaction interface {
	Begin() (*sql.Tx, error)
}

type IRedisMember interface {
	IMember
}

type IRedisMenu interface {
	// Shop 。
	IShop
}

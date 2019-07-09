package common

import (
	"orderfood/src/database/models"

	"github.com/jmoiron/sqlx"
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
	AddMember(*models.Member, IExecer) error
	UpdateMember(*models.Member, IExecer) (int64, error)
	DeleteMember(*models.Member, IExecer) (int64, error)
}

type IShop interface {
	GetShop(*models.Shop) ([]*models.Shop, error)
	AddShop(*models.Shop, IExecer) error
	DeleteShop(*models.Shop, IExecer) (int64, error)
	UpdateShop(*models.Shop, IExecer) (int64, error)
}

type IMenu interface {
	IDbTransaction

	// Shop 。
	IShop

	// Item 。
	GetItem(*models.Item) ([]*models.Item, error)
	AddItem(*models.Item, IExecer) error
	DeleteItem(*models.Item, IExecer) (int64, error)
	UpdateItem(*models.Item, IExecer) (int64, error)

	// Option 。
	GetOption(*models.Option) ([]*models.Option, error)
	AddOption(*models.Option, IExecer) error
	DeleteOption(*models.Option, IExecer) (int64, error)
	UpdateOption(*models.Option, IExecer) (int64, error)

	// ItemOption 。
	GetItemOption(*models.ItemOption) ([]*models.ItemOption, error)
	AddItemOption(*models.ItemOption, IExecer) error
	DeleteItemOption(*models.ItemOption, IExecer) (int64, error)
	UpdateItemOption(*models.ItemOption, IExecer) (int64, error)

	// selection
	GetSelection(*models.Selection) ([]*models.Selection, error)
	AddSelection(*models.Selection, IExecer) error
	DeleteSelection(*models.Selection, IExecer) (int64, error)
	UpdateSelection(*models.Selection, IExecer) (int64, error)

	// item option view
	GetItemOptionView(*models.ItemOptionView) ([]*models.ItemOptionView, error)

	// option selection view
	GetOptionSelectionView(*models.OptionSelectionView) ([]*models.OptionSelectionView, error)
}

type IExecer interface {
	sqlx.Execer
}

type ITransaction interface {
	IExecer
	Commit() error
	Rollback() error
}

type IDbTransaction interface {
	Begin() (ITransaction, error)
}

type IRedisMember interface {
	IMember
}

type IRedisMenu interface {
	// Shop 。
	IShop
}

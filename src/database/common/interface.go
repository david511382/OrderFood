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

type IMenu interface {
	// Shop 。
	GetShop(*models.Shop) ([]*models.Shop, error)
	AddShop(*models.Shop) error
	DeleteShop(*models.Shop) (int64, error)
	UpdateShop(*models.Shop) (int64, error)

	// Item 。
	GetItem(*models.Item) ([]*models.Item, error)
	AddItem(*models.Item) error
	DeleteItem(*models.Item) (int64, error)
	UpdateItem(*models.Item) (int64, error)

	// Option 。
	GetGroup(*models.Group) ([]*models.Group, error)
	AddGroup(*models.Group) error
	DeleteGroup(*models.Group) (int64, error)
	UpdateGroup(*models.Group) (int64, error)

	// selection
	GetSelection(*models.Selection) ([]*models.Selection, error)
	AddSelection(*models.Selection) error
	DeleteSelection(*models.Selection) (int64, error)
	UpdateSelection(*models.Selection) (int64, error)
}

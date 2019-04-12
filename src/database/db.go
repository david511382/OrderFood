package database

import (
	"orderfood/src/config"
	"orderfood/src/database/models"
	"orderfood/src/database/mysql"
	"orderfood/src/database/txt"

	_ "github.com/go-sql-driver/mysql"
)

var Db IDb

type IDb interface {
	GetMembers() ([]models.Member, error)
	GetMenus(shop string) ([]models.MenuItem, error)

	AddMembers(models.Member) error
	UpdateMembers(models.Member) error
}

func InitMysql(dbCfg config.DbConfig) error {
	d, err := mysql.NewDb(dbCfg)

	Db = d

	return err
}

func RebuildMysql(dbCfg config.DbConfig) error {
	return mysql.Rebuild(dbCfg)
}

func InitTxt(dbCfg config.DbConfig) error {
	d, err := txt.NewDb(dbCfg)

	Db = d

	return err
}

func RebuildTxt(dbCfg config.DbConfig) error {
	return txt.Rebuild(dbCfg)
}

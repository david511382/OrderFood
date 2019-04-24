package database

import (
	"orderfood/src/config"
	"orderfood/src/database/common"
	"orderfood/src/database/mysql"
	"orderfood/src/database/txt"

	_ "github.com/go-sql-driver/mysql"
)

var Db common.IDb

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

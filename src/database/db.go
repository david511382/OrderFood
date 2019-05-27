package database

import (
	"orderfood/src/config"
	"orderfood/src/database/common"
	"orderfood/src/database/mysql"
	"orderfood/src/database/txt"
	"orderfood/src/util"

	_ "github.com/go-sql-driver/mysql"
)

var Db common.IDb

func InitMysql(dbCfg config.DbConfig) error {
	d := mysql.NewDb(dbCfg)

	//check db
	err := d.CheckDb()
	if err != nil {
		err = d.RebuildDb()
		if err != nil {
			return err
		}
	}

	Db = d

	return err
}

func InitTxt(dbCfg config.DbConfig) error {
	d, err := txt.NewDb(dbCfg)
	if err != nil {
		err = d.RebuildDb()
		if err != nil { // no folder
			path, err := util.GetFilePath("")
			if err != nil {
				return err
			}
			path += `\data`

			err = util.MakeFolderOn(path)
			if err != nil {
				return err
			}

			err = d.RebuildDb()
			if err != nil {
				return err
			}
		}
	}

	Db = d

	return err
}

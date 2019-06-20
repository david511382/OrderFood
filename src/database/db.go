package database

import (
	"orderfood/src/config"
	"orderfood/src/database/common"
	"orderfood/src/database/mysql"
	"orderfood/src/database/redis"
	"orderfood/src/database/txt"
	"orderfood/src/util"

	_ "github.com/go-sql-driver/mysql"
)

type db struct {
	shopDb   *shopDbSwitch
	memberDb *memberDbSwitch
}

var (
	Db common.IDb

	memberDb common.IMember
	menuDb   common.IMenu
	dbmDb    common.IDBM

	redisMemberDb common.IRedisMember
	redisMenuDb   common.IRedisMenu
)

func InitMysql(cfg *config.Config) error {
	memberDb = mysql.NewMemberDb(cfg.MySQLMember)
	menuDb = mysql.NewMenuDb(cfg.MySQLMenu)
	dbmDb = mysql.NewDBMdb(cfg.MySQLdbm)

	memberSwitcher := &memberDbSwitch{}
	shopSwitcher := &shopDbSwitch{}
	Db = &db{
		shopDb:   shopSwitcher,
		memberDb: memberSwitcher,
	}

	//check db
	err := Db.DBM().CheckDb()
	if err != nil {
		err = Db.DBM().RebuildDb()
		if err != nil {
			return err
		}
	}

	redisMemberDb, err = redis.NewMemberDb(cfg.RedisMember)
	if err != nil {
		memberSwitcher.redisStatus = false
	} else {
		err = memberSwitcher.initRedis()
		if err != nil {
			return err
		}
	}

	redisMenuDb, err = redis.NewMenuDb(cfg.RedisMenu)
	if err != nil {
		shopSwitcher.redisStatus = false
	} else {
		err = shopSwitcher.initRedis()
		if err != nil {
			return err
		}
	}

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

	//Db = d

	return err
}

func (d *db) Member() common.IMember {
	return d.memberDb
}

func (d *db) Menu() common.IMenu {
	return menuDb
}

func (d *db) MenuShop() common.IShop {
	return d.shopDb
}

func (d *db) DBM() common.IDBM {
	return dbmDb
}

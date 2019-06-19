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
	member common.IMember
	menu   common.IMenu
	dbm    common.IDBM

	redisMember common.IRedisMember
	redisMenu common.IRedisMenu
}

var Db common.IDb

func InitMysql(cfg *config.Config) error {
	redisMemberDb, _ := redis.NewMemberDb(cfg.RedisMember)
	redisMenuDb, _ := redis.NewMenuDb(cfg.RedisMenu)

	memberDb := mysql.NewMemberDb(cfg.MySQLMember)
	menuDb := mysql.NewMenuDb(cfg.MySQLMenu)
	dbmDb := mysql.NewDBMdb(cfg.MySQLdbm)
	Db = &db{
		member: memberDb,
		menu:   menuDb,
		dbm:    dbmDb,
		redisMember:  redisMemberDb,
		redisMenu:  redisMenuDb,
	}

	//check db
	err := Db.DBM().CheckDb()
	if err != nil {
		err = Db.DBM().RebuildDb()
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

	Db = d

	return err
}

func (d *db) Member() common.IMember {
	return d.member
}

func (d *db) Menu() common.IMenu {
	return d.menu
}

func (d *db) DBM() common.IDBM {
	return d.dbm
}

package mysql

import (
	"orderfood/src/config"
	"orderfood/src/database/common"
	"orderfood/src/database/mysql/member"
	"orderfood/src/database/mysql/menu"

	"orderfood/src/util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type mysqlDb struct {
	Connect func() (*sqlx.DB, error)
}

func NewMemberDb(dbCfg config.DbConfig) common.IMember {
	d := newDb(dbCfg)
	db := &member.MemberDb{
		Connect: d.Connect,
	}
	return db
}

func NewMenuDb(dbCfg config.DbConfig) common.IMenu {
	d := newDb(dbCfg)
	db := &menu.MenuDb{
		Connect: d.Connect,
	}
	return db
}

func NewDBMdb(dbCfg config.DbConfig) common.IDBM {
	d := newDb(dbCfg)
	return d
}

func newDb(dbCfg config.DbConfig) *mysqlDb {
	d := &mysqlDb{Connect: func() (*sqlx.DB, error) {
		db, err := sqlx.Open("mysql", dbCfg.MysqlURL())
		return db, err
	}}

	return d
}

func (d *mysqlDb) CheckDb() error {
	db, err := d.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("use orderfood_menu;")
	if err != nil {
		return err
	}
	_, err = db.Exec("use orderfood_member;")
	if err != nil {
		return err
	}

	return nil
}

func (d *mysqlDb) RebuildDb() error {
	//check db struct
	db, err := d.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	// Loads queries from file
	data, err := util.ReadFile("src/database/mysql/init_mysql.sql")
	if err != nil {
		return err
	}

	const end string = ";"
	endByte := ([]byte(end))[0]
	start := 0
	for i, v := range data {
		if v == endByte {
			sql := string(data[start:i])
			_, err = db.Exec(sql)
			if err != nil {
				return err
			}
			start = i + 1
		}
	}

	return err
}

package mysql

import (
	"fmt"
	"orderfood/src/database/common"
	"orderfood/src/util"
	"orderfood/tags"
	"testing"

	"github.com/jmoiron/sqlx"
)

type testDBM struct {
	*mysqlDb
}

var (
	memberDb common.IMember
)

func TestMain(m *testing.M) {
	cfg, _ := tags.InitConfig("../../config/test-config.yml")

	d := newDb(cfg.MySQLdbm)
	dbm := &testDBM{d}

	if err := dbm.CheckDb(); err != nil {
		if err = dbm.RebuildDb(); err != nil {
			panic(err)
		}
	}

	dbm.cleanDb()

	memberDb = NewMemberDb(cfg.MySQLMember)

	m.Run()
}

func (db *testDBM) cleanDb() {
	tables := []struct {
		schema string
		table  common.DbTable
	}{
		{
			schema: "orderfood_test_member",
			table:  common.MemberDt,
		},
	}

	d, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer d.Close()

	for _, dt := range tables {
		err = truncateTable(d, dt.schema, dt.table)
		if err != nil {
			panic(err)
		}
	}
}

func truncateTable(d *sqlx.DB, schema string, dt common.DbTable) error {
	sqlStr := "use %s;"
	sqlStr = fmt.Sprintf(sqlStr, schema)
	_, err := d.Exec(sqlStr)
	if err != nil {
		return err
	}

	sqlStr = "TRUNCATE TABLE %s;"
	sqlStr = fmt.Sprintf(sqlStr, dt.TableName)
	_, err = d.Exec(sqlStr)
	if err != nil {
		return err
	}

	return nil
}

func (db *testDBM) CheckDb() error {
	d, err := db.Connect()
	if err != nil {
		return err
	}
	defer d.Close()

	_, err = d.Exec("use orderfood_test_menu;")
	if err != nil {
		return err
	}
	_, err = d.Exec("use orderfood_test_member;")
	if err != nil {
		return err
	}

	return nil
}

func (d *testDBM) RebuildDb() error {
	//check db struct
	db, err := d.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	// Loads queries from file
	data, err := util.ReadFile("./init_test_mysql.sql")
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

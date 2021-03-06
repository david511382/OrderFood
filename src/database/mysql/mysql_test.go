package mysql

import (
	"fmt"
	"orderfood/src/database/common"
	"orderfood/src/database/models"
	"orderfood/src/util"
	"orderfood/tags"
	"testing"

	"github.com/jmoiron/sqlx"
)

type testDBM struct {
	*mysqlDb
}

const (
	i1 int = 1
	i2 int = 2
	i3 int = 3
	i4 int = 4
	i5 int = 5

	s1 string = "1"
	s2 string = "2"
	s3 string = "3"
	s4 string = "4"
	s5 string = "5"

	b1 bool = false
	b2 bool = true
)

var (
	memberDb common.IMember
	menuDb   common.IMenu

	memberDbMembers = []models.Member{
		models.Member{
			ID:       int32(i1),
			Name:     s1,
			Username: s1,
			Password: s1,
		},
		models.Member{
			ID:        int32(i2),
			Name:     s2,
			Username: s2,
			Password: s2,
		},
		models.Member{
			ID:        int32(i3),
			Name:     s3,
			Username: s3,
			Password: s3,
		},
		models.Member{
			ID:       int32(i4),
			Name:     s4,
			Username: s4,
			Password: s4,
		},
		models.Member{
			ID:        int32(i5),
			Name:     s5,
			Username: s5,
			Password: s5,
		},
	}

	menuDbShops = []models.Shop{
		models.Shop{
			ID:    int32(i1),
			Name: s1,
		},
		models.Shop{
			ID:    int32(i2),
			Name: s2,
		},
		models.Shop{
			ID:    int32(i3),
			Name: s3,
		},
	}
	menuDbItems = []models.Item{
		models.Item{
			ID:       i1,
			Name:     s1,
			Shop_ID: i1,
			Price:    i1,
		},
		models.Item{
			ID:       i2,
			Name:     s2,
			Shop_ID: i1,
			Price:    i2,
		},
		models.Item{
			ID:       i3,
			Name:     s3,
			Shop_ID: i2,
			Price:    i3,
		},
		models.Item{
			ID:       i4,
			Name:     s4,
			Shop_ID: i1,
			Price:    i4,
		},
		models.Item{
			ID:       i5,
			Name:     s5,
			Shop_ID: i1,
			Price:    i3,
		},
	}
	menuDbItemOption = []models.ItemOption{
		models.ItemOption{
			ID:               i1,
			Item_ID:          i1,
			Option_ID: i1,
		},
		models.ItemOption{
			ID:               i2,
			Item_ID:          i2,
			Option_ID: i2,
		},
		models.ItemOption{
			ID:               i3,
			Item_ID:          i3,
			Option_ID: i3,
		},
		models.ItemOption{
			ID:               i4,
			Item_ID:          i4,
			Option_ID: i1,
		},
	}
	menuDbOptions = []models.Option{
		models.Option{
			ID:               i1,
			Select_Num: i1,
		},
		models.Option{
			ID:               i2,
			Select_Num: i2,
		},
		models.Option{
			ID:               i3,
			Select_Num: i3,
		},
	}	
	menuDbSelections = []models.Selection{
		models.Selection{
			ID:       i1,
			Name:     s1,
			Option_ID: i1,
			Price:    i1,
		},
		models.Selection{
			ID:       i2,
			Name:     s2,
			Option_ID: i2,
			Price:    i2,
		},
		models.Selection{
			ID:       i3,
			Name:     s3,
			Option_ID: i3,
			Price:    i3,
		},
		models.Selection{
			ID:       i4,
			Name:     s4,
			Option_ID: i1,
			Price:    i4,
		},
		models.Selection{
			ID:       i5,
			Name:     s5,
			Option_ID: i2,
			Price:    i5,
		},
	}
)

func TestMain(m *testing.M) {
	cfg, _ := tags.InitConfig("../../config/test-config.yml")

	d := newDb(cfg.MySQLdbm)
	dbm := &testDBM{d}

	if err := dbm.RebuildDb(); err != nil {
		panic(err)
	}

	dbm.initDb()

	memberDb = NewMemberDb(cfg.MySQLMember)
	menuDb = NewMenuDb(cfg.MySQLMenu)

	m.Run()
}

func (db *testDBM) initDb() {
	const (
		memberSchema = "orderfood_test_member"
		menuSchema   = "orderfood_test_menu"
	)

	schemas := []struct {
		schema string
		table  []common.DbTable
		init   func(*sqlx.DB)
	}{
		{
			schema: memberSchema,
			table:  []common.DbTable{common.MemberDt},
			init: func(d *sqlx.DB) {
				sqlStr := `
				INSERT INTO %s
				(id,name,username,password)				
				VALUES
				(?,?,?,?)
				`
				sqlStr = fmt.Sprintf(sqlStr, memberSchema+"."+common.MemberDt.TableName)
				for _, member := range memberDbMembers {
					r, err := d.Exec(sqlStr, []interface{}{
						member.GetID(),
						member.GetName(),
						member.GetUsername(),
						member.GetPassword(),
					}...)
					if err != nil {
						panic(err)
					} else if count, err := r.RowsAffected(); count != 1 || err != nil {
						panic("insert fail")
					}
				}
			},
		},
		{
			schema: menuSchema,
			table: []common.DbTable{
				common.ItemDt,
			},
			init: func(d *sqlx.DB) {
				sqlStr := `
				INSERT INTO %s
				(id,name)				
				VALUES
				(?,?)
				`
				sqlStr = fmt.Sprintf(sqlStr, menuSchema+"."+common.ShopDt.TableName)
				for _, shop := range menuDbShops {
					r, err := d.Exec(sqlStr, []interface{}{
						shop.GetID(),
						shop.GetName(),
					}...)
					if err != nil {
						panic(err)
					} else if count, err := r.RowsAffected(); count != 1 || err != nil {
						panic("insert fail")
					}
				}

				sqlStr = `
				INSERT INTO %s
				(id,select_num)				
				VALUES
				(?,?)
				`
				sqlStr = fmt.Sprintf(sqlStr, menuSchema+"."+common.OptionDt.TableName)
				for _, option := range menuDbOptions {
					r, err := d.Exec(sqlStr, []interface{}{
						option.GetID(),
						option.GetSelect_Num(),
					}...)
					if err != nil {
						panic(err)
					} else if count, err := r.RowsAffected(); count != 1 || err != nil {
						panic("insert fail")
					}
				}

				sqlStr = `
				INSERT INTO %s
				(id,name,shop_id,price)				
				VALUES
				(?,?,?,?)
				`
				sqlStr = fmt.Sprintf(sqlStr, menuSchema+"."+common.ItemDt.TableName)
				for _, item := range menuDbItems {
					r, err := d.Exec(sqlStr, []interface{}{
						item.GetID(),
						item.GetName(),
						item.GetShop_ID(),
						item.GetPrice(),
					}...)
					if err != nil {
						panic(err)
					} else if count, err := r.RowsAffected(); count != 1 || err != nil {
						panic("insert fail")
					}
				}

				sqlStr = `
				INSERT INTO %s
				(id,item_id,option_id)				
				VALUES
				(?,?,?)
				`
				sqlStr = fmt.Sprintf(sqlStr, menuSchema+"."+common.ItemOptionDt.TableName)
				for _, itemOption := range menuDbItemOption {
					r, err := d.Exec(sqlStr, []interface{}{
						itemOption.GetID(),
						itemOption.GetItem_ID(),
						itemOption.GetOption_ID(),
					}...)
					if err != nil {
						panic(err)
					} else if count, err := r.RowsAffected(); count != 1 || err != nil {
						panic("insert fail")
					}
				}

				sqlStr = `
				INSERT INTO %s
				(id,name,option_id,price)				
				VALUES
				(?,?,?,?)
				`
				sqlStr = fmt.Sprintf(sqlStr, menuSchema+"."+common.SelectionDt.TableName)
				for _, selection := range menuDbSelections {
					r, err := d.Exec(sqlStr, []interface{}{
						selection.GetID(),
						selection.GetName(),
						selection.GetOption_ID(),
						selection.GetPrice(),
					}...)
					if err != nil {
						panic(err)
					} else if count, err := r.RowsAffected(); count != 1 || err != nil {
						panic("insert fail")
					}
				}
			},
		},
	}

	d, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer d.Close()

	for _, schema := range schemas {
		schema.init(d)
	}
}

// func truncateTable(d *sqlx.DB, schema string, dt common.DbTable) error {
// 	sqlStr := "use %s;"
// 	sqlStr = fmt.Sprintf(sqlStr, schema)
// 	_, err := d.Exec(sqlStr)
// 	if err != nil {
// 		return err
// 	}

// 	sqlStr = "TRUNCATE TABLE %s;"
// 	sqlStr = fmt.Sprintf(sqlStr, dt.TableName)
// 	_, err = d.Exec(sqlStr)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

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

package txt

import (
	"orderfood/src/config"
	"orderfood/src/database/common"
	// "orderfood/src/database/txt/member"
	// "orderfood/src/database/txt/menu"
	"orderfood/src/database/txt/orm"
	//"orderfood/src/util"
	"os"
)

type txtDb struct {
	member common.IMember
	menu   common.IMenu
}

func (d *txtDb) Member() common.IMember {
	return d.member
}

func (d *txtDb) Menu() common.IMenu {
	return d.menu
}

func (d *txtDb) DBM() common.IDBM {
	return nil
}

func (db *txtDb) Connect(filename string) (*os.File, error) {
	f, _, err := orm.Connect(filename)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func NewDb(dbCfg config.DbConfig) (*txtDb, error) {
	// path, err := util.GetFilePath(dbCfg.Domain)
	// if err != nil {
	// 	return nil, err
	// }
	// orm.Init(path)

	// db := &txtDb{}
	// // db.member = &member.MemberDb{}
	// // db.menu = &menu.MenuDb{}

	// //check db
	// if err := orm.CheckDb(); err != nil {
	// 	return db, err
	// }

	// return db, nil
	return nil, nil
}

func (d *txtDb) RebuildDb() error {
	allFileNames := orm.GetAllFilePaths()
	for _, file := range allFileNames {
		f, err := os.Create(file)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	return nil
}

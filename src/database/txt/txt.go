package txt

import (
	"orderfood/src/config"
	"orderfood/src/util"
	"os"
)

var (
	dataPath string
)

type txtDb struct {
	Connect func(file string) (*os.File, error)
}

func NewDb(dbCfg config.DbConfig) (*txtDb, error) {
	path, err := util.GetFilePath(dbCfg.Domain)
	if err != nil {
		return nil, err
	}
	dataPath = path

	d := &txtDb{Connect: func(filename string) (*os.File, error) {
		file := filepath(filename)

		f, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND, 0660)

		return f, err
	}}

	//check db
	f, err := d.Connect("order_member.user_info.txt")
	defer f.Close()

	return d, err
}

func Rebuild(dbCfg config.DbConfig) error {
	filename := "order_member.user_info.txt"
	file := dataPath + string(os.PathSeparator) + filename

	f, err := os.Create(file)
	defer f.Close()

	return err
}

func filepath(filename string) string {
	return dataPath + string(os.PathSeparator) + filename
}

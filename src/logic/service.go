package logic

import (
	"orderfood/src/config"
	"orderfood/src/database"
	"orderfood/src/util"

	"strconv"
	"strings"
)

const (
	R = "rice"
	V = "vag"
)

var (
	targetView = V
	UserOrders = make(map[string][]string)
)

func Init(cfg *config.Config) {
	initTxt(cfg.Txt)

	LoadMembers()
}

func initTxt(dbCfg config.DbConfig) {
	err := database.InitTxt(dbCfg)
	if err != nil {
		err = database.Db.RebuildDb(dbCfg)
		if err != nil { // no folder
			path, err := util.GetFilePath("")
			if err != nil {
				panic(err)
			}
			path += `\data`

			err = util.MakeFolderOn(path)
			if err != nil {
				panic(err)
			}

			err = database.Db.RebuildDb(dbCfg)
			if err != nil {
				panic(err)
			}
		}

		return
	}
}

func initMySQL(dbCfg config.DbConfig) {
	err := database.InitMysql(dbCfg)
	if err != nil {
		panic(err)
	}

	Members, err = database.Db.GetMembers()
	if err != nil {
		err = database.Db.RebuildDb(dbCfg)
		if err != nil {
			panic(err)
		}
	}
}

func SetView(view string) {
	targetView = view
}

func GetView() string {
	return targetView
}

func IntegrationOrders() (totalList string) {
	totalOrders := make(map[string]int)

	for _, orders := range UserOrders {
		for _, order := range orders {
			orderElements := strings.Split(order, " ")
			amount, _ := strconv.Atoi(orderElements[len(orderElements)-3])
			orderElements = orderElements[:len(orderElements)-3]

			clearOrder := strings.Join(orderElements, " ")

			totalOrders[clearOrder] += amount
		}
	}

	for order, amount := range totalOrders {
		totalList += order + " * " + strconv.Itoa(amount) + "\n"
	}

	return
}

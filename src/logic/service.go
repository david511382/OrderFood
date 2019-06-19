package logic

import (
	"orderfood/src/config"
	"orderfood/src/database"

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
	// init database
	err := initMySQL(cfg)
	if err != nil {
		initTxt(cfg.Txt)
	}

	LoadMembers()
}

func initTxt(dbCfg config.DbConfig) {
	err := database.InitTxt(dbCfg)
	if err != nil {
		panic(err)
	}
}

func initMySQL(Cfg *config.Config) error {
	err := database.InitMysql(Cfg)
	if err != nil {
		return err
	}

	return nil
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

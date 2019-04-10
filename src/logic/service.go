package logic

import (
	"orderfood/src/config"
	"orderfood/src/database"
	"orderfood/src/database/models"
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

	Members = make([]models.Member, 0)
)

func Init(cfg *config.Config) {
	// err := database.InitMysql(cfg.MySQL)
	// if err != nil {
	// 	panic(err)
	// }

	// Members, err = database.Db.GetMembers()
	// if err != nil {
	// 	err = database.RebuildMysql(cfg.MySQL)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	err := database.InitTxt(cfg.Txt)
	if err != nil {
		err = database.RebuildTxt(cfg.Txt)
		if err != nil {
			panic(err)
		}

		return
	}

	Members, err = database.Db.GetMembers()
	if err != nil {
		panic(err)
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

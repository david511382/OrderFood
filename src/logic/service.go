package logic

import (
	"strconv"
	"strings"
)

const R = "rice"
const V = "vag"

var targetView = V

var UserOrders = make(map[string][]string)

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

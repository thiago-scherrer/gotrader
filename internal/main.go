package main

import (
	"fmt"
)

func main() {
	for {
		daemonize()
	}
}

func daemonize() {
	initFlag()

	fmt.Println(
		helloMsg(),
	)
	telegramSend(
		helloMsg(),
	)

	typeOrder := candleRunner()
	waitCreateOrder()

	if typeOrder == "Buy" {
		closePositionProfitBuy()
	} else {
		closePositionProfitSell()
	}
	getProfit()
}

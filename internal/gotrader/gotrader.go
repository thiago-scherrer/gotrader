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

	msg1 := "Starting a new round! GoTrader!"
	fmt.Println(msg1)
	telegramSend(msg1)

	typeOrder := candleRunner()
	waitCreateOrder()
	closePositionProfit(typeOrder)
	getProfit()
}

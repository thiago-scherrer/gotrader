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

	asset := asset()
	candleTime := candle()
	hand := getHand()

	msg1 := "Starting a new round! GoTrader!"
	msg2 := "Asset: " + asset + "Candle time: " + string(candleTime) + "Hand: " + string(hand)

	fmt.Println(msg1)
	telegramSend(msg1)
	fmt.Println(msg2)
	telegramSend(msg2)

	typeOrder := candleRunner()
	waitCreateOrder()
	closePositionProfit(typeOrder)
	getProfit()
}

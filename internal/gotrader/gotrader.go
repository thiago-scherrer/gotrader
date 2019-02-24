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

	fmt.Println("Starting a new round! GoTrader!")
	fmt.Println("Asset: ", asset, "Candle time: ", candleTime, "Hand: ", hand)

	typeOrder := candleRunner()
	waitCreateOrder()
	closePositionProfit(typeOrder)
	getProfit()
}

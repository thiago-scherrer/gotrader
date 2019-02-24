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
	fmt.Println("Asset:", asset)
	fmt.Println("Candle time:", candleTime)
	fmt.Println("Hand:", hand)

	typeOrder := candleRunner()

	waitCreateOrder()
	closePositionProfit(typeOrder)

}

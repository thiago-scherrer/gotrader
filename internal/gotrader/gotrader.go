package main

import (
	"fmt"
	"time"
)

func main() {
	initFlag()

	trigger := threshold()
	var cSell int
	var cBuy int
	var oderid string
	var typeOrder string
	speed := speed()
	asset := asset()
	candleTime := candle()
	hand := getHand()

	fmt.Println("Starting gotrader!")
	fmt.Println("Asset:", asset)
	fmt.Println("Candle time:", candleTime)
	fmt.Println("Hand:", hand)

	for index := 0; index < trigger; index++ {
		fmt.Println("New candle: ", index)
		result := logicSystem()
		if result == "Buy" {
			cBuy++
		} else if result == "Sell" {
			cSell++
		}
	}

	for {
		if cBuy > cSell {
			oderid = makeBuy()
			typeOrder = "Buy"
			break
		} else if cSell > cBuy {
			oderid = makeSell()
			typeOrder = "Sell"
			break
		} else {
			panic("error to create a logic trigger")
		}
	}
	fmt.Println("Nice, order created: ", oderid)

	for {
		if statusOrder() {
			fmt.Println("Done, good Luck!")
			break
		} else {
			fmt.Println("Wainting order: ", oderid)
			time.Sleep(time.Duration(speed) * time.Second)
		}
	}

	for {
		if closePositionBuy() && typeOrder == "Buy" {
			fmt.Println("Closing buy position!")
			closePosition()
			break
		} else if closePositionSell() && typeOrder == "Sell" {
			fmt.Println("Closing sell position!")
			closePosition()
			break
		} else {
			time.Sleep(time.Duration(speed) * time.Second)
		}
	}

	for {
		if statusOrder() {
			fmt.Println("Profit done!")
			break
		} else {
			fmt.Println("Wainting for closing ...")
			time.Sleep(time.Duration(speed) * time.Second)
		}
	}

}

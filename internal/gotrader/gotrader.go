package main

import (
	"fmt"
	"time"
)

func main() {
	initFlag()

	trigger := StringToIntBit(threshold())
	var cSell int
	var cBuy int
	var oderid string
	var typeOrder string
	speed := StringToInt(
		speed(),
	)
	asset := asset()
	candleTime := candle()
	logic := logic()
	hand := getHand()

	fmt.Println("Starting gotrader!")
	fmt.Println("Asset:", asset)
	fmt.Println("Candle time:", candleTime)
	fmt.Println("Logic:", logic)
	fmt.Println("Hand:", hand)

	for index := 0; index <= trigger; index++ {
		fmt.Println("New candle: ", index)
		result := volume()
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

	fmt.Println("Ordem criada: ", oderid)
	for {
		if closePositionBuy() && typeOrder == "Buy" {
			closePosition()
			break
		} else if closePositionSell() && typeOrder == "Sell" {
			closePosition()
			break
		} else {
			time.Sleep(time.Duration(speed) * time.Second)
		}
	}
}

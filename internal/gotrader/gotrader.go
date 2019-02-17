package main

import (
	"fmt"
)

func main() {
	initFlag()

	asset := asset()
	candleTime := candle()
	logic := logic()
	hand := hand()

	fmt.Println("Starting gotrader!")
	fmt.Println("Asset:", asset)
	fmt.Println("Candle time:", candleTime)
	fmt.Println("Logic:", logic)
	fmt.Println("Hand:", hand)

	volume()
	
}

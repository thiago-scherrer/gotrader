package main

import "fmt"

func main() {
	cfile := configFile()
	asset := configReader("asset", cfile)
	candleTime := configReader("candle", cfile)
	logic := configReader("logic", cfile)
	hand := configReader("hand", cfile)

	fmt.Println("Starting gotrader!")
	fmt.Println("Asset:", asset)
	fmt.Println("Candle time:", candleTime)
	fmt.Println("Logic:", logic)
	fmt.Println("Hand:", hand)

	fmt.Println(volume(cfile))
}

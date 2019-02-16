package main

import (
	"encoding/json"
	"fmt"
)

func main() {
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

func sellOrder(configFile string) {
	path := "/api/v1/user/wallet"
	data := map[string]string{"message": "TDDRobot =)", "channelID": "1"}
	dataB, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	asset := asset()
	path = "/api/v1/order/&symbol=" + asset + "&side=SELL&orderQty=" + "1" + "&price=" + "3603,5" + "&ordType=Limit"
	getResult := clientRobot("GET", path, dataB)

	fmt.Println(BytesToString(getResult))
}

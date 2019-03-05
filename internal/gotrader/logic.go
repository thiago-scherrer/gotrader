package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func logicSystem() string {
	var apiresponse []APIResponseComplex
	var countSell int
	var countBuy int
	var result string

	asset := asset()
	candleTime := candle()
	path := "/api/v1/orderBook/L2?symbol=" + asset + "&depth=" + IntToString(depth())
	speed := speed()

	data := StringToBytes("message=GoTrader bot&channelID=1")

	for count := 0; count < candleTime; count++ {

		getResult := clientRobot("GET", path, data)
		err := json.Unmarshal(getResult, &apiresponse)
		if err != nil {
			fmt.Println("Error to get data to the logic, got", err)
		}

		for _, value := range apiresponse[:] {
			if value.Side == "Sell" {
				countSell = countSell + value.Size
			} else if value.Side == "Buy" {
				countBuy = countBuy + value.Size
			}
		}

		if verboseMode() {
			fmt.Println("New candle: ", count)
			fmt.Println("Number of Sell orders: ", countSell)
			fmt.Println("Number of Buy orders: ", countBuy)

		}
		time.Sleep(time.Duration(speed) * time.Second)
	}

	if countBuy > countSell {
		result = "Buy"
	} else if countSell > countBuy {
		result = "Sell"
	} else if countSell == countBuy {
		result = "Draw"
	} else {
		result = "Error"
		fmt.Println("Api result noting working! Buy: ", countBuy, " Sell: ", countSell)
	}

	return result
}

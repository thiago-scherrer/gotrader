package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func volume(confFile string) string {
	var apiresponse []APIResponseComplex
	var countSell int
	var countBuy int
	var result string

	asset := configReader("asset", confFile)
	candleTime := StringToIntBit(
		configReader("candle", confFile)) * 60
	path := "/api/v1/orderBook/L2?symbol=" + asset + "&depth=0"
	speed := StringToInt(
		configReader("speed", confFile),
	)

	for count := 0; count < candleTime; count++ {
		getResult := clientRobot("GET", confFile, path)
		err := json.Unmarshal(getResult, &apiresponse)
		if err != nil {
			panic(err)
		}

		for _, value := range apiresponse[:] {
			if value.Side == "Sell" {
				countSell = countSell + value.Size
			} else if value.Side == "Buy" {
				countBuy = countBuy + value.Size
			}
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
		fmt.Println("api result noting working! Buy: ", countBuy, " Sell: ", countSell)
	}

	return result
}

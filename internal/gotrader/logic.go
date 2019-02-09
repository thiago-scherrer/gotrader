package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func volume(userIDquery, secretQuery, endpoint, asset, candle string, hand, speed int64) string {
	var point []APIResponseComplex
	var countSell int
	var countBuy int
	var result string
	candleTime := StringToIntBit(candle) * 60

	for count := 0; count < candleTime; count++ {

		expired := IntToString((timeExpired()))
		path := "/api/v1/orderBook/L2?symbol=" + asset + "&depth=10"
		hexResult := hexCreator(secretQuery, "GET", path, expired)
		getResult := clientGet(hexResult,
			endpoint, path, expired, userIDquery)
		getByte := StringToBytes(getResult)

		err := json.Unmarshal(getByte, &point)
		if err != nil {
			panic(err)
		}

		for _, value := range point[:] {
			if value.Side == "Sell" {
				countSell = countSell + value.Size
			} else if value.Side == "Buy" {
				countBuy = countBuy + value.Size
			}
		}
		fmt.Println("Buy: ", countSell, "Sell: ", countBuy)
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

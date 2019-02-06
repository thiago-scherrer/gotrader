package main

import (
	"encoding/json"
	"fmt"
)

func volume(expired, userIDquery, secretQuery, endpoint, asset string, hand int64) string {
	var point []APIResponseComplex
	path := "/api/v1/orderBook/L2?symbol=" + asset + "&depth=0"
	hexResult := hexCreator(secretQuery, "GET", path, expired)
	getResult := clientGet(hexResult,
		endpoint, path, expired, userIDquery)
	getByte := StringToBytes(getResult)

	err := json.Unmarshal(getByte, &point)
	if err != nil {
		panic(err)
	}

	countSell := 0
	countBuy := 0
	for _, value := range point[:] {
		if value.Side == "Sell" {
			countSell++
		} else if value.Side == "Buy" {
			countBuy++
		}
	}
	var result string
	if countBuy > countSell {
		result = "Buy"
	} else if countSell > countBuy {
		result = "Sell"
	} else if countSell == countBuy {
		result = "Draw"
	} else {
		fmt.Println("api result noting working! Buy: ", countBuy, " Sell: ", countSell)
		panic(504)
	}
	fmt.Println("Buy:", countBuy)
	fmt.Println("Sell:", countSell)

	return result
}

package logic

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/thiago-scherrer/gotrader/central"
	"github.com/thiago-scherrer/gotrader/convert"
)

// CandleRunner verify the api and start the logic system
func CandleRunner() string {
	trigger := central.Threshold()
	var cSell int
	var cBuy int

	for index := 0; index < trigger; index++ {
		result := logicSystem()
		if result == "Buy" {
			cBuy++
		} else if result == "Sell" {
			cSell++
		} else {
			index = -1
		}
	}
	if central.VerboseMode() {
		fmt.Println("Buy orders:", cBuy, "Sell orders: ", cSell)
	}
	return order(cBuy, cSell)
}
func order(cBuy, cSell int) string {
	var typeOrder string

	for {
		if cBuy > cSell {
			typeOrder = "Buy"
			break
		} else if cSell > cBuy {
			typeOrder = "Sell"
			break
		} else {
			log.Fatalf("Draw, Starting a new round!")
		}
	}
	return typeOrder
}

func logicSystem() string {
	var apiresponse []central.APIResponseComplex
	var countSell int
	var countBuy int
	var result string

	asset := central.Asset()
	candleTime := central.Candle()
	path := "/api/v1/orderBook/L2?symbol=" + asset + "&depth=" + convert.IntToString(central.Depth())
	speed := central.Speed()

	data := convert.StringToBytes("message=GoTrader bot&channelID=1")

	for count := 0; count < candleTime; count++ {

		getResult := central.ClientRobot("GET", path, data)
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

		if central.VerboseMode() {
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
	if central.VerboseMode() {
		fmt.Println("Candle result:", result)
	}
	return result
}

package logic

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/thiago-scherrer/gotrader/central"
	"github.com/thiago-scherrer/gotrader/convert"
)

// Path from api to view the orderbook
const orderbook string = "/api/v1/orderBook/L2?"

// Used to return Buy to te bone
const tbuy = "Buy"

// Used to return Sell to te bone
const tsell = "Sell"

// Used to return Draw to te bone
const tdraw = "Draw"

// CandleRunner verify the api and start the logic system
func CandleRunner() string {
	trigger := central.Threshold()
	var cSell int
	var cBuy int

	for index := 0; index < trigger; index++ {
		result := logicSystem()
		if result == tbuy {
			cBuy++
		} else if result == tsell {
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
			typeOrder = tbuy
			break
		} else if cSell > cBuy {
			typeOrder = tsell
			break
		}
		log.Fatalf("Draw, Starting a new round!")
		typeOrder = tdraw
		break
	}
	return typeOrder
}

func logicSystem() string {
	var apiResponse []central.APIResponseComplex
	var countSell int
	var countBuy int
	depth := convert.IntToString(central.Depth())
	asset := central.Asset()
	candleTime := central.Candle()

	urlmap := url.Values{}
	urlmap.Set("symbol", asset)
	urlmap.Add("depth", depth)
	path := orderbook + urlmap.Encode()
	speed := central.Speed()

	// There is nothing important here,
	// but I can not leave empty so as not to break the request
	data := convert.StringToBytes("message=GoTrader bot&channelID=1")

	for count := 0; count < candleTime; count++ {

		getResult := central.ClientRobot("GET", path, data)
		err := json.Unmarshal(getResult, &apiResponse)
		if err != nil {
			fmt.Println("Error to get data to the logic, got", err)
		}

		for _, value := range apiResponse[:] {
			if value.Side == tsell {
				countSell = countSell + value.Size
			} else if value.Side == tbuy {
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
		return tbuy
	} else if countSell > countBuy {
		return tsell
	}
	return tdraw
}

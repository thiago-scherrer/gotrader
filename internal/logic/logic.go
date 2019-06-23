package logic

import (
	"encoding/json"
	"log"
	"net/url"
	"time"

	"github.com/thiago-scherrer/gotrader/internal/api"
	rd "github.com/thiago-scherrer/gotrader/internal/reader"
)

// Path from api to view the orderbook
const orb string = "/api/v1/orderBook/L2?"

// Return types to Buy
const tby = "Buy"
const tll = "Sell"
const tdw = "Draw"

// CandleRunner verify the api and start the logic system
func CandleRunner() string {
	t := rd.Threshold()
	var tsl int
	var cby int

	for i := 1; i < t; i++ {
		res := returnDepth()
		if res == tby {
			cby++
		} else if res == tll {
			tsl++
		} else {
			i = -1
		}
	}

	if cby > tsl {
		return "Buy"
	} else if tsl > cby {
		return "Sell"
	}

	return "Draw"
}

func logicSystem(buy, sell int) string {
	if buy > sell {
		return tby
	} else if sell > buy {
		return tll
	}
	return tdw
}

func timeStamp() string {
	ctm := rd.Candle()
	t := time.Now().UTC().Add(time.Duration(-ctm) * time.Minute)
	return t.Format("15:04")
}

func returnDepth() string {
	var sell int
	var buy int

	poh := "/api/v1/trade?"
	ap := rd.APIArray()
	t := timeStamp()
	d := rd.Data()

	u := url.Values{}
	u.Set("symbol", rd.Asset())
	u.Add("binSize", "1m")
	u.Add("partial", "false")
	u.Add("count", "500")
	u.Add("reverse", "false")
	u.Add("filter", `{"timestamp.time":"`+t+`"}`)

	p := poh + u.Encode()
	res := api.ClientRobot("GET", p, d)
	err := json.Unmarshal(res, &ap)
	if err != nil {
		log.Println("Error to get trade numbers:", err)
		return "Draw"
	}

	for _, v := range ap[:] {
		if v.Side == tll {
			sell = sell + v.Size
		} else if v.Side == tby {
			buy = buy + v.Size
		}
	}
	return logicSystem(buy, sell)
}

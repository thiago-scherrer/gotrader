package logic

import (
	"encoding/json"
	"log"
	"net/url"
	"time"

	"github.com/thiago-scherrer/gotrader/internal/api"
	"github.com/thiago-scherrer/gotrader/internal/convert"
	cvt "github.com/thiago-scherrer/gotrader/internal/convert"
	rd "github.com/thiago-scherrer/gotrader/internal/reader"
)

// Path from api to view the orderbook
const orb string = "/api/v1/orderBook/L2?"

// Used to return Buy to te bone
const tby = "Buy"

// Used to return Sell to te bone
const tll = "Sell"

// Used to return Draw to te bone
const tdw = "Draw"

// CandleRunner verify the api and start the logic system
func CandleRunner() string {
	trg := rd.Threshold()
	var tsl int
	var cbu int

	for index := 0; index < trg; index++ {
		res := logicSystem()
		if res == tby {
			cbu++
		} else if res == tll {
			tsl++
		} else {
			index = -1
		}
	}
	return order(cbu, tsl)
}

// order return the type of the oder to create, buy and sell
func order(cbu, tsl int) string {
	var trd string

	for {
		if cbu > tsl {
			trd = tby
			break
		} else if tsl > cbu {
			trd = tll
			break
		}
		log.Println("Draw, Starting a new round!")
		trd = tdw
		break
	}
	return trd
}

func logicSystem() string {
	ap := rd.APIArray()

	var cl int
	var cby int
	dth := returnDepth()
	ast := rd.Asset()
	ctm := rd.Candle()

	u := url.Values{}
	u.Set("symbol", ast)
	u.Add("depth", dth)
	pth := orb + u.Encode()
	spd := rd.Speed()

	// There is nothing important here,
	// but I can not leave empty so as not to break the request
	d := cvt.StringToBytes("message=GoTrader bot&channelID=1")

	for i := 0; i < ctm; i++ {

		g := api.ClientRobot("GET", pth, d)
		err := json.Unmarshal(g, &ap)
		if err != nil {
			log.Println("Error to get data to the logic, got", err)
		}

		for _, v := range ap[:] {
			if v.Side == tll {
				cl = cl + v.Size
			} else if v.Side == tby {
				cby = cby + v.Size
			}
		}
		time.Sleep(time.Duration(spd) * time.Second)
	}

	if cby > cl {
		return tby
	} else if cl > cby {
		return tll
	}
	return tdw
}

func returnDepth() string {
	poh := "/api/v1/trade/bucketed?"
	ap := rd.APIArray()
	t := time.Now().UTC()
	timestamp := t.Format("2006-01-02 15:04")

	data := cvt.StringToBytes("message=GoTrader bot&channelID=1")

	u := url.Values{}
	u.Set("symbol", rd.Asset())
	u.Add("binSize", "1m")
	u.Add("partial", "false")
	u.Add("count", "1")
	u.Add("reverse", "false")
	u.Add("filter", `{"timestamp":"`+timestamp+`"}`)

	p := poh + u.Encode()

	res := api.ClientRobot("GET", p, data)

	err := json.Unmarshal(res, &ap)
	if err != nil {
		log.Println("Error to get trade numbers:", err)
	}

	for _, v := range ap[:] {
		return convert.IntToString(v.Trades)
	}

	return "30"
}

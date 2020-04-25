package logic

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/thiago-scherrer/gotrader/internal/api"
	"github.com/thiago-scherrer/gotrader/internal/central"
	"github.com/thiago-scherrer/gotrader/internal/convert"
	"github.com/thiago-scherrer/gotrader/internal/display"
	rd "github.com/thiago-scherrer/gotrader/internal/reader"
)

// Path from api to view the orderbook
const orb string = "/api/v1/orderBook/L2?"

// CandleRunner verify the api and start the logic system
func CandleRunner() string {
	t := rd.Threshold()
	c := rd.Candle()
	var tsl int
	var cby int

	for i := int64(0); i < t; i++ {
		for ii := int64(0); ii < c; ii++ {
			res := returnDepth()
			if res == "Buy" {
				cby++
			} else if res == "Sell" {
				tsl++
			} else {
				i = -1
			}
		}
	}

	if cby > tsl {
		return "Buy"
	} else if tsl > cby {
		return "Sell"
	}

	return "Draw"
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
	u.Add("partial", "false")
	u.Add("count", "500")
	u.Add("reverse", "false")
	u.Add("filter", t)

	for index := 1; index <= 60; index++ {
		u.Set("start", strconv.Itoa(index))

		p := poh + u.Encode()
		res, code := api.ClientRobot("GET", p, d)

		if code != 200 {
			log.Println("Something wrong with api:", code, "Response: ", convert.BytesToString(res))
			time.Sleep(time.Duration(2) * time.Second)
		}

		err := json.Unmarshal(res, &ap)
		if err != nil {
			break
		}
		if len(res) <= 5 {
			break
		}

		for _, v := range ap[:] {
			if v.Side == "Sell" {
				sell = sell + v.Size
			} else if v.Side == "Buy" {
				buy = buy + v.Size
			}

		}
		time.Sleep(time.Duration(1) * time.Second)

	}

	return logicSystem(buy, sell)
}

func timeStamp() string {
	ctm := rd.Candle()

	t := time.Now().UTC().Add(time.Duration(-ctm) * time.Minute)

	date := t.Format("2006-01-02")

	time := t.Format("15:04")

	return `{"timestamp.date": "` + date + `", "timestamp.minute": "` + time + `" }`
}

func logicSystem(buy, sell int) string {
	if buy > sell {
		return "Buy"
	} else if sell > buy {
		return "Sell"
	}

	return "Draw"
}

func stopLossBuy(pst float64, price float64) bool {
	return price <= (pst - ((pst / 100) * rd.StopLoss()))
}

func stopLossSell(pst float64, price float64) bool {
	return price >= (pst + ((pst / 100) * rd.StopLoss()))
}

func closePositionBuy(pst float64, price float64) bool {
	return price >= (pst + ((pst / 100) * rd.Profit()))
}

func closePositionSell(pst float64, price float64) bool {
	return price <= (pst - ((pst / 100) * rd.Profit()))
}

func priceCloseBuy(price float64) string {
	priceClose := fmt.Sprintf("%2.f",
		(price + 0.5),
	)

	return priceClose
}

func priceCloseSell(price float64) string {
	priceClose := fmt.Sprintf("%2.f",
		(price - 0.5),
	)

	return priceClose
}

func lossTargetBuy(price float64) string {
	priceClose := fmt.Sprintf("%2.f",
		(price - 0.5),
	)

	return priceClose
}

func lossTargetSell(price float64) string {
	priceClose := fmt.Sprintf("%2.f",
		(price + 0.5),
	)

	return priceClose
}

// ClosePositionProfitBuy is responsible for closing a Buy order
func ClosePositionProfitBuy() {
	pst := central.GetPosition()

	for {
		price := central.Price()

		if closePositionBuy(pst, price) {
			profitTarget := priceCloseBuy(price)

			log.Println(display.OrdertriggerMsg(rd.Asset()))

			central.ClosePosition(profitTarget)
			if central.GetProfit() {
				break
			}
		} else if stopLossBuy(pst, price) {
			log.Println(display.StopLossMsg(rd.Asset()))

			central.ClosePosition(lossTargetBuy(price))
			if central.GetProfit() {
				break
			}
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}

// ClosePositionProfitSell is responsible for closing a Sell position
func ClosePositionProfitSell() {
	pst := central.GetPosition()

	for {
		price := central.Price()

		if closePositionSell(pst, price) {
			profitTarget := priceCloseSell(price)

			log.Println(display.OrdertriggerMsg(rd.Asset()))

			central.ClosePosition(profitTarget)
			if central.GetProfit() {
				break
			}
		} else if stopLossSell(pst, price) {
			log.Println(display.StopLossMsg(rd.Asset()))

			central.ClosePosition(lossTargetSell(price))
			if central.GetProfit() {
				break
			}
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}

package central

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/thiago-scherrer/gotrader/internal/api"
	cvt "github.com/thiago-scherrer/gotrader/internal/convert"
	dpl "github.com/thiago-scherrer/gotrader/internal/display"
	rd "github.com/thiago-scherrer/gotrader/internal/reader"
)

// Order path to use on API Request
const oph string = "/api/v1/order"

// Position path to use on API Request
const poh string = "/api/v1/position?"

// Basic path to use on API Request
const ith string = "/api/v1/instrument?"

// Laverage path to use on API Request
const lth = "/api/v1/position/leverage"

// A random number to make a sleep before staring a new round
const tlp = 50

// parserAmount unmarshal a r API to return the wallet amount
func parserAmount(data []byte) int {
	ap := rd.APISimple()
	err := json.Unmarshal(data, &ap)
	if err != nil {
		log.Println("Error to get Amount: ", err)
	}
	return ap.Amount
}

// lastPrice unmarshal a r API to return the last price
func lastPrice(d []byte) float64 {
	ap := rd.APIArray()
	var r float64

	err := json.Unmarshal(d, &ap)
	if err != nil {
		log.Println("Error to get last price: ", err)
	}
	for _, v := range ap[:] {
		r = v.LastPrice
	}
	return r
}

func makeOrder(orderType string) string {
	ap := rd.APISimple()
	hfl := cvt.IntToString(rd.Hand())
	ast := rd.Asset()
	prc := cvt.FloatToString(price())
	rtp := "POST"

	u := url.Values{}
	u.Set("symbol", ast)
	u.Add("side", orderType)
	u.Add("orderQty", hfl)
	u.Add("price", prc)
	u.Add("ordType", "Limit")
	data := cvt.StringToBytes(u.Encode())

	for {
		glt := api.ClientRobot(rtp, oph, data)
		err := json.Unmarshal(glt, &ap)
		if err != nil {
			log.Println("Error to make a order:", err)
			time.Sleep(time.Duration(5) * time.Second)
		} else {
			return ap.OrderID
		}
	}
}

func getPosition() float64 {
	ap := rd.APIArray()
	var r float64
	pth := poh + `filter={"symbol":"` + rd.Asset() + `"}&count=1`
	rtp := "GET"
	dt := cvt.StringToBytes("message=GoTrader bot&channelID=1")
	glt := api.ClientRobot(rtp, pth, dt)
	err := json.Unmarshal(glt, &ap)
	if err != nil {
		log.Println("Error to get pst:", err)
	}

	for _, v := range ap[:] {
		r = v.AvgEntryPrice
	}
	return r
}

func price() float64 {
	ast := rd.Asset()

	u := url.Values{}
	u.Set("symbol", ast)
	u.Add("count", "100")
	u.Add("reverse", "false")
	u.Add("columns", "lastPrice")

	p := ith + u.Encode()
	d := cvt.StringToBytes("message=GoTrader bot&channelID=1")
	g := api.ClientRobot("GET", p, d)

	return lastPrice(g)
}

func closePositionBuy(pst float64) bool {
	return price() >= (pst + ((pst / 100) * rd.Profit()))
}

func closePositionSell(pst float64) bool {
	return price() <= (pst - ((pst / 100) * rd.Profit()))
}

func closePosition() string {
	ast := rd.Asset()
	path := oph
	rtp := "POST"
	pst := getPosition()
	priceClose := fmt.Sprintf("%2.f", (pst +
		((pst / 100) * rd.Profit())))

	u := url.Values{}
	u.Set("symbol", ast)
	u.Add("execInst", "Close")
	u.Add("price", priceClose)
	u.Add("ordType", "Limit")
	data := cvt.StringToBytes(u.Encode())
	glt := api.ClientRobot(rtp, path, data)

	return cvt.BytesToString(glt)
}

func setLeverge() {
	ast := rd.Asset()
	path := lth
	rtp := "POST"
	l := rd.Leverage()
	u := url.Values{}
	u.Set("symbol", ast)
	u.Add("leverage", l)
	data := cvt.StringToBytes(u.Encode())

	api.ClientRobot(rtp, path, data)
	log.Println(dpl.SetleverageMsg(rd.Asset(), l))
	api.MatrixSend(dpl.SetleverageMsg(rd.Asset(), l))

}

func statusOrder() bool {
	path := poh + `filter={"symbol":"` + rd.Asset() + `"}&count=1`
	data := cvt.StringToBytes("message=GoTrader bot&channelID=1")
	glt := api.ClientRobot("GET", path, data)
	return opening(glt)
}

func opening(data []byte) bool {
	ap := rd.APIArray()
	var r bool

	err := json.Unmarshal(data, &ap)
	if err != nil {
		log.Println("json open error:", err)
	}
	for _, v := range ap[:] {
		r = v.IsOpen
	}
	return r
}

// CreateOrder create the order on bitmex
func CreateOrder(typeOrder string) {

	for {
		setLeverge()
		makeOrder(typeOrder)
		if waitCreateOrder() {
			log.Println(dpl.OrderCreatedMsg(rd.Asset(), typeOrder))
			api.MatrixSend(dpl.OrderCreatedMsg(rd.Asset(), typeOrder))
			break
		}
		time.Sleep(time.Duration(10) * time.Second)

	}
}

func waitCreateOrder() bool {

	for {
		if statusOrder() == true {
			log.Println(dpl.OrderDoneMsg(rd.Asset()))
			api.MatrixSend(dpl.OrderDoneMsg(rd.Asset()))
			return true
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}

// ClosePositionProfitBuy the Buy pst
func ClosePositionProfitBuy() bool {
	pst := getPosition()

	for {
		if closePositionBuy(pst) {
			log.Println(dpl.OrdertriggerMsg(rd.Asset()))
			api.MatrixSend(dpl.OrdertriggerMsg(rd.Asset()))
			closePosition()
			return true
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}

// ClosePositionProfitSell cloe the Buy position
func ClosePositionProfitSell() bool {
	pst := getPosition()

	for {
		if closePositionSell(pst) {
			log.Println(dpl.OrdertriggerMsg(rd.Asset()))
			api.MatrixSend(dpl.OrdertriggerMsg(rd.Asset()))
			closePosition()
			return true
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}

// GetProfit waint to start a new trade round
func GetProfit() bool {
	log.Println(dpl.OrderWaintMsg(rd.Asset()))
	api.MatrixSend(dpl.OrderWaintMsg(rd.Asset()))

	for {
		if statusOrder() == false {
			log.Println(dpl.ProfitMsg(rd.Asset()))
			api.MatrixSend(dpl.ProfitMsg(rd.Asset()))
			time.Sleep(time.Duration(tlp) * time.Second)
			return true
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}

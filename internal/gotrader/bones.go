package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// APIResponseComplex used to struct data from API response,
// thanks https://mholt.github.io/json-to-go/
type APIResponseComplex struct {
	Amount        int     `json:"amount"`
	AvgEntryPrice float64 `json:"avgEntryPrice"`
	ChannelID     int     `json:"channelID"`
	IsOpen        bool    `json:"isOpen"`
	ID            int64   `json:"id"`
	LastPrice     float64 `json:"lastPrice"`
	OrderID       string  `json:"orderID"`
	OrderQty      int     `json:"orderQty"`
	Price         float64 `json:"price"`
	Side          string  `json:"side"`
	Size          int     `json:"size"`
}

// Conf instruction are the file yaml on disc
type Conf struct {
	Asset     string  `yaml:"asset"`
	Candle    int     `yaml:"candle"`
	Depth     int64   `yaml:"depth"`
	Endpoint  string  `yaml:"endpoint"`
	Hand      int     `yaml:"hand"`
	Profit    float64 `yaml:"profit"`
	Secret    string  `yaml:"secret"`
	Threshold int     `yaml:"threshold"`
	Userid    string  `yaml:"userid"`
}

func initFlag() string {
	var config string
	if len(os.Args[1:]) == 0 {
		panic("Usage : config config.yml")
	}
	if os.Args[1] == "config" {
		config = os.Args[2]
	} else {
		panic("Usage : config config.yml")
	}
	return config
}

func configReader() *Conf {
	confFile := initFlag()
	conf := Conf{}
	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic(err)
	}
	return &conf
}

func userid() string {
	conf := configReader()
	return conf.Userid
}

func secret() string {
	conf := configReader()
	return conf.Secret
}

func endpoint() string {
	conf := configReader()
	return conf.Endpoint
}

func hand() int {
	conf := configReader()
	return conf.Hand
}

func speed() int {
	return 10
}

func asset() string {
	conf := configReader()
	return conf.Asset
}

func candle() int {
	conf := configReader()
	return conf.Candle * 6
}

func profit() float64 {
	conf := configReader()
	return conf.Profit
}

func threshold() int {
	conf := configReader()
	return conf.Threshold
}

func handRoll(getParser, hand int) int {
	return (getParser * hand) / 100
}

func depth() int64 {
	conf := configReader()
	return conf.Depth
}

func hexCreator(secret, requestTipe, path, expired, data string) string {
	concat := requestTipe + path + expired + data
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(concat))
	hexResult := hex.EncodeToString(h.Sum(nil))
	return hexResult
}

func parserAmount(data []byte) int {
	apiresponse := APIResponseComplex{}
	err := json.Unmarshal(data, &apiresponse)
	if err != nil {
		fmt.Println(err)
	}
	return apiresponse.Amount
}

func opening(data []byte) bool {
	var apiresponse []APIResponseComplex
	var result bool

	err := json.Unmarshal(data, &apiresponse)
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range apiresponse[:] {
		result = value.IsOpen
	}
	return result
}

func lastPrice(data []byte) float64 {
	var apiresponse []APIResponseComplex
	var result float64

	err := json.Unmarshal(data, &apiresponse)
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range apiresponse[:] {
		result = value.LastPrice
	}
	return result
}

func timeExpired() int64 {
	timeExpired := timeStamp() + 60
	return timeExpired
}

func timeStamp() int64 {
	now := time.Now()
	timestamp := now.Unix()
	return timestamp
}

func getHand() int {
	path := "/api/v1/user/wallet"
	requestTipe := "GET"
	hand := hand()
	data := StringToBytes("message=GoTrader bot&channelID=1")
	getResult := clientRobot(requestTipe, path, data)
	return (parserAmount(getResult) * hand) / 100
}

func makeSell() string {
	apiresponse := APIResponseComplex{}
	qtyOrerFloat := (price() * float64(getHand())) / 10000000
	qtyOrder := FloatToInt(qtyOrerFloat)
	asset()
	path := "/api/v1/order"
	requestTipe := "POST"
	data := StringToBytes("symbol=" + asset() + "&side=Sell&orderQty=" +
		IntToString(qtyOrder) + "&price=" + FloatToString(price()) + "&ordType=Limit")

	getResult := clientRobot(requestTipe, path, data)

	err := json.Unmarshal(getResult, &apiresponse)
	if err != nil {
		fmt.Println(err)
	}

	return apiresponse.OrderID
}

func makeBuy() string {
	apiresponse := APIResponseComplex{}
	qtyOrerFloat := (price() * float64(getHand())) / 10000000
	qtyOrder := FloatToInt(qtyOrerFloat)
	asset()
	path := "/api/v1/order"
	requestTipe := "POST"
	data := StringToBytes("symbol=" + asset() + "&side=Buy&orderQty=" +
		IntToString(qtyOrder) + "&price=" + FloatToString(price()) + "&ordType=Limit")

	getResult := clientRobot(requestTipe, path, data)

	err := json.Unmarshal(getResult, &apiresponse)
	if err != nil {
		fmt.Println(err)
	}

	return apiresponse.OrderID
}

func getPosition() float64 {
	var apiresponse []APIResponseComplex
	var result float64

	path := "/api/v1/position" + "?symbol=" + asset() + "&count=1"
	requestTipe := "GET"
	data := StringToBytes("message=GoTrader bot&channelID=1")
	getResult := clientRobot(requestTipe, path, data)

	err := json.Unmarshal(getResult, &apiresponse)
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range apiresponse[:] {
		result = value.AvgEntryPrice
	}
	return result
}

func price() float64 {
	asset := asset()
	path := "/api/v1/instrument?symbol=" + asset + "&count=100&reverse=false&columns=lastPrice"
	data := StringToBytes("message=GoTrader bot&channelID=1")
	getResult := clientRobot("GET", path, data)
	return lastPrice(getResult)
}

func closePositionBuy() bool {
	return price() >= (getPosition() + ((getPosition() / 100) * profit()))
}

func closePositionSell() bool {
	return price() <= (getPosition() - ((getPosition() / 100) * profit()))
}

func closePosition() string {
	path := "/api/v1/order"
	requestTipe := "POST"
	priceClose := fmt.Sprintf("%2.f", (getPosition() + ((getPosition() / 100) * profit())))
	data := StringToBytes("symbol=" + asset() +
		"&execInst=Close" + "&price=" + priceClose + "&ordType=Limit")

	getResult := clientRobot(requestTipe, path, data)
	return BytesToString(getResult)
}

func statusOrder() bool {
	asset := asset()
	path := "/api/v1/position?symbol=" + asset + "&count=1"
	data := StringToBytes("message=GoTrader bot&channelID=1")
	getResult := clientRobot("GET", path, data)

	return opening(getResult)
}

func candleRunner() string {
	trigger := threshold()
	var cSell int
	var cBuy int

	for index := 0; index < trigger; index++ {
		fmt.Println("New candle: ", index)
		result := logicSystem()
		if result == "Buy" {
			cBuy++
		} else if result == "Sell" {
			cSell++
		}
	}
	return createOrder(cBuy, cSell)
}

func createOrder(cBuy, cSell int) string {
	var oderid string
	var typeOrder string

	for {
		if cBuy > cSell {
			oderid = makeBuy()
			typeOrder = "Buy"
			fmt.Println("Nice, BUY order created: ", oderid)
			break
		} else if cSell > cBuy {
			oderid = makeSell()
			typeOrder = "Sell"
			fmt.Println("Nice, SELL order created: ", oderid)
			break
		} else {
			typeOrder = "Draw"
		}
	}
	return typeOrder
}

func waitCreateOrder() bool {
	speed := speed()

	for {
		if statusOrder() == true {
			fmt.Println("Done, good Luck!")
			return true
		}
		time.Sleep(time.Duration(speed) * time.Second)
	}
}

func closePositionProfit(typeOrder string) bool {
	speed := speed()
	fmt.Println("Wainting position close...")

	for {
		if closePositionBuy() == true && typeOrder == "Buy" {
			fmt.Println("Closing buy position!")
			closePosition()
			return true
		} else if closePositionSell() == true && typeOrder == "Sell" {
			fmt.Println("Closing sell position!")
			closePosition()
			return true
		} else {
			time.Sleep(time.Duration(speed) * time.Second)
		}
	}
}

func getProfit() bool {
	speed := speed()
	fmt.Println("Wainting get profit...")

	for {
		if !statusOrder() {
			fmt.Println("Profit done!")
			time.Sleep(time.Duration(speed+50) * time.Second)
			return true
		}
		time.Sleep(time.Duration(speed) * time.Second)
	}
}

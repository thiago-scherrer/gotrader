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
	Asset           string  `yaml:"asset"`
	Candle          int     `yaml:"candle"`
	Depth           int64   `yaml:"depth"`
	Endpoint        string  `yaml:"endpoint"`
	Hand            int     `yaml:"hand"`
	Leverage        string  `yaml:"leverage"`
	Profit          float64 `yaml:"profit"`
	Secret          string  `yaml:"secret"`
	Threshold       int     `yaml:"threshold"`
	Userid          string  `yaml:"userid"`
	TelegramUse     bool    `yaml:"telegramuse"`
	TelegramKey     string  `yaml:"telegram_key"`
	TelegramURL     string  `yaml:"telegramurl"`
	TelegramChannel string  `yaml:"telegramchannel"`
	Verbose         bool    `yaml:"verbose"`
}

func initFlag() string {
	var config string
	if len(os.Args[1:]) == 0 {
		panic(usageMsg())
	}
	if os.Args[1] == "config" {
		config = os.Args[2]
	} else {
		panic(usageMsg())
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

func telegramurl() string {
	conf := configReader()
	return conf.TelegramURL
}

func secret() string {
	conf := configReader()
	return conf.Secret
}

func endpoint() string {
	conf := configReader()
	return conf.Endpoint
}

func telegramUse() bool {
	conf := configReader()
	return conf.TelegramUse
}

func telegramKey() string {
	conf := configReader()
	return conf.TelegramKey
}

func telegramChannel() string {
	conf := configReader()
	return conf.TelegramChannel
}

func hand() int {
	conf := configReader()
	return conf.Hand
}

func leverage() string {
	conf := configReader()
	return conf.Leverage
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

func verboseMode() bool {
	conf := configReader()
	return conf.Verbose
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
	if err != nil && verboseMode() {
		fmt.Println("Error to get Amount: ", err)
	}
	return apiresponse.Amount
}

func lastPrice(data []byte) float64 {
	var apiresponse []APIResponseComplex
	var result float64

	err := json.Unmarshal(data, &apiresponse)
	if err != nil && verboseMode() {
		fmt.Println("Error to get last price: ", err)
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

	if verboseMode() {
		fmt.Println("DATA get hand: ", hand)
	}

	data := StringToBytes("message=GoTrader bot&channelID=1")
	getResult := clientRobot(requestTipe, path, data)
	return (parserAmount(getResult) * hand) / 100
}

func makeOrder(orderType string) string {
	apiresponse := APIResponseComplex{}
	qtyOrerFloat := (price() * float64(getHand())) / 10000000
	qtyOrder := FloatToInt(qtyOrerFloat)
	asset()
	path := "/api/v1/order"
	requestTipe := "POST"

	if verboseMode() {
		fmt.Println("DATA make order: " + "symbol=" + asset() + "&side=" +
			orderType + "&orderQty=" + IntToString(qtyOrder) + "&price=" +
			FloatToString(price()) + "&ordType=Limit")
	}

	data := StringToBytes("symbol=" + asset() + "&side=" + orderType + "&orderQty=" +
		IntToString(qtyOrder) + "&price=" + FloatToString(price()) + "&ordType=Limit")

	getResult := clientRobot(requestTipe, path, data)

	err := json.Unmarshal(getResult, &apiresponse)
	if err != nil && verboseMode() {
		fmt.Println("Error to make a order:", err)
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
	if err != nil && verboseMode() {
		fmt.Println("Error to get position:", err)
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

func closePositionBuy(position float64) bool {
	if verboseMode() {
		fmt.Println("Close Position Buy: ", position+
			(position/100)*profit())
	}
	return price() >= (position + ((position / 100) * profit()))
}

func closePositionSell(position float64) bool {
	if verboseMode() {
		fmt.Println("Close Position Sell: ", position+
			(position/100)*profit())
	}
	return price() <= (position - ((position / 100) * profit()))
}

func closePosition() string {
	path := "/api/v1/order"
	requestTipe := "POST"
	position := getPosition()
	priceClose := fmt.Sprintf("%2.f", (position +
		((position / 100) * profit())))

	if verboseMode() {
		fmt.Println("Data close position: " + "symbol=" + asset() +
			"&execInst=Close" + "&price=" + priceClose + "&ordType=Limit")
	}

	data := StringToBytes("symbol=" + asset() +
		"&execInst=Close" + "&price=" + priceClose + "&ordType=Limit")
	getResult := clientRobot(requestTipe, path, data)
	return BytesToString(getResult)
}

func setLeverge() {
	asset()
	path := "/api/v1/position/leverage"
	requestTipe := "POST"

	if verboseMode() {
		fmt.Println("Data leverge: " + "symbol=" + asset() +
			"&leverage=" + leverage())
	}

	data := StringToBytes("symbol=" + asset() + "&leverage=" + leverage())
	clientRobot(requestTipe, path, data)

	fmt.Println(setlavarageMsg())
	telegramSend(setlavarageMsg())
}

func statusOrder() bool {
	asset := asset()
	path := "/api/v1/position?symbol=" + asset + "&count=1"

	data := StringToBytes("message=GoTrader bot&channelID=1")
	getResult := clientRobot("GET", path, data)
	if verboseMode() {
		fmt.Println("Data status order: " + BytesToString(getResult))
	}
	return opening(getResult)
}

func opening(data []byte) bool {
	var apiresponse []APIResponseComplex
	var result bool

	err := json.Unmarshal(data, &apiresponse)
	if err != nil && verboseMode() {
		fmt.Println("Check if open error:", err)
	}
	for _, value := range apiresponse[:] {
		result = value.IsOpen
	}
	return result
}

func candleRunner() string {
	trigger := threshold()
	var cSell int
	var cBuy int

	for index := 0; index < trigger; index++ {
		result := logicSystem()
		if result == "Buy" {
			cBuy++
		} else if result == "Sell" {
			cSell++
		}
	}
	if verboseMode() {
		fmt.Println("Buy orders:", cBuy, "Sell orders: ", cSell)
	}
	return createOrder(cBuy, cSell)
}

func createOrder(cBuy, cSell int) string {
	var typeOrder string

	for {
		if cBuy > cSell {
			setLeverge()
			makeOrder("Buy")
			typeOrder = "Buy"

			fmt.Println(orderCreatedMsg(typeOrder))
			telegramSend(orderCreatedMsg(typeOrder))
			break
		} else if cSell > cBuy {
			setLeverge()
			makeOrder("Sell")
			typeOrder = "Sell"

			fmt.Println(orderCreatedMsg(typeOrder))
			telegramSend(orderCreatedMsg(typeOrder))
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
			fmt.Println(orderDoneMsg())
			telegramSend(orderDoneMsg())
			return true
		}
		time.Sleep(time.Duration(speed) * time.Second)
	}
}

func closePositionProfitBuy() bool {
	speed := speed()
	position := getPosition()

	for {
		if closePositionBuy(position) {
			fmt.Println(ordertriggerMsg())
			telegramSend(ordertriggerMsg())

			closePosition()
			return true
		}
		time.Sleep(time.Duration(speed) * time.Second)
	}
}

func closePositionProfitSell() bool {
	speed := speed()
	position := getPosition()

	for {
		if closePositionSell(position) {
			fmt.Println(ordertriggerMsg())
			telegramSend(ordertriggerMsg())

			closePosition()
			return true
		}
		time.Sleep(time.Duration(speed) * time.Second)
	}
}

func getProfit() bool {
	speed := speed()
	fmt.Println(orderWaintMsg())
	telegramSend(orderWaintMsg())

	for {
		if statusOrder() == false {
			fmt.Println(profitMsg())
			telegramSend(profitMsg())
			time.Sleep(time.Duration(speed+50) * time.Second)
			return true
		}
		time.Sleep(time.Duration(speed) * time.Second)
	}
}

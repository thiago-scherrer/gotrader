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
	ID            int64   `json:"id"`
	LastPrice     float64 `json:"lastPrice"`
	Side          string  `json:"side"`
	Size          int     `json:"size"`
	Price         float64 `json:"price"`
	ChannelID     int     `json:"channelID"`
	OrderID       string  `json:"orderID"`
	OrderQty      int     `json:"orderQty"`
	AvgEntryPrice float64 `json:"avgEntryPrice"`
}

// BotData are json send to the API
type BotData struct {
	ChannelID string `json:"channelID"`
	Message   string `json:"message"`
}

// Conf instruction are the file yaml on disc
type Conf struct {
	Asset     string `yaml:"asset"`
	Candle    string `yaml:"candle"`
	Endpoint  string `yaml:"endpoint"`
	Hand      string `yaml:"hand"`
	Logic     string `yaml:"logic"`
	Speed     string `yaml:"speed"`
	Secret    string `yaml:"secret"`
	Threshold string `yaml:"threshold"`
	Userid    string `yaml:"userid"`
	Profit    string `yaml:"profit"`
}

// KeyConfig struc from yaml file result
type KeyConfig struct {
	result string
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
	var keyconfig KeyConfig
	keyconfig.result = conf.Userid
	return keyconfig.result
}

func secret() string {
	conf := configReader()
	var keyconfig KeyConfig
	keyconfig.result = conf.Secret
	return keyconfig.result
}

func endpoint() string {
	conf := configReader()
	var keyconfig KeyConfig
	keyconfig.result = conf.Endpoint
	return keyconfig.result
}

func hand() string {
	conf := configReader()
	var keyconfig KeyConfig
	keyconfig.result = conf.Hand
	return keyconfig.result
}

func speed() string {
	conf := configReader()
	var keyconfig KeyConfig
	keyconfig.result = conf.Speed
	return keyconfig.result
}

func logic() string {
	conf := configReader()
	var keyconfig KeyConfig
	keyconfig.result = conf.Logic
	return keyconfig.result
}

func asset() string {
	conf := configReader()
	var keyconfig KeyConfig
	keyconfig.result = conf.Asset
	return keyconfig.result
}

func candle() string {
	conf := configReader()
	var keyconfig KeyConfig
	keyconfig.result = conf.Candle
	return keyconfig.result
}

func profit() string {
	conf := configReader()
	var keyconfig KeyConfig
	keyconfig.result = conf.Profit
	return keyconfig.result
}

func threshold() string {
	conf := configReader()
	var keyconfig KeyConfig
	keyconfig.result = conf.Threshold
	return keyconfig.result
}

func handRoll(getParser, hand int) int {
	return (getParser * hand) / 100
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
	hand := StringToIntBit(hand())
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
	return (price() + ((price() / 100) * 3.0)) > getPosition()
}

func closePositionSell() bool {
	return (price() + ((price() / 100) * 3.0)) < getPosition()
}

func closePosition() string {
	asset()
	path := "/api/v1/order"
	requestTipe := "POST"
	data := StringToBytes("symbol=" + asset() +
		"&execInst=Close" + "&price=" + FloatToString(price()) + "&ordType=Limit")

	getResult := clientRobot(requestTipe, path, data)

	return BytesToString(getResult)
}

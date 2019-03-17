package central

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gotrader/internal/convert"
	"gotrader/internal/display"
	"io/ioutil"
	"net/http"
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
	Leverage      string  `json:"leverage,float64,omitempty"`
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
	Hand            int64   `yaml:"hand"`
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

// InitFlag verify if config file has found
func InitFlag() string {
	var config string
	if len(os.Args[1:]) == 0 {
		panic(display.UsageMsg())
	}
	if os.Args[1] == "config" {
		config = os.Args[2]
	} else {
		panic(display.UsageMsg())
	}
	return config
}

// ConfigReader ascsacacscac
func configReader() *Conf {
	confFile := InitFlag()
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

// Asset set the contract type to trade
func Asset() string {
	conf := configReader()
	return conf.Asset
}

// Candle return the time of candle setting
func Candle() int {
	conf := configReader()
	return conf.Candle * 6
}

// Depth get how many ordersbooks can see
func Depth() int64 {
	conf := configReader()
	return conf.Depth
}

func endpoint() string {
	conf := configReader()
	return conf.Endpoint
}

func hand() int64 {
	conf := configReader()
	return conf.Hand
}

func leverage() string {
	conf := configReader()
	return conf.Leverage
}

func profit() float64 {
	conf := configReader()
	return conf.Profit
}

func secret() string {
	conf := configReader()
	return conf.Secret
}

// Threshold return the the value from config file
func Threshold() int {
	conf := configReader()
	return conf.Threshold
}

func userid() string {
	conf := configReader()
	return conf.Userid
}

func telegramUse() bool {
	conf := configReader()
	return conf.TelegramUse
}

func telegramKey() string {
	conf := configReader()
	return conf.TelegramKey
}

func telegramurl() string {
	conf := configReader()
	return conf.TelegramURL
}

func telegramChannel() string {
	conf := configReader()
	return conf.TelegramChannel
}

// VerboseMode if true show debbug mode
func VerboseMode() bool {
	conf := configReader()
	return conf.Verbose
}

// Speed set the daemon daemon, dont change
func Speed() int {
	return 10
}

// ClientRobot are the curl from the bot
func ClientRobot(requestType, path string, data []byte) []byte {
	client := &http.Client{}
	endpoint := endpoint()
	secretQuery := secret()
	userIDquery := userid()
	expire := convert.IntToString((timeExpired()))
	hexResult := hexCreator(secretQuery, requestType, path, expire, convert.BytesToString(data))

	url := endpoint + path

	request, err := http.NewRequest(requestType, url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error create a request on bitmex, got: ", err)
	}

	request.Header.Set("api-signature", hexResult)
	request.Header.Set("api-expires", expire)
	request.Header.Set("api-key", userIDquery)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "gotrader-r0b0tnull")

	for {
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("Error to send the request to the API bitmex, got: ", err)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		if VerboseMode() {
			fmt.Println("Bitmex API Status code are: ", response.StatusCode)
		}

		if response.StatusCode == 200 {
			return body
		}
		Speed := Speed()
		time.Sleep(time.Duration(Speed) * time.Second)
	}
}

// TelegramSend send a msg to the user on settings
func TelegramSend(msg string) int {
	if telegramUse() == false {
		return 200
	}

	client := &http.Client{}
	telegramurl := telegramurl()
	telegramChannel := telegramChannel()
	token := telegramKey()
	data := convert.StringToBytes("chat_id=" + telegramChannel + "&text=" + msg)
	url := telegramurl + "/bot" + token + "/sendMessage"

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error create a request on telegram, got: ", err)
	}

	request.Header.Set("User-Agent", "gotrader-r0b0tnull")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil && VerboseMode() {
		fmt.Println("Error to get body from Telegram API, got", err)
	}

	if VerboseMode() {
		fmt.Println("Telegram API Status code are: ", response.StatusCode)
	}

	return response.StatusCode
}

func hexCreator(secret, requestTipe, path, expired, data string) string {
	concat := requestTipe + path + expired + data
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(concat))
	hexResult := hex.EncodeToString(h.Sum(nil))
	return hexResult
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

func parserAmount(data []byte) int {
	apiresponse := APIResponseComplex{}
	err := json.Unmarshal(data, &apiresponse)
	if err != nil && VerboseMode() {
		fmt.Println("Error to get Amount: ", err)
	}
	return apiresponse.Amount
}

func lastPrice(data []byte) float64 {
	var apiresponse []APIResponseComplex
	var result float64

	err := json.Unmarshal(data, &apiresponse)
	if err != nil {
		fmt.Println("Error to get last price: ", err)
	}
	for _, value := range apiresponse[:] {
		result = value.LastPrice
	}
	return result
}

func leverageResult(data []byte) string {
	apiresponse := APIResponseComplex{}
	err := json.Unmarshal(data, &apiresponse)
	if err != nil {
		fmt.Println("Error to set leverage: ", err)
	}
	return apiresponse.Leverage
}

func makeOrder(orderType string) string {
	Speed := Speed()
	apiresponse := APIResponseComplex{}
	qtyOrerFloat := convert.IntToString(hand())
	Asset()
	path := "/api/v1/order"
	requestTipe := "POST"

	if VerboseMode() {
		fmt.Println("DATA make order: " + "symbol=" + Asset() + "&side=" +
			orderType + "&orderQty=" + qtyOrerFloat + "&price=" +
			convert.FloatToString(price()) + "&ordType=Limit")
	}

	data := convert.StringToBytes("symbol=" + Asset() + "&side=" + orderType + "&orderQty=" +
		qtyOrerFloat + "&price=" + convert.FloatToString(price()) + "&ordType=Limit")

	for {
		getResult := ClientRobot(requestTipe, path, data)
		err := json.Unmarshal(getResult, &apiresponse)
		if err != nil && VerboseMode() {
			fmt.Println("Error to make a order:", err)
			time.Sleep(time.Duration(Speed) * time.Second)
		} else {
			return apiresponse.OrderID
		}
	}
}

func getPosition() float64 {
	var apiresponse []APIResponseComplex
	var result float64
	path := "/api/v1/position" + `?filter={"symbol":"` + Asset() + `"}&count=1`
	requestTipe := "GET"
	data := convert.StringToBytes("message=GoTrader bot&channelID=1")
	getResult := ClientRobot(requestTipe, path, data)
	if VerboseMode() {
		fmt.Println("Data get position" + convert.BytesToString(getResult))
	}
	err := json.Unmarshal(getResult, &apiresponse)
	if err != nil && VerboseMode() {
		fmt.Println("Error to get position:", err)
	}

	for _, value := range apiresponse[:] {
		result = value.AvgEntryPrice
	}
	return result
}

func price() float64 {
	asset := Asset()
	path := "/api/v1/instrument?symbol=" + asset + "&count=100&reverse=false&columns=lastPrice"
	data := convert.StringToBytes("message=GoTrader bot&channelID=1")
	getResult := ClientRobot("GET", path, data)

	return lastPrice(getResult)
}

func closePositionBuy(position float64) bool {
	if VerboseMode() {
		fmt.Println("Close Position Buy: ", position+
			(position/100)*profit())
	}
	return price() >= (position + ((position / 100) * profit()))
}

func closePositionSell(position float64) bool {
	if VerboseMode() {
		fmt.Println("Close Position Sell: ", position-
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

	if VerboseMode() {
		fmt.Println("Data close position: " + "symbol=" + Asset() +
			"&execInst=Close" + "&price=" + priceClose + "&ordType=Limit")
	}

	data := convert.StringToBytes("symbol=" + Asset() +
		"&execInst=Close" + "&price=" + priceClose + "&ordType=Limit")
	getResult := ClientRobot(requestTipe, path, data)
	return convert.BytesToString(getResult)
}

func setLeverge() {
	Asset()
	path := "/api/v1/position/leverage"
	requestTipe := "POST"
	leverage := leverage()
	if VerboseMode() {
		fmt.Println("Data leverge: " + "symbol=" + Asset() +
			"&leverage=" + leverage)
	}

	data := convert.StringToBytes("symbol=" + Asset() + "&leverage=" + leverage)

	ClientRobot(requestTipe, path, data)
	fmt.Println(display.SetlavarageMsg(Asset(), leverage))
	TelegramSend(display.SetlavarageMsg(Asset(), leverage))

}

func statusOrder() bool {
	path := "/api/v1/position" + `?filter={"symbol":"` + Asset() + `"}&count=1`

	data := convert.StringToBytes("message=GoTrader bot&channelID=1")
	getResult := ClientRobot("GET", path, data)
	if VerboseMode() {
		fmt.Println("Data status order: " + convert.BytesToString(getResult))
	}
	return opening(getResult)
}

func opening(data []byte) bool {
	var apiresponse []APIResponseComplex
	var result bool

	err := json.Unmarshal(data, &apiresponse)
	if err != nil && VerboseMode() {
		fmt.Println("Check if open error:", err)
	}
	for _, value := range apiresponse[:] {
		result = value.IsOpen
	}
	return result
}

// CreateOrder create the order on bitmex
func CreateOrder(typeOrder string) {
	Speed := Speed()

	for {
		setLeverge()
		makeOrder("Buy")
		if waitCreateOrder() {
			fmt.Println(display.OrderCreatedMsg(Asset(), typeOrder))
			TelegramSend(display.OrderCreatedMsg(Asset(), typeOrder))
			break
		}
		time.Sleep(time.Duration(Speed) * time.Second)

	}
}

func waitCreateOrder() bool {
	Speed := Speed()

	for {
		if statusOrder() == true {
			fmt.Println(display.OrderDoneMsg(Asset()))
			TelegramSend(display.OrderDoneMsg(Asset()))
			return true
		}
		time.Sleep(time.Duration(Speed) * time.Second)
	}
}

// ClosePositionProfitBuy the Buy position
func ClosePositionProfitBuy() bool {
	Speed := Speed()
	position := getPosition()

	for {
		if closePositionBuy(position) {
			fmt.Println(display.OrdertriggerMsg(Asset()))
			TelegramSend(display.OrdertriggerMsg(Asset()))

			closePosition()
			return true
		}
		time.Sleep(time.Duration(Speed) * time.Second)
	}
}

// ClosePositionProfitSell the Buy position
func ClosePositionProfitSell() bool {
	Speed := Speed()
	position := getPosition()

	for {
		if closePositionSell(position) {
			fmt.Println(display.OrdertriggerMsg(Asset()))
			TelegramSend(display.OrdertriggerMsg(Asset()))

			closePosition()
			return true
		}
		time.Sleep(time.Duration(Speed) * time.Second)
	}
}

// GetProfit waint to start a new trade round
func GetProfit() bool {
	Speed := Speed()
	fmt.Println(display.OrderWaintMsg(Asset()))
	TelegramSend(display.OrderWaintMsg(Asset()))

	for {
		if statusOrder() == false {
			fmt.Println(display.ProfitMsg(Asset()))
			TelegramSend(display.ProfitMsg(Asset()))
			time.Sleep(time.Duration(Speed+50) * time.Second)
			return true
		}
		time.Sleep(time.Duration(Speed) * time.Second)
	}
}

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
	Amount                int           `json:"amount"`
	Symbol                string        `json:"symbol"`
	ID                    int64         `json:"id"`
	LastPrice             float64       `json:"lastPrice"`
	Side                  string        `json:"side"`
	Size                  int           `json:"size"`
	Price                 float64       `json:"price"`
	Account               int           `json:"account"`
	Currency              string        `json:"currency"`
	PrevDeposited         int           `json:"prevDeposited"`
	PrevWithdrawn         int           `json:"prevWithdrawn"`
	PrevTransferIn        int           `json:"prevTransferIn"`
	PrevTransferOut       int           `json:"prevTransferOut"`
	PrevAmount            int           `json:"prevAmount"`
	PrevTimestamp         time.Time     `json:"prevTimestamp"`
	DeltaDeposited        int           `json:"deltaDeposited"`
	DeltaWithdrawn        int           `json:"deltaWithdrawn"`
	DeltaTransferIn       int           `json:"deltaTransferIn"`
	DeltaTransferOut      int           `json:"deltaTransferOut"`
	DeltaAmount           int           `json:"deltaAmount"`
	Deposited             int           `json:"deposited"`
	Withdrawn             int           `json:"withdrawn"`
	TransferIn            int           `json:"transferIn"`
	TransferOut           int           `json:"transferOut"`
	PendingCredit         int           `json:"pendingCredit"`
	PendingDebit          int           `json:"pendingDebit"`
	ConfirmedDebit        int           `json:"confirmedDebit"`
	Timestamp             time.Time     `json:"timestamp"`
	Addr                  string        `json:"addr"`
	Script                string        `json:"script"`
	WithdrawalLock        []interface{} `json:"withdrawalLock"`
	Date                  time.Time     `json:"date"`
	User                  string        `json:"user"`
	Message               string        `json:"message"`
	HTML                  string        `json:"html"`
	FromBot               bool          `json:"fromBot"`
	ChannelID             int           `json:"channelID"`
	OrderID               string        `json:"orderID"`
	ClOrdID               string        `json:"clOrdID"`
	ClOrdLinkID           string        `json:"clOrdLinkID"`
	SimpleOrderQty        interface{}   `json:"simpleOrderQty"`
	OrderQty              int           `json:"orderQty"`
	DisplayQty            interface{}   `json:"displayQty"`
	StopPx                interface{}   `json:"stopPx"`
	PegOffsetValue        interface{}   `json:"pegOffsetValue"`
	PegPriceType          string        `json:"pegPriceType"`
	SettlCurrency         string        `json:"settlCurrency"`
	OrdType               string        `json:"ordType"`
	TimeInForce           string        `json:"timeInForce"`
	ExecInst              string        `json:"execInst"`
	ContingencyType       string        `json:"contingencyType"`
	ExDestination         string        `json:"exDestination"`
	OrdStatus             string        `json:"ordStatus"`
	Triggered             string        `json:"triggered"`
	WorkingIndicator      bool          `json:"workingIndicator"`
	OrdRejReason          string        `json:"ordRejReason"`
	SimpleLeavesQty       interface{}   `json:"simpleLeavesQty"`
	LeavesQty             int           `json:"leavesQty"`
	SimpleCumQty          interface{}   `json:"simpleCumQty"`
	CumQty                int           `json:"cumQty"`
	AvgPx                 int           `json:"avgPx"`
	MultiLegReportingType string        `json:"multiLegReportingType"`
	Text                  string        `json:"text"`
	TransactTime          time.Time     `json:"transactTime"`
	Underlying            string        `json:"underlying"`
	QuoteCurrency         string        `json:"quoteCurrency"`
	Commission            float64       `json:"commission"`
	InitMarginReq         float64       `json:"initMarginReq"`
	MaintMarginReq        float64       `json:"maintMarginReq"`
	RiskLimit             int64         `json:"riskLimit"`
	Leverage              int           `json:"leverage"`
	CrossMargin           bool          `json:"crossMargin"`
	DeleveragePercentile  int           `json:"deleveragePercentile"`
	RebalancedPnl         int           `json:"rebalancedPnl"`
	PrevRealisedPnl       int           `json:"prevRealisedPnl"`
	PrevUnrealisedPnl     int           `json:"prevUnrealisedPnl"`
	PrevClosePrice        float64       `json:"prevClosePrice"`
	OpeningTimestamp      time.Time     `json:"openingTimestamp"`
	OpeningQty            int           `json:"openingQty"`
	OpeningCost           int           `json:"openingCost"`
	OpeningComm           int           `json:"openingComm"`
	OpenOrderBuyQty       int           `json:"openOrderBuyQty"`
	OpenOrderBuyCost      int           `json:"openOrderBuyCost"`
	OpenOrderBuyPremium   int           `json:"openOrderBuyPremium"`
	OpenOrderSellQty      int           `json:"openOrderSellQty"`
	OpenOrderSellCost     int           `json:"openOrderSellCost"`
	OpenOrderSellPremium  int           `json:"openOrderSellPremium"`
	ExecBuyQty            int           `json:"execBuyQty"`
	ExecBuyCost           int           `json:"execBuyCost"`
	ExecSellQty           int           `json:"execSellQty"`
	ExecSellCost          int           `json:"execSellCost"`
	ExecQty               int           `json:"execQty"`
	ExecCost              int           `json:"execCost"`
	ExecComm              int           `json:"execComm"`
	CurrentTimestamp      time.Time     `json:"currentTimestamp"`
	CurrentQty            int           `json:"currentQty"`
	CurrentCost           int           `json:"currentCost"`
	CurrentComm           int           `json:"currentComm"`
	RealisedCost          int           `json:"realisedCost"`
	UnrealisedCost        int           `json:"unrealisedCost"`
	GrossOpenCost         int           `json:"grossOpenCost"`
	GrossOpenPremium      int           `json:"grossOpenPremium"`
	GrossExecCost         int           `json:"grossExecCost"`
	IsOpen                bool          `json:"isOpen"`
	MarkPrice             float64       `json:"markPrice"`
	MarkValue             int           `json:"markValue"`
	RiskValue             int           `json:"riskValue"`
	HomeNotional          float64       `json:"homeNotional"`
	ForeignNotional       int           `json:"foreignNotional"`
	PosState              string        `json:"posState"`
	PosCost               int           `json:"posCost"`
	PosCost2              int           `json:"posCost2"`
	PosCross              int           `json:"posCross"`
	PosInit               int           `json:"posInit"`
	PosComm               int           `json:"posComm"`
	PosLoss               int           `json:"posLoss"`
	PosMargin             int           `json:"posMargin"`
	PosMaint              int           `json:"posMaint"`
	PosAllowance          int           `json:"posAllowance"`
	TaxableMargin         int           `json:"taxableMargin"`
	InitMargin            int           `json:"initMargin"`
	MaintMargin           int           `json:"maintMargin"`
	SessionMargin         int           `json:"sessionMargin"`
	TargetExcessMargin    int           `json:"targetExcessMargin"`
	VarMargin             int           `json:"varMargin"`
	RealisedGrossPnl      int           `json:"realisedGrossPnl"`
	RealisedTax           int           `json:"realisedTax"`
	RealisedPnl           int           `json:"realisedPnl"`
	UnrealisedGrossPnl    int           `json:"unrealisedGrossPnl"`
	LongBankrupt          int           `json:"longBankrupt"`
	ShortBankrupt         int           `json:"shortBankrupt"`
	TaxBase               int           `json:"taxBase"`
	IndicativeTaxRate     int           `json:"indicativeTaxRate"`
	IndicativeTax         int           `json:"indicativeTax"`
	UnrealisedTax         int           `json:"unrealisedTax"`
	UnrealisedPnl         int           `json:"unrealisedPnl"`
	UnrealisedPnlPcnt     float64       `json:"unrealisedPnlPcnt"`
	UnrealisedRoePcnt     float64       `json:"unrealisedRoePcnt"`
	SimpleQty             interface{}   `json:"simpleQty"`
	SimpleCost            interface{}   `json:"simpleCost"`
	SimpleValue           interface{}   `json:"simpleValue"`
	SimplePnl             interface{}   `json:"simplePnl"`
	SimplePnlPcnt         interface{}   `json:"simplePnlPcnt"`
	AvgCostPrice          int           `json:"avgCostPrice"`
	AvgEntryPrice         float64       `json:"avgEntryPrice"`
	BreakEvenPrice        float64       `json:"breakEvenPrice"`
	MarginCallPrice       int           `json:"marginCallPrice"`
	LiquidationPrice      int           `json:"liquidationPrice"`
	BankruptPrice         int           `json:"bankruptPrice"`
	LastValue             int           `json:"lastValue"`
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
		"&execInst=" + "Close" + "&price=" + FloatToString(price()) + "&ordType=Limit")

	getResult := clientRobot(requestTipe, path, data)

	return BytesToString(getResult)
}

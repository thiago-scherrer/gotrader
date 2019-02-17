package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// APIResponseComplex used to struct data from API response,
// thanks https://mholt.github.io/json-to-go/
type APIResponseComplex struct {
	Amount           int           `json:"amount"`
	Symbol           string        `json:"symbol"`
	ID               int64         `json:"id"`
	LastPrice        float64       `json:"lastPrice"`
	Side             string        `json:"side"`
	Size             int           `json:"size"`
	Price            float64       `json:"price"`
	Account          int           `json:"account"`
	Currency         string        `json:"currency"`
	PrevDeposited    int           `json:"prevDeposited"`
	PrevWithdrawn    int           `json:"prevWithdrawn"`
	PrevTransferIn   int           `json:"prevTransferIn"`
	PrevTransferOut  int           `json:"prevTransferOut"`
	PrevAmount       int           `json:"prevAmount"`
	PrevTimestamp    time.Time     `json:"prevTimestamp"`
	DeltaDeposited   int           `json:"deltaDeposited"`
	DeltaWithdrawn   int           `json:"deltaWithdrawn"`
	DeltaTransferIn  int           `json:"deltaTransferIn"`
	DeltaTransferOut int           `json:"deltaTransferOut"`
	DeltaAmount      int           `json:"deltaAmount"`
	Deposited        int           `json:"deposited"`
	Withdrawn        int           `json:"withdrawn"`
	TransferIn       int           `json:"transferIn"`
	TransferOut      int           `json:"transferOut"`
	PendingCredit    int           `json:"pendingCredit"`
	PendingDebit     int           `json:"pendingDebit"`
	ConfirmedDebit   int           `json:"confirmedDebit"`
	Timestamp        time.Time     `json:"timestamp"`
	Addr             string        `json:"addr"`
	Script           string        `json:"script"`
	WithdrawalLock   []interface{} `json:"withdrawalLock"`
	Date             time.Time     `json:"date"`
	User             string        `json:"user"`
	Message          string        `json:"message"`
	HTML             string        `json:"html"`
	FromBot          bool          `json:"fromBot"`
	ChannelID        int           `json:"channelID"`
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
		panic(err)
	}

	return apiresponse.Amount
}

func lastPrice(data []byte) float64 {
	var apiresponse []APIResponseComplex
	var result float64

	err := json.Unmarshal(data, &apiresponse)
	if err != nil {
		panic(err)
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

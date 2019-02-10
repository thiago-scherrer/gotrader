package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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

func configFile() string {
	var config string
	var help string

	flag.StringVar(&config, "config", "", "")
	flag.StringVar(&help, "help", "off", "")

	flag.Parse()

	if len(config) < 1 {
		config = "usage: ./gotrade -config config.yml"
		return config
	}
	return config

}

func configReader(keyname, confFile string) string {
	conf := Conf{}
	var keyconfig KeyConfig

	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic(err)
	}

	switch keyname {
	case "userid":
		keyconfig.result = conf.Userid
	case "secret":
		keyconfig.result = conf.Secret
	case "api":
		keyconfig.result = conf.Endpoint
	case "hand":
		keyconfig.result = conf.Hand
	case "speed":
		keyconfig.result = conf.Speed
	case "logic":
		keyconfig.result = conf.Logic
	case "asset":
		keyconfig.result = conf.Asset
	case "candle":
		keyconfig.result = conf.Candle
	case "threshold":
		keyconfig.result = conf.Threshold
	}

	return keyconfig.result
}

func handRoll(getParser, hand int) int {
	return (getParser * hand) / 100
}

func hexCreator(secret, requestTipe, path, expired string) string {
	concat := requestTipe + path + expired
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(concat))

	hexResult := hex.EncodeToString(h.Sum(nil))
	return hexResult
}

func parserAmount(data []byte) int {
	var apiresponse []APIResponseComplex
	var result int
	fmt.Println(BytesToString(data))

	err := json.Unmarshal(data, &apiresponse)
	if err != nil {
		panic(err)
	}

	for _, value := range apiresponse[:] {
		result = value.Amount
	}
	return result
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

func requiredConfig(confFile string) bool {
	var result bool

	userid := configReader("userid", confFile)
	secret := configReader("secret", confFile)
	endpoint := configReader("api", confFile)
	hand := configReader("hand", confFile)
	speed := configReader("speed", confFile)
	logic := configReader("logic", confFile)
	asset := configReader("asset", confFile)
	candle := configReader("candle", confFile)
	threshold := configReader("threshold", confFile)

	if len(userid) == 0 {
		result = true
		panic("user id not found!")
	} else if len(secret) == 0 {
		result = true
		panic("secret not found!")
	} else if len(endpoint) == 0 {
		result = true
		panic("api endpoint not found!")
	} else if len(hand) == 0 {
		result = true
		panic("hand not found!")
	} else if len(speed) == 0 {
		result = true
		panic("speed not found!")
	} else if len(logic) == 0 {
		result = true
		panic("logic not found!")
	} else if len(asset) == 0 {
		result = true
		panic("asset not found!")
	} else if len(candle) == 0 {
		result = true
		panic("candle time not found!")
	} else if len(threshold) == 0 {
		result = true
		panic("threshold not found!")
	}
	return result
}

package reader

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/go-redis/redis"
	"github.com/thiago-scherrer/gotrader/internal/convert"
	"gopkg.in/yaml.v2"
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
	Trades        int64   `json:"trades"`
}

// Conf instruction are the file yaml on disc
type Conf struct {
	Asset     string  `yaml:"asset"`
	Candle    int     `yaml:"candle"`
	Endpoint  string  `yaml:"endpoint"`
	Hand      int64   `yaml:"hand"`
	Leverage  string  `yaml:"leverage"`
	Profit    float64 `yaml:"profit"`
	Secret    string  `yaml:"secret"`
	StopLoss  float64 `yaml:"stoploss"`
	Threshold int     `yaml:"threshold"`
	Userid    string  `yaml:"userid"`
}

// Use to get the right time of the candle time
const fixtime int = 6

// ConfigPath verify where is config file
func ConfigPath() string {
	var config string

	if os.Getenv("GOTRADER_MODE") == "prod" {
		config = "/opt/config.yml"
	} else if os.Getenv("GOTRADER_MODE") == "testnet" {
		config = "./opt/config-testnet.yml"
	} else {
		config = "../../configs/config-test.yml"
	}

	return config
}

// ConfigReader - read the file from PC
func configReader() *Conf {
	confFile := ConfigPath()
	conf := Conf{}

	var once sync.Once

	onceReader := func() {
		config, _ := ioutil.ReadFile(confFile)
		yaml.Unmarshal(config, &conf)
	}

	once.Do(onceReader)
	return &conf

}

// RDclient create a client to the redis container
func RDclient() *redis.Client {

	redisurl := os.Getenv("REDISURL")

	if redisurl == "" {
		redisurl = "redis:6379"
	}

	return redis.NewClient(&redis.Options{
		Addr:     redisurl,
		Password: "",
		DB:       0,
	})
}

// Boot create the initial Bootstrap config
func Boot() {
	conf := configReader()
	db := RDclient()
	defer db.Close()
	bootStatus, _ := db.Get("reload").Result()

	if bootStatus != "true" {
		db.Set("hand", conf.Hand, 0).Err()
		db.Set("asset", conf.Asset, 0).Err()
		db.Set("candle", conf.Candle, 0).Err()
		db.Set("threshold", conf.Threshold, 0).Err()
		db.Set("leverage", conf.Leverage, 0).Err()
		db.Set("profit", conf.Profit, 0).Err()
		db.Set("stoploss", conf.StopLoss, 0).Err()
	}
}

// Asset set the contract type to trade
func Asset() string {
	db := RDclient()
	defer db.Close()
	result, _ := db.Get("asset").Result()
	return result
}

// Candle return the time of candle setting
func Candle() int64 {
	db := RDclient()
	defer db.Close()
	result, _ := db.Get("candle").Result()
	return convert.StringToInt(result)
}

// Endpoint return url from bitmex
func Endpoint() string {
	conf := configReader()
	return conf.Endpoint
}

// Hand return value to trade
func Hand() string {
	db := RDclient()
	defer db.Close()
	result, _ := db.Get("hand").Result()
	return result
}

// Leverage return the value to set on laverage trading
func Leverage() string {
	db := RDclient()
	defer db.Close()
	result, _ := db.Get("leverage").Result()
	return result
}

// Secret return API password
func Secret() string {
	conf := configReader()
	return conf.Secret
}

// Threshold return the the value from config file
func Threshold() int64 {
	db := RDclient()
	defer db.Close()
	result, _ := db.Get("threshold").Result()
	return convert.StringToInt(result)
}

// Userid return user identify from bitmex
func Userid() string {
	conf := configReader()
	return conf.Userid
}

// Profit return the profit percentage
func Profit() float64 {
	db := RDclient()
	defer db.Close()
	result, _ := db.Get("profit").Result()
	return convert.StringToFloat64(result)
}

// StopLoss return the StopLoss percentage
func StopLoss() float64 {
	db := RDclient()
	defer db.Close()
	result, _ := db.Get("stoploss").Result()
	return convert.StringToFloat64(result)
}

//APISimple return JSON
func APISimple() APIResponseComplex {
	var ar APIResponseComplex
	return ar
}

//APIArray return JSON Array
func APIArray() []APIResponseComplex {
	var ar []APIResponseComplex
	return ar
}

// Data return a default data to the bitmet api
func Data() []byte {
	return convert.StringToBytes("message=GoTrader bot&channelID=1")
}

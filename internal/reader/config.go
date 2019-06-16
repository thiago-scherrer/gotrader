package reader

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/thiago-scherrer/gotrader/internal/display"
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
}

// Use to get the right time of the candle time
const fixtime int = 6

// InitFlag verify if config file has found
func InitFlag() string {
	var config string
	if len(os.Args[1:]) == 0 {
		log.Fatalf(display.UsageMsg())
	}
	if os.Args[1] == "config" {
		config = os.Args[2]
	} else {
		log.Fatalf(display.UsageMsg())
	}
	return config
}

// ConfigReader - read the file from PC
func configReader() *Conf {
	confFile := InitFlag()
	conf := Conf{}
	var once sync.Once

	onceReader := func() {
		config, _ := ioutil.ReadFile(confFile)
		yaml.Unmarshal(config, &conf)
	}
	once.Do(onceReader)
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
	return conf.Candle * fixtime
}

// Depth get how many ordersbooks can see
func Depth() int64 {
	conf := configReader()
	return conf.Depth
}

// Endpoint return url from bitmex
func Endpoint() string {
	conf := configReader()
	return conf.Endpoint
}

// Hand return value to trade
func Hand() int64 {
	conf := configReader()
	return conf.Hand
}

// Leverage return the value to set on laverage trading
func Leverage() string {
	conf := configReader()
	return conf.Leverage
}

// Profit return the porcentage to exit a trade
func Profit() float64 {
	conf := configReader()
	return conf.Profit
}

// Secret return API password
func Secret() string {
	conf := configReader()
	return conf.Secret
}

// Threshold return the the value from config file
func Threshold() int {
	conf := configReader()
	return conf.Threshold
}

// Userid return user identify from bitmex
func Userid() string {
	conf := configReader()
	return conf.Userid
}

// TelegramUse return if enable or not telegram
func TelegramUse() bool {
	conf := configReader()
	return conf.TelegramUse
}

// TelegramKey return API Key from telegram
func TelegramKey() string {
	conf := configReader()
	return conf.TelegramKey
}

// Telegramurl return API endpoint from telegram
func Telegramurl() string {
	conf := configReader()
	return conf.TelegramURL
}

// TelegramChannel return the channel to send a msg
func TelegramChannel() string {
	conf := configReader()
	return conf.TelegramChannel
}

// Speed set the daemon daemon, dont change
func Speed() int {
	return 10
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

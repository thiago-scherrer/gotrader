package reader

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/thiago-scherrer/gotrader/internal/convert"
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
	Trades        int64   `json:"trades"`
}

// Conf instruction are the file yaml on disc
type Conf struct {
	Asset         string  `yaml:"asset"`
	Candle        int     `yaml:"candle"`
	Endpoint      string  `yaml:"endpoint"`
	Hand          int64   `yaml:"hand"`
	Leverage      string  `yaml:"leverage"`
	Profit        float64 `yaml:"profit"`
	Secret        string  `yaml:"secret"`
	Threshold     int     `yaml:"threshold"`
	Userid        string  `yaml:"userid"`
	MatrixUse     bool    `yaml:"matrixuse"`
	MatrixKey     string  `yaml:"matrix_key"`
	MatrixURL     string  `yaml:"matrixurl"`
	MatrixChannel string  `yaml:"matrixchannel"`
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
	return conf.Candle
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

// MatrixUse return if enable or not Matrix
func MatrixUse() bool {
	conf := configReader()
	return conf.MatrixUse
}

// MatrixKey return API Key from Matrix
func MatrixKey() string {
	conf := configReader()
	return conf.MatrixKey
}

// Matrixurl return API endpoint from Matrix
func Matrixurl() string {
	conf := configReader()
	return conf.MatrixURL
}

// MatrixChannel return the channel to send a msg
func MatrixChannel() string {
	conf := configReader()
	return conf.MatrixChannel
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

// Data I can not be leave empty
func Data() []byte {
	return convert.StringToBytes("message=GoTrader bot&channelID=1")
}

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// APIResponse are used to struc json result from API
type APIResponse struct {
	Amount int64 `json:"amount"`
}

// APIResponseComplex used to struct data from API response,
// thanks https://mholt.github.io/json-to-go/
type APIResponseComplex struct {
	Symbol string  `json:"symbol"`
	ID     int64   `json:"id"`
	Side   string  `json:"side"`
	Size   int     `json:"size"`
	Price  float64 `json:"price"`
}

// Conf instruction are the file yaml on disc
type Conf struct {
	Asset    string `yaml:"asset"`
	Endpoint string `yaml:"endpoint"`
	Hand     string `yaml:"hand"`
	Logic    string `yaml:"logic"`
	Speed    string `yaml:"speed"`
	Secret   string `yaml:"secret"`
	Userid   string `yaml:"userid"`
}

// KeyConfig struc from yaml file result
type KeyConfig struct {
	result string
}

func configReader(keyname string, yamlFile []byte) string {
	conf := Conf{}
	var keyconfig KeyConfig

	err := yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic(err)
	}

	if keyname == "userid" {
		keyconfig.result = conf.Userid
	} else if keyname == "secret" {
		keyconfig.result = conf.Secret
	} else if keyname == "api" {
		keyconfig.result = conf.Endpoint
	} else if keyname == "hand" {
		keyconfig.result = conf.Hand
	} else if keyname == "speed" {
		keyconfig.result = conf.Speed
	} else if keyname == "logic" {
		keyconfig.result = conf.Logic
	} else if keyname == "asset" {
		keyconfig.result = conf.Asset
	}
	return keyconfig.result
}

func handRoll(getParser, hand int64) int64 {
	return (getParser * hand) / 100
}

func hexCreator(secret, requestTipe, path, expired string) string {
	concat := requestTipe + path + expired
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(concat))

	hexResult := hex.EncodeToString(h.Sum(nil))
	return hexResult
}

func parserAmount(getResult string) int64 {
	getByte := StringToBytes(getResult)
	var apiresponse APIResponse

	json.Unmarshal(getByte, &apiresponse)

	return apiresponse.Amount
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
	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		panic(err)
	}

	userid := configReader("userid", yamlFile)
	secret := configReader("secret", yamlFile)
	endpoint := configReader("api", yamlFile)
	hand := configReader("hand", yamlFile)
	speed := configReader("speed", yamlFile)
	logic := configReader("logic", yamlFile)
	asset := configReader("asset", yamlFile)

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
	}

	return result
}

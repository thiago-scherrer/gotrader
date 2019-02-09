package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var logicResult string
	var strikeBuy int
	var strikeSell int
	var strike int
	confFile := "config.yml"
	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		panic(err)
	}
	requiredConfig(confFile)
	asset := configReader("asset", yamlFile)
	candle := configReader("candle", yamlFile)
	endpoint := configReader("api", yamlFile)
	lConfig := configReader("logic", yamlFile)
	secretQuery := configReader("secret", yamlFile)
	speed := StringToInt(configReader("speed", yamlFile))
	strike = StringToIntBit(configReader("threshold", yamlFile))
	userIDquery := configReader("userid", yamlFile)
	expire := IntToString((timeExpired()))
	hexResult := hexCreator(secretQuery, "GET", "/api/v1/user/wallet", expire)

	// get $$ from account
	money := parserAmount(clientGet(hexResult,
		endpoint, "/api/v1/user/wallet", expire, userIDquery),
	)

	hand := handRoll(money,
		StringToInt(configReader("hand", yamlFile)),
	)

	for count := 0; count <= strike; count++ {

		switch lConfig {
		case "volume":
			logicResult = volume(userIDquery, secretQuery, endpoint, asset, candle, hand, speed)
		default:
			panic("logic nog found!")
		}
		if logicResult == "Buy" {
			strikeBuy++
		} else if logicResult == "Sell" {
			strikeSell++
		} else {
			count--
		}
		fmt.Println(logicResult)
	}

}

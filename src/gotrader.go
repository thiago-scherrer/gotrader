package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	confFile := "config.yml"
	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		panic(err)
	}
	requiredConfig(confFile)
	asset := configReader("asset", yamlFile)
	endpoint := configReader("api", yamlFile)
	lConfig := configReader("logic", yamlFile)
	secretQuery := configReader("secret", yamlFile)
	speed := StringToInt(configReader("speed", yamlFile))
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

	for {
		expired := IntToString((timeExpired()))

		switch lConfig {
		case "volume":
			result := volume(expired, userIDquery, secretQuery, endpoint, asset, hand)
			fmt.Print(result, "\n")
		default:
			panic("logic nog found!")
		}
		time.Sleep(time.Duration(speed) * time.Second)
	}
}

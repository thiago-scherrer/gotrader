package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

// TimeStamp struct to validate expired time api
type TimeStamp struct {
	timeResult  int64
	timeExpired int64
}

func TestConfigFile(t *testing.T) {
	config := "../../configs/config-test.yml"
	_, err := os.Stat(config)

	if err == nil {
	} else if os.IsNotExist(err) {
		t.Error("config file not found! ", config)
	} else {
		t.Error("error to look at config file.")
	}
}

func TestConfigReader(t *testing.T) {
	useridKey := "userid"
	secretKey := "secret"
	endpointKey := "api"
	confFile := "../../configs/config-test.yml"

	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		panic(err)
	}

	userIDquery := configReader(useridKey, yamlFile)
	if userIDquery != "dontpanic" {
		t.Error("the yml file not return the user, the result are: ", userIDquery)
	}
	secretQuery := configReader(secretKey, yamlFile)
	if secretQuery != "123456" {
		t.Error("the yml file not return the password, the result are: ", secretQuery)
	}
	endpointQuery := configReader(endpointKey, yamlFile)
	if endpointQuery != "https://testnet.bitmex.com" {
		t.Error("the yml file not return the endpoint, the result are: ", endpointQuery)
	}
	if !strings.HasPrefix(endpointQuery, "https://") {
		t.Error("the endpoint in yml file is not not secured (https), the result are: ", endpointQuery)
	}
}

func TestExpiresTime(t *testing.T) {
	var timeStampResult TimeStamp

	now := time.Now()
	timestamp := now.Unix()
	timeStampResult.timeResult = timeStamp()
	diff := timestamp - timeStampResult.timeResult
	timeStampResult.timeExpired = timeExpired()
	timeStampResultExpected := timestamp + 60

	if diff != 0 {
		t.Error("time stamp function noting working, result are: ", diff)
	}
	if timeStampResult.timeExpired != timeStampResultExpected {
		t.Error("expired time not working, resulte are: ", timeStampResult.timeExpired, "expected: ", timeStampResultExpected)
	}
}

func TestHmac(t *testing.T) {
	expired := "1518064236"
	path := "/api/v1/instrument"
	requestTipe := "GET"
	secretQuery := "chNOOS4KvNXR_Xq4k4c9qsfoKWvnDecLATCRlcBwyKDYnWgO"
	hexExpected := "c7682d435d0cfe87c16098df34ef2eb5a549d4c5a3c2b1f0f77b8af73423bf00"
	hexResult := hexCreator(secretQuery, requestTipe, path, expired)

	if hexExpected != hexResult {
		t.Error("GET hex not match, got: ", hexResult, "need: ", hexExpected)
	}
}

func TestGetAnnounement(t *testing.T) {
	confFile := "../../configs/config.yml"
	expired := IntToString((timeExpired()))
	path := "/api/v1/user/affiliateStatus"
	requestTipe := "GET"

	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		panic(err)
	}
	secretQuery := configReader("secret", yamlFile)
	userIDquery := configReader("userid", yamlFile)
	endpoint := configReader("api", yamlFile)

	hexResult := hexCreator(secretQuery, requestTipe, path, expired)

	getResult := clientGet(hexResult, endpoint, path, expired, userIDquery)

	if len(getResult) <= 3 {
		t.Error("GET response not woring, got: ", getResult)
	}
}

func TestGetWalletAmount(t *testing.T) {
	confFile := "../../configs/config.yml"
	expired := IntToString((timeExpired()))
	path := "/api/v1/user/wallet"
	requestTipe := "GET"

	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		panic(err)
	}
	secretQuery := configReader("secret", yamlFile)
	userIDquery := configReader("userid", yamlFile)
	endpoint := configReader("api", yamlFile)
	hexResult := hexCreator(secretQuery, requestTipe, path, expired)

	getResult := clientGet(hexResult, endpoint, path, expired, userIDquery)

	getParser := parserAmount(getResult)

	if getParser <= 1 {
		t.Error("error to get wallet value, got: ", getParser)
	}
}

func TestPostLogout(t *testing.T) {
	confFile := "../../configs/config.yml"
	expired := IntToString((timeExpired()))
	path := "/api/v1/user/logout"
	requestTipe := "POST"

	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		panic(err)
	}
	secretQuery := configReader("secret", yamlFile)
	userIDquery := configReader("userid", yamlFile)
	endpoint := configReader("api", yamlFile)

	hexResult := hexCreator(secretQuery, requestTipe, path, expired)

	postResult := clientPost(hexResult, endpoint, path, expired, userIDquery)

	if postResult != "" {
		t.Error("POST response not woring, got: ", postResult)
	}
}

func TestTradeValue(t *testing.T) {
	confFile := "../../configs/config.yml"
	expired := IntToString((timeExpired()))
	path := "/api/v1/user/wallet"
	requestTipe := "GET"

	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		panic(err)
	}
	secretQuery := configReader("secret", yamlFile)
	userIDquery := configReader("userid", yamlFile)
	endpoint := configReader("api", yamlFile)
	hand := StringToInt(configReader("hand", yamlFile))
	hexResult := hexCreator(secretQuery, requestTipe, path, expired)
	getResult := clientGet(hexResult, endpoint, path, expired, userIDquery)
	getParser := parserAmount(getResult)
	handRollEspected := (getParser * hand) / 100
	result := handRoll(getParser, hand)

	if handRollEspected != result {
		t.Error("the value to trade not working, got: ", result, ", want: ", handRollEspected)
	}
}

func TestQuote(t *testing.T) {
	confFile := "../../configs/config.yml"
	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		panic(err)
	}
	secretQuery := configReader("secret", yamlFile)
	userIDquery := configReader("userid", yamlFile)
	endpoint := configReader("api", yamlFile)
	asset := configReader("asset", yamlFile)
	path := 
	requestTipe := "GET"
	expired := IntToString((timeExpired()))
	hexResult := hexCreator(secretQuery, requestTipe, path, expired)
	getResult := clientGet(hexResult, endpoint, path, expired, userIDquery)

	//result := getQuote()

	fmt.Println(getResult)
}

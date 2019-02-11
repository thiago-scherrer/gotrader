package main

import (
	"encoding/json"
	"fmt"
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

	userIDquery := configReader(useridKey, confFile)
	if userIDquery != "dontpanic" {
		t.Error("the yml file not return the user, the result are: ", userIDquery)
	}
	secretQuery := configReader(secretKey, confFile)
	if secretQuery != "123456" {
		t.Error("the yml file not return the password, the result are: ", secretQuery)
	}
	endpointQuery := configReader(endpointKey, confFile)
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
	path := "/api/v1/user/affiliateStatus"
	requestTipe := "GET"
	data := map[string]string{"message": "TDDRobot =)", "channelID": "1"}
	dataB, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	getResult := clientRobot(requestTipe, confFile, path, dataB)

	if len(getResult) <= 3 {
		t.Error("GET response not woring, got: ", getResult)
	}
}

func TestPostChat(t *testing.T) {
	confFile := "../../configs/config.yml"
	path := "/api/v1/chat"
	requestTipe := "POST"
	data := map[string]string{"message": "TDDRobot =)", "channelID": "1"}
	dataB, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	getResult := clientRobot(requestTipe, confFile, path, dataB)

	//postResult := BytesToString(getResult)

	fmt.Println(getResult)
}

func TestTradeValue(t *testing.T) {
	confFile := "../../configs/config.yml"
	path := "/api/v1/user/wallet"
	requestTipe := "GET"
	hand := StringToIntBit(configReader("hand", confFile))
	data := map[string]string{"message": "TDDRobot =)", "channelID": "1"}
	dataB, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	getResult := clientRobot(requestTipe, confFile, path, dataB)
	getParser := parserAmount(getResult)
	handRollEspected := (getParser * hand) / 100
	result := handRoll(getParser, hand)

	if handRollEspected != result {
		t.Error("the value to trade not working, got: ", result, ", want: ", handRollEspected)
	}
}

func TestConfigLoad(t *testing.T) {
	getResult := configFile()
	expected := "usage: ./gotrade -config config.yml"

	if getResult != expected {
		t.Error("config load not working, got: ", getResult)
	}
}
func TestQuote(t *testing.T) {
	confFile := "../../configs/config.yml"
	asset := configReader("asset", confFile)
	path := "/api/v1/instrument?symbol=" + asset + "&count=100&reverse=false&columns=lastPrice"
	data := map[string]string{"message": "TDDRobot =)", "channelID": "1"}
	dataB, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	getResult := clientRobot("GET", confFile, path, dataB)

	getPrice := lastPrice(getResult)
	if getPrice <= 3 {
		t.Error("erro to get last price, got: ", getPrice)
	}
}

func TestGetWalletAmount(t *testing.T) {
	confFile := "../../configs/config.yml"
	path := "/api/v1/user/wallet"
	requestTipe := "GET"
	data := map[string]string{"message": "TDDRobot =)", "channelID": "1"}
	dataB, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	getResult := clientRobot(requestTipe, confFile, path, dataB)

	getParser := parserAmount(getResult)

	if getParser <= 1 {
		t.Error("error to get wallet value, got: ", getParser)
	}
}

func TestSellOrder(t *testing.T) {

}

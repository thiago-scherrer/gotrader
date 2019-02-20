package main

import (
	"encoding/json"
	"testing"
	"time"
)

// TimeStamp struct to validate expired time api
type TimeStamp struct {
	timeResult  int64
	timeExpired int64
}

func TestExpiresTime(t *testing.T) {
	initFlag()
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
	initFlag()

	expired := "1518064236"
	path := "/api/v1/instrument"
	requestTipe := "GET"
	secretQuery := "chNOOS4KvNXR_Xq4k4c9qsfoKWvnDecLATCRlcBwyKDYnWgO"
	hexExpected := "9c37199dd75f47b63774ddbb5e2851998848d5ec62b9a2bbc380a48f620b305e"
	hexResult := hexCreator(secretQuery, requestTipe, path, expired, "data")

	if hexExpected != hexResult {
		t.Error("GET hex not match, got: ", hexResult, "need: ", hexExpected)
	}
}

func TestGetAnnounement(t *testing.T) {
	initFlag()
	path := "/api/v1/user/affiliateStatus"
	requestTipe := "GET"
	data := StringToBytes("message=GoTrader bot&channelID=1")

	getResult := clientRobot(requestTipe, path, data)

	if len(getResult) <= 3 {
		t.Error("GET response not woring, got: ", getResult)
	}

}

func TestTradeValue(t *testing.T) {
	initFlag()
	path := "/api/v1/user/wallet"
	requestTipe := "GET"
	hand := hand()
	data := StringToBytes("message=GoTrader bot&channelID=1")
	getResult := clientRobot(requestTipe, path, data)
	getParser := parserAmount(getResult)
	handRollEspected := (getParser * hand) / 100
	result := handRoll(getParser, hand)

	if handRollEspected != result {
		t.Error("the value to trade not working, got: ", result, ", want: ", handRollEspected)
	}
}

func TestQuote(t *testing.T) {
	initFlag()
	asset := asset()
	path := "/api/v1/instrument?symbol=" + asset + "&count=100&reverse=false&columns=lastPrice"
	data := StringToBytes("message=GoTrader bot&channelID=1")

	getResult := clientRobot("GET", path, data)

	getPrice := lastPrice(getResult)
	if getPrice <= 3 {
		t.Error("erro to get last price, got: ", getPrice)
	}
}

func TestPostChat(t *testing.T) {
	apiresponse := APIResponseComplex{}

	initFlag()
	path := "/api/v1/chat"
	requestTipe := "POST"
	data := StringToBytes("message=GoTrader bot&channelID=1")
	getResult := clientRobot(requestTipe, path, data)

	err := json.Unmarshal(getResult, &apiresponse)
	if err != nil {
		panic(err)
	}

	postResult := BytesToString(getResult)
	if apiresponse.ChannelID != 1 {
		t.Error("error to use chat, got: ", postResult)
	}
}

func TestGetWalletAmount(t *testing.T) {
	initFlag()
	path := "/api/v1/user/wallet"
	requestTipe := "GET"
	data := StringToBytes("message=GoTrader bot&channelID=1")
	getResult := clientRobot(requestTipe, path, data)

	getParser := parserAmount(getResult)

	if getParser <= 1 {
		t.Error("error to get wallet value, got: ", getParser)
	}
}

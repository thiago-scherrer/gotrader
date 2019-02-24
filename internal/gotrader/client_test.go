package main

import (
	"testing"
)

// TimeStamp struct to validate expired time api
type TimeStamp struct {
	timeResult  int64
	timeExpired int64
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

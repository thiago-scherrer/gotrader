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

package main

import (
	"testing"
	"time"
)

func TestFlag(t *testing.T) {
	getResult := initFlag()
	if len(getResult) <= 1 {
		t.Error("init flag not working, got: ", getResult)
	}
}

func TestReader(t *testing.T) {
	getResult := configReader()

	if getResult.Asset != "XBTUSD" {
		t.Error("error to read config file, got:", getResult)
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

func TestParserAmount(t *testing.T) {
	mock := `{ "Amount": 10 }`
	getResult := parserAmount(StringToBytes(mock))

	if getResult != 10 {
		t.Error("json parser not working, got:", getResult)
	}
}

func TestLastPriceJson(t *testing.T) {
	mock := `[{ "LastPrice": 10.1 }]`
	getResult := lastPrice(StringToBytes(mock))

	if getResult != 10.1 {
		t.Error("LastPrice json parser not working, got:", getResult)
	}
}

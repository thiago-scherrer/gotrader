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

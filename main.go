package main

import (
	"fmt"

	"github.com/thiago-scherrer/gotrader/central"
	"github.com/thiago-scherrer/gotrader/display"
	"github.com/thiago-scherrer/gotrader/logic"
)

func main() {
	for {
		daemonize()
	}
}

func daemonize() {
	central.InitFlag()
	fmt.Println(
		display.HelloMsg(central.Asset()),
	)

	central.TelegramSend(
		display.HelloMsg(central.Asset()),
	)

	typeOrder := logic.CandleRunner()
	central.CreateOrder(typeOrder)

	if typeOrder == "Buy" {
		central.ClosePositionProfitBuy()
	} else {
		central.ClosePositionProfitSell()
	}
	central.GetProfit()
}

package main

import (
	"log"

	"github.com/thiago-scherrer/gotrader/internal/central"
	"github.com/thiago-scherrer/gotrader/internal/display"
	"github.com/thiago-scherrer/gotrader/internal/logic"
)

func main() {
	for {
		daemonize()
	}
}

func daemonize() {
	central.InitFlag()
	log.Println(
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

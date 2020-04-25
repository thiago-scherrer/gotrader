package main

import (
	"log"

	"github.com/thiago-scherrer/gotrader/internal/central"
	"github.com/thiago-scherrer/gotrader/internal/display"
	"github.com/thiago-scherrer/gotrader/internal/logic"
	"github.com/thiago-scherrer/gotrader/internal/reader"
)

func main() {
	for {
		reader.Boot()
		daemonize()
	}
}

func daemonize() {
	reader.ConfigPath()

	log.Println(
		display.HelloMsg(reader.Asset()),
	)

	trd := logic.CandleRunner()
	hand := reader.Hand()

	if central.CreateOrder(trd, hand) == false {
		trd = "Error"
	}

	log.Println(
		display.OrderPrice(
			reader.Asset(),
			central.GetPosition(),
		),
	)

	if trd == "Buy" {
		logic.ClosePositionProfitBuy()
	} else if trd == "Sell" {
		logic.ClosePositionProfitSell()
	} else {
		display.OrderCancelMsg()
	}
}

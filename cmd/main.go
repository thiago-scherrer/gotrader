package main

import (
	"log"

	"github.com/thiago-scherrer/gotrader/internal/api"
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
	reader.InitFlag()

	log.Println(
		display.HelloMsg(reader.Asset()),
	)
	api.MatrixSend(
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
			central.Price(),
		),
	)

	api.MatrixSend(
		display.OrderPrice(
			reader.Asset(),
			central.Price(),
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

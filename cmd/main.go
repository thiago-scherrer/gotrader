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

	if central.CreateOrder(trd) == false {
		trd = "Error"
	}

	if trd == "Buy" {
		logic.ClosePositionProfitBuy()
	} else if trd == "Sell" {
		logic.ClosePositionProfitSell()
	} else {
		display.OrderCancelMsg()
	}
}

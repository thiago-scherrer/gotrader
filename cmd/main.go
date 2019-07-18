package main

import (
	"log"

	"github.com/thiago-scherrer/gotrader/internal/api"
	"github.com/thiago-scherrer/gotrader/internal/central"
	dp "github.com/thiago-scherrer/gotrader/internal/display"
	"github.com/thiago-scherrer/gotrader/internal/logic"
	rd "github.com/thiago-scherrer/gotrader/internal/reader"
)

func main() {
	for {
		daemonize()
	}
}

func daemonize() {
	rd.InitFlag()

	log.Println(
		dp.HelloMsg(rd.Asset()),
	)
	api.MatrixSend(
		dp.HelloMsg(rd.Asset()),
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
		dp.OrderCancelMsg()
	}
}

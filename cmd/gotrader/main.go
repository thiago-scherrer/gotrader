package main

import (
	"log"

	"github.com/thiago-scherrer/gotrader/internal/api"
	ctr "github.com/thiago-scherrer/gotrader/internal/central"
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

	if ctr.CreateOrder(trd) == false {
		dp.OrderCancelMsg()
		return
	}

	if trd == "Buy" {
		ctr.ClosePositionProfitBuy()
	} else {
		ctr.ClosePositionProfitSell()
	}
	ctr.GetProfit()
}

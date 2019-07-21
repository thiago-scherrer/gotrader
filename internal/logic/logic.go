package logic

import (
	"github.com/thiago-scherrer/gotrader/internal/convert"
	"github.com/thiago-scherrer/gotrader/internal/reader"
)

// Path from api to view the orderbook
const orb string = "/api/v1/orderBook/L2?"

// CandleRunner verify the api and start the logic system
func CandleRunner() string {
	return "statusLogic"
}

func returnDepth() string {
	return "Buy/Sell/Draw"
}

// ClosePositionProfitBuy the Buy position
func ClosePositionProfitBuy() {

}

// ClosePositionProfitSell close the Sell position
func ClosePositionProfitSell() {
}

// GetHand change the hand according to a strategy
func GetHand() string {
	return convert.IntToString(
		reader.Hand(),
	)
}

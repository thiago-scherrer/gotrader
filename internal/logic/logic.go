package logic

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

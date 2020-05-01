package display

import (
	"github.com/thiago-scherrer/gotrader/internal/convert"
)

// HelloMsg will be see when the bot start a new trade
func HelloMsg(a string) string {
	return " " + a + " - Starting a new round"
}

// UsageMsg display a basic msg when not found the config file
func UsageMsg() string {
	return "Config not found! Usage: gotrader config some_config.yml"
}

// SetleverageMsg display set the l
func SetleverageMsg(a, l string) string {
	return " " + a + " - Setting leverage: " + l
}

// OrderCreatedMsg display the type order to the user
func OrderCreatedMsg(a, t string) string {
	return " " + a + " - A new order type: " + t + " as been created! "
}

// OrderCancelMsg display cancell msg
func OrderCancelMsg() string {
	return "Canceling trade, order not executed"
}

// OrderDoneMsg display a msg when order fulfilled
func OrderDoneMsg(a string) string {
	return " " + a + " - Order fulfilled!"
}

// OrdertriggerMsg display a msg when order trigged
func OrdertriggerMsg(a string) string {
	return " " + a + " - Profit target trigged"
}

// OrderWaintMsg display a msg when will waint
func OrderWaintMsg(a string) string {
	return " " + a + " - Waiting to get order fulfilled..."
}

//StopLossMsg show stop loss msg
func StopLossMsg(a string) string {
	return " " + a + " - Stop Loss trigged!"
}

// ProfitMsg display msg  when the trade get profit
func ProfitMsg(a string) string {
	return " " + a + " - Round done."
}

// OrderPrice display the current price
func OrderPrice(asset string, price float64) string {
	return " " + asset + " - Price now: " + convert.FloatToString(price)
}

// DrawMode display the draw msg
func DrawMode() string {
	return "The number of buy and sell orders are equal, leaving."
}

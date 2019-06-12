package display

// HelloMsg will be see when the bot start a new trade
func HelloMsg(asset string) string {
	return " " + asset + " - Starting a new round"
}

// UsageMsg display a basic msg when not found the config file
func UsageMsg() string {
	return "Config not found! Usage: config config.yml"
}

// SetlavarageMsg display set the leverage
func SetlavarageMsg(asset, leverage string) string {
	return " " + asset + " - Setting leverage: " + leverage
}

// OrderCreatedMsg display the type order to the user
func OrderCreatedMsg(asset, typeOrder string) string {
	return " " + asset + " - A new order type: " + typeOrder + " as been created! "
}

// OrderDoneMsg display a msg when order fulfilled
func OrderDoneMsg(asset string) string {
	return " " + asset + " - Order fulfilled!"
}

// OrdertriggerMsg display a msg when order trigged
func OrdertriggerMsg(asset string) string {
	return " " + asset + " - Profit target trigged"
}

// OrderWaintMsg display a msg when will waint
func OrderWaintMsg(asset string) string {
	return " " + asset + " - Waiting to get order fulfilled..."
}

// ProfitMsg display msg  when the trade get profit
func ProfitMsg(asset string) string {
	return " " + asset + " - Profit done!"
}

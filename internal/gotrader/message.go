package main

func helloMsg() string {
	return "Starting a new round => GoTrader Bot"
}

func toLowerMsg() string {
	return "Hand to lower, setting 1 "
}

func usageMsg() string {
	return "Usage: config config.yml"
}

func setlavarageMsg() string {
	return "Setting leverage: " + leverage()
}

func orderCreatedMsg(typeOrder string) string {
	return "A new order type: " + typeOrder + " as been created! "
}

func orderDoneMsg() string {
	return "Order fulfilled!"
}

func ordertriggerMsg() string {
	return "Profit target trigged"
}

func orderWaintMsg() string {
	return "Waiting to get order fulfilled..."
}

func profitMsg() string {
	return "Profit done!"
}

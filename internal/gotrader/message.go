package main

func helloMsg() string {
	return " " + asset() + " - Starting a new round"
}

func usageMsg() string {
	return "Usage: config config.yml"
}

func setlavarageMsg() string {
	return " " + asset() + " - Setting leverage: " + leverage()
}

func orderCreatedMsg(typeOrder string) string {
	return " " + asset() + " - A new order type: " + typeOrder + " as been created! "
}

func orderDoneMsg() string {
	return " " + asset() + " - Order fulfilled!"
}

func ordertriggerMsg() string {
	return " " + asset() + " - Profit target trigged"
}

func orderWaintMsg() string {
	return " " + asset() + " - Waiting to get order fulfilled..."
}

func profitMsg() string {
	return " " + asset() + " - Profit done!"
}

package main

func main() {
	/*

		var logicResult string
		var strikeBuy int
		var strikeSell int
		var strike int
		confFile := "../../configs/config.yml"
			requiredConfig(confFile)
			asset := configReader("asset", confFile)
			candle := configReader("candle", confFile)
			endpoint := configReader("api", confFile)
			lConfig := configReader("logic", confFile)
			secretQuery := configReader("secret", confFile)
			speed := StringToInt(configReader("speed", confFile))
			strike = StringToIntBit(configReader("threshold", confFile))
			userIDquery := configReader("userid", confFile)
			expire := IntToString((timeExpired()))
			hexResult := hexCreator(secretQuery, "GET", "/api/v1/user/wallet", expire)

			// get $$ from account
			money := parserAmount(clientGet(hexResult,
				endpoint, "/api/v1/user/wallet", expire, userIDquery),
			)

			hand := handRoll(money,
				StringToInt(configReader("hand", confFile)),
			)

			for count := 0; count <= strike; count++ {

				switch lConfig {
				case "volume":
					logicResult = volume(userIDquery, secretQuery, endpoint, asset, candle, hand, speed)
				default:
					panic("logic nog found!")
				}
				if logicResult == "Buy" {
					strikeBuy++
				} else if logicResult == "Sell" {
					strikeSell++
				} else {
					count--
				}
				fmt.Println(logicResult)
			}
	*/
}

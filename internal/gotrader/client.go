package main

import (
	"io/ioutil"
	"net/http"
)

func clientRobot(requestType, config, path string) []byte {
	client := &http.Client{}

	endpoint := configReader("api", config)
	secretQuery := configReader("secret", config)
	userIDquery := configReader("userid", config)
	expire := IntToString((timeExpired()))
	hexResult := hexCreator(secretQuery, requestType, path, expire)
	url := endpoint + path

	request, err := http.NewRequest(requestType, url, nil)
	request.Header.Add("api-signature", hexResult)
	request.Header.Add("api-expires", expire)
	request.Header.Add("api-key", userIDquery)
	request.Header.Add("Content-Type", "text/plain; charset=utf-8")
	request.Header.Add("User-Agent", "gotrader-r0b0tnull")
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func getQuote() {

}

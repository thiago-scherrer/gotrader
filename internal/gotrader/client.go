package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func clientRobot(requestType, config, path string, data []byte) []byte {
	client := &http.Client{}
	endpoint := configReader("api", config)
	secretQuery := configReader("secret", config)
	userIDquery := configReader("userid", config)
	expire := IntToString((timeExpired()))
	hexResult := hexCreator(secretQuery, requestType, path, expire)
	url := endpoint + path
	
	request, err := http.NewRequest(requestType, url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("api-signature", hexResult)
	request.Header.Add("api-expires", expire)
	request.Header.Add("api-key", userIDquery)
	request.Header.Add("Content-Type", "text/plain; charset=utf-8")
	request.Header.Add("User-Agent", "gotrader-r0b0tnull")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	if response.StatusCode == 401 {
		fmt.Println("quiting, API response are: ")
		panic(response.StatusCode)
	} else if response.StatusCode == 404 {
		fmt.Println("quiting, API response are: ")
		panic(response.StatusCode)
	}

	return body
}

func getQuote() {

}

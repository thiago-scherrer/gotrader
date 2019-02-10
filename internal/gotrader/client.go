package main

import (
	"io/ioutil"
	"net/http"
)

func clientGet(hex, endpoint, path, expired, userid string) string {
	url := endpoint + path

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("api-signature", hex)
	request.Header.Add("api-expires", expired)
	request.Header.Add("api-key", userid)
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

	return BytesToString(body)
}

func clientPost(hex, endpoint, path, expired, userid string) string {
	url := endpoint + path

	client := &http.Client{}

	request, err := http.NewRequest("POST", url, nil)
	request.Header.Add("api-signature", hex)
	request.Header.Add("api-expires", expired)
	request.Header.Add("api-key", userid)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
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

	return BytesToString(body)
}

func clientRobot(requestType, config, path string) []byte {
	client := &http.Client{}

	endpoint := configReader("api", config)
	secretQuery := configReader("secret", config)
	userIDquery := configReader("userid", config)
	expire := IntToString((timeExpired()))
	hexResult := hexCreator(secretQuery, requestType, path, expire)
	url := endpoint + path

	request, err := http.NewRequest("GET", url, nil)
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

	//return BytesToString(body)
	return body
}

func getQuote() {

}

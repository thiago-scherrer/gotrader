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

func getQuote() {
	
}

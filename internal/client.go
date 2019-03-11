package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func clientRobot(requestType, path string, data []byte) []byte {
	client := &http.Client{}
	endpoint := endpoint()
	secretQuery := secret()
	userIDquery := userid()
	expire := IntToString((timeExpired()))
	hexResult := hexCreator(secretQuery, requestType, path, expire, BytesToString(data))

	url := endpoint + path

	request, err := http.NewRequest(requestType, url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error create a request on bitmex, got: ", err)
	}

	request.Header.Set("api-signature", hexResult)
	request.Header.Set("api-expires", expire)
	request.Header.Set("api-key", userIDquery)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "gotrader-r0b0tnull")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error to send the request to the API bitmex, got: ", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	if verboseMode() {
		fmt.Println("Bitmex API Status code are: ", response.StatusCode)
	}

	return body
}

func telegramSend(msg string) int {
	if telegramUse() == false {
		return 200
	}

	client := &http.Client{}
	telegramurl := telegramurl()
	telegramChannel := telegramChannel()
	token := telegramKey()
	data := StringToBytes("chat_id=" + telegramChannel + "&text=" + msg)
	url := telegramurl + "/bot" + token + "/sendMessage"

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error create a request on telegram, got: ", err)
	}

	request.Header.Set("User-Agent", "gotrader-r0b0tnull")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil && verboseMode() {
		fmt.Println("Error to get body from Telegram API, got", err)
	}

	if verboseMode() {
		fmt.Println("Telegram API Status code are: ", response.StatusCode)
	}

	return response.StatusCode
}

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
		fmt.Println(err)
	}

	request.Header.Set("api-signature", hexResult)
	request.Header.Set("api-expires", expire)
	request.Header.Set("api-key", userIDquery)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "gotrader-r0b0tnull")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 401 {
		fmt.Println("quiting, API response are: ", response.StatusCode)
	} else if response.StatusCode == 404 {
		fmt.Println("quiting, API response are: ", response.StatusCode)
	}

	return body
}

package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/thiago-scherrer/gotrader/internal/convert"
	"github.com/thiago-scherrer/gotrader/internal/reader"
)

// hexCreator encode a string to a sha256, this is needed in bitme API
func hexCreator(Secret, requestTipe, path, expired, data string) string {
	concat := requestTipe + path + expired + data
	h := hmac.New(sha256.New, []byte(Secret))
	h.Write([]byte(concat))
	hexResult := hex.EncodeToString(h.Sum(nil))
	return hexResult
}

// timeExpired create a time to expire a session API.
func timeExpired() int64 {
	timeExpired := timeStamp() + 60
	return timeExpired
}

// timeStamp create a timestamp to be used in a session.
func timeStamp() int64 {
	now := time.Now()
	timestamp := now.Unix()
	return timestamp
}

// ClientRobot are the curl from the bot
func ClientRobot(requestType, path string, data []byte) []byte {
	for {
		cl := &http.Client{}
		ep := reader.Endpoint()
		sq := reader.Secret()
		uid := reader.Userid()
		exp := convert.IntToString((timeExpired()))
		hex := hexCreator(sq, requestType, path, exp, convert.BytesToString(data))
		url := ep + path

		req, err := http.NewRequest(requestType, url, bytes.NewBuffer(data))
		if err != nil {
			log.Println("Error create a request on bitmex, got: ", err)
		}

		req.Header.Set("api-signature", hex)
		req.Header.Set("api-expires", exp)
		req.Header.Set("api-key", uid)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("User-Agent", "gotrader-r0b0tnull")

		rsp, err := cl.Do(req)
		if err != nil {
			log.Println("Error to send the request to the API bitmex, got: ", err)
		}

		body, _ := ioutil.ReadAll(rsp.Body)

		if rsp.StatusCode != 200 {
			log.Println("Bitmex API Status code are: ", rsp.StatusCode)
			log.Println("Body: ", convert.BytesToString(body))

			time.Sleep(time.Duration(60) * time.Second)
		} else {
			return body
		}
	}
}

// MatrixSend send a msg to the user on settings
func MatrixSend(msg string) int {
	if reader.MatrixUse() == false {
		return 200
	}

	cl := &http.Client{}
	turl := reader.Matrixurl()
	tch := reader.MatrixChannel()
	tkn := reader.MatrixKey()
	dt := convert.StringToBytes(`{"msgtype":"m.text", "body": "` + msg + `"}`)
	url := turl + "_matrix/client/r0/rooms/" + tch + "/send/m.room.message?access_token=" + tkn

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dt))
	if err != nil {
		log.Println("Error create a request on matrix, got: ", err)
	}

	req.Header.Set("User-Agent", "gotrader-r0b0tnull")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	rsp, err := cl.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer rsp.Body.Close()
	_, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Println("Error to get body from Matrix API, got", err)
	}
	return rsp.StatusCode
}

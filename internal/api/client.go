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

	hexResult := hex.EncodeToString(
		h.Sum(nil),
	)

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

// ClientRobot make all request to the bitmex API
func ClientRobot(requestType, path string, data []byte) ([]byte, int) {
	cl := &http.Client{}
	ep := reader.Endpoint()
	sq := reader.Secret()
	uid := reader.Userid()
	exp := convert.IntToString(
		(timeExpired()),
	)
	hex := hexCreator(
		sq,
		requestType,
		path,
		exp,
		convert.BytesToString(data),
	)
	url := ep + path

	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(data))
	if err != nil {
		log.Println("Error create a request, got: ", err)
	}

	req.Header.Set("api-signature", hex)
	req.Header.Set("api-expires", exp)
	req.Header.Set("api-key", uid)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "gotrader")
	rsp, err := cl.Do(req)
	if err != nil {
		log.Println("Error to send the request to the API bitmex, got: ", err)
		rsp.Body.Close()
	}
	body, _ := ioutil.ReadAll(rsp.Body)

	return body, rsp.StatusCode
}

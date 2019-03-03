package main

import "testing"

func TestUsageMsg(t *testing.T) {
	if usageMsg() != "Usage: config config.yml" {
		t.Error("error to get usage msg, got:", usageMsg())
	}
}

func TestTelegram(t *testing.T) {
	getResult := telegramSend(helloMsg())

	if getResult != 200 {
		t.Error("Telegram not working, got: ", getResult)
	}
}

package main

import (
	"testing"
)

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

func Test_setlavarageMsg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"test: ", " " + asset() + " - Setting leverage: 0.1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setlavarageMsg(); got != tt.want {
				t.Errorf("setlavarageMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderCreatedMsg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", " " + asset() + " - A new order type: Sell as been created! "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderCreatedMsg("Sell"); got != tt.want {
				t.Errorf("orderCreatedMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderDoneMsg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", " " + asset() + " - Order fulfilled!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderDoneMsg(); got != tt.want {
				t.Errorf("orderDoneMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ordertriggerMsg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", " " + asset() + " - Profit target trigged"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ordertriggerMsg(); got != tt.want {
				t.Errorf("ordertriggerMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderWaintMsg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", " " + asset() + " - Waiting to get order fulfilled..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderWaintMsg(); got != tt.want {
				t.Errorf("orderWaintMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_profitMsg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", " " + asset() + " - Profit done!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := profitMsg(); got != tt.want {
				t.Errorf("profitMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

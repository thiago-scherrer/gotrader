package main

import (
	"reflect"
	"testing"
	"time"
)

func TestFlag(t *testing.T) {
	getResult := initFlag()
	if len(getResult) <= 1 {
		t.Error("init flag not working, got: ", getResult)
	}
}

func TestReader(t *testing.T) {
	getResult := configReader()

	if getResult.Asset != "XBTUSD" {
		t.Error("error to read config file, got:", getResult)
	}
}

func TestHmac(t *testing.T) {
	initFlag()

	expired := "1518064236"
	path := "/api/v1/instrument"
	requestTipe := "GET"
	secretQuery := "chNOOS4KvNXR_Xq4k4c9qsfoKWvnDecLATCRlcBwyKDYnWgO"
	hexExpected := "9c37199dd75f47b63774ddbb5e2851998848d5ec62b9a2bbc380a48f620b305e"
	hexResult := hexCreator(secretQuery, requestTipe, path, expired, "data")

	if hexExpected != hexResult {
		t.Error("GET hex not match, got: ", hexResult, "need: ", hexExpected)
	}
}

func TestExpiresTime(t *testing.T) {
	initFlag()
	var timeStampResult TimeStamp

	now := time.Now()
	timestamp := now.Unix()
	timeStampResult.timeResult = timeStamp()
	diff := timestamp - timeStampResult.timeResult
	timeStampResult.timeExpired = timeExpired()
	timeStampResultExpected := timestamp + 60

	if diff != 0 {
		t.Error("time stamp function noting working, result are: ", diff)
	}
	if timeStampResult.timeExpired != timeStampResultExpected {
		t.Error("expired time not working, resulte are: ", timeStampResult.timeExpired, "expected: ", timeStampResultExpected)
	}
}

func TestParserAmount(t *testing.T) {
	mock := `{ "Amount": 10 }`
	getResult := parserAmount(StringToBytes(mock))

	if getResult != 10 {
		t.Error("json parser not working, got:", getResult)
	}
}

func TestLastPriceJson(t *testing.T) {
	mock := `[{ "LastPrice": 10.1 }]`
	getResult := lastPrice(StringToBytes(mock))

	if getResult != 10.1 {
		t.Error("LastPrice json parser not working, got:", getResult)
	}
}

func Test_initFlag(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "../../configs/config-test.yml"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initFlag(); got != tt.want {
				t.Errorf("initFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_configReader(t *testing.T) {
	tests := []struct {
		name string
		want *Conf
	}{
		{"Test", configReader()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configReader(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asset(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "XBTUSD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asset(); got != tt.want {
				t.Errorf("asset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_candle(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Test", 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candle(); got != tt.want {
				t.Errorf("candle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_depth(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"Test", 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := depth(); got != tt.want {
				t.Errorf("depth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_endpoint(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "https://testnet.bitmex.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := endpoint(); got != tt.want {
				t.Errorf("endpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Test", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hand(); got != tt.want {
				t.Errorf("hand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leverage(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "0.1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leverage(); got != tt.want {
				t.Errorf("leverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_profit(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		{"Test", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := profit(); got != tt.want {
				t.Errorf("profit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secret(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "chNOOS4KvNXR_Xq4k4c9qsfoKWvnDecLATCRlcBwyKDYnWgO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secret(); got != tt.want {
				t.Errorf("secret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threshold(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Test", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threshold(); got != tt.want {
				t.Errorf("threshold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userid(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "LAqUlngMIQkIUjXMUreyu3qn"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := userid(); got != tt.want {
				t.Errorf("userid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_telegramUse(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := telegramUse(); got != tt.want {
				t.Errorf("telegramUse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_telegramKey(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "xxxxx"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := telegramKey(); got != tt.want {
				t.Errorf("telegramKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_telegramurl(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "https://api.telegram.org"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := telegramurl(); got != tt.want {
				t.Errorf("telegramurl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_telegramChannel(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "@"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := telegramChannel(); got != tt.want {
				t.Errorf("telegramChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_verboseMode(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verboseMode(); got != tt.want {
				t.Errorf("verboseMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_speed(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Test", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := speed(); got != tt.want {
				t.Errorf("speed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handRoll(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Test", 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handRoll(42, 42); got != tt.want {
				t.Errorf("handRoll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hexCreator(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "4210b09a7ec51b6399c8b32284925bce0c28156b4800d97f4cf5815ab059fd4b"},
	}
	secret := "42"
	requestTipe := "42"
	path := "42"
	expired := "never"
	data := "42"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hexCreator(secret, requestTipe, path,
				expired, data); got != tt.want {
				t.Errorf("hexCreator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parserAmount(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Test", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parserAmount(StringToBytes(`{"amount":1}`)); got != tt.want {
				t.Errorf("parserAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lastPrice(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		{"Test", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastPrice(StringToBytes(`[{"lastPrice":1}]`)); got != tt.want {
				t.Errorf("lastPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_opening(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Test", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := opening(StringToBytes(`[{"isOpen":true}]`)); got != tt.want {
				t.Errorf("opening() = %v, want %v", got, tt.want)
			}
		})
	}
}

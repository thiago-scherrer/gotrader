package central

import (
	"testing"

	"github.com/thiago-scherrer/gotrader/internal/convert"
)

func TestParserAmount(t *testing.T) {
	mock := `{ "Amount": 10 }`
	getResult := parserAmount(convert.StringToBytes(mock))

	if getResult != 10 {
		t.Error("json parser not working, got:", getResult)
	}
}

func TestLastPriceJson(t *testing.T) {
	mock := `[{ "LastPrice": 10.1 }]`
	getResult := lastPrice(convert.StringToBytes(mock))

	if getResult != 10.1 {
		t.Error("LastPrice json parser not working, got:", getResult)
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
			if got := lastPrice(convert.StringToBytes(`[{"lastPrice":1}]`)); got != tt.want {
				t.Errorf("lastPrice() = %v, want %v", got, tt.want)
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
			if got := parserAmount(convert.StringToBytes(`{"amount":1}`)); got != tt.want {
				t.Errorf("parserAmount() = %v, want %v", got, tt.want)
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
			if got := opening(convert.StringToBytes(`[{"isOpen":true}]`)); got != tt.want {
				t.Errorf("opening() = %v, want %v", got, tt.want)
			}
		})
	}
}

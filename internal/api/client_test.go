package api

import (
	"testing"

	"github.com/thiago-scherrer/gotrader/internal/reader"
)

func TestHmac(t *testing.T) {
	reader.InitFlag()

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

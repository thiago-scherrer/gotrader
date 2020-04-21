package reader

import (
	"reflect"
	"testing"
)

func TestFlag(t *testing.T) {
	getResult := ConfigPath()
	if len(getResult) <= 1 {
		t.Error("init flag not working, got: ", getResult)
	}
}

func TestReader(t *testing.T) {
	Boot()

	asset := Asset()

	if asset != "XBTUSD" {
		t.Error("error to read config file, got:", asset)
	}
}

func Test_configReader(t *testing.T) {
	Boot()
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
	Boot()
	tests := []struct {
		name string
		want string
	}{
		{"Test", "XBTUSD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Asset(); got != tt.want {
				t.Errorf("Asset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Candle(t *testing.T) {
	Boot()
	tests := []struct {
		name string
		want int64
	}{
		{"Test", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Candle(); got != tt.want {
				t.Errorf("Candle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_endpoint(t *testing.T) {
	Boot()
	tests := []struct {
		name string
		want string
	}{
		{"Test", "https://testnet.bitmex.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Endpoint(); got != tt.want {
				t.Errorf("Endpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leverage(t *testing.T) {
	Boot()
	tests := []struct {
		name string
		want string
	}{
		{"Test", "0.1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Leverage(); got != tt.want {
				t.Errorf("leverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secret(t *testing.T) {
	Boot()
	tests := []struct {
		name string
		want string
	}{
		{"Test", "chNOOS4KvNXR_Xq4k4c9qsfoKWvnDecLATCRlcBwyKDYnWgO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Secret(); got != tt.want {
				t.Errorf("secret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Threshold(t *testing.T) {
	Boot()
	tests := []struct {
		name string
		want int64
	}{
		{"Test", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Threshold(); got != tt.want {
				t.Errorf("Threshold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userid(t *testing.T) {
	Boot()
	tests := []struct {
		name string
		want string
	}{
		{"Test", "LAqUlngMIQkIUjXMUreyu3qn"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Userid(); got != tt.want {
				t.Errorf("userid() = %v, want %v", got, tt.want)
			}
		})
	}
}

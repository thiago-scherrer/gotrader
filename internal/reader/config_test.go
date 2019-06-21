package reader

import (
	"reflect"
	"testing"
)

func TestFlag(t *testing.T) {
	getResult := InitFlag()
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

func Test_InitFlag(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "/opt/config-test.yml"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitFlag(); got != tt.want {
				t.Errorf("InitFlag() = %v, want %v", got, tt.want)
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
			if got := Asset(); got != tt.want {
				t.Errorf("Asset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Candle(t *testing.T) {
	tests := []struct {
		name string
		want int
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

func Test_profit(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		{"Test", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Profit(); got != tt.want {
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
			if got := Secret(); got != tt.want {
				t.Errorf("secret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Threshold(t *testing.T) {
	tests := []struct {
		name string
		want int
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

func Test_matrixUse(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatrixUse(); got != tt.want {
				t.Errorf("matrixUse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matrixKey(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "!xxxxxxx:matrix.org"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatrixKey(); got != tt.want {
				t.Errorf("matrixKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matrixurl(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "https://matrix.org"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Matrixurl(); got != tt.want {
				t.Errorf("matrixurl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matrixChannel(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "@"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatrixChannel(); got != tt.want {
				t.Errorf("matrixChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}

package display

import (
	"testing"
)

func TestUsageMsg(t *testing.T) {
	if UsageMsg() != "Config not found! Usage: gotrader config some_config.yml" {
		t.Error("error to get usage msg, got:", UsageMsg())
	}
}

func Test_setleverageMsg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"test: ", " " + "BTC" + " - Setting leverage: 0.1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetleverageMsg("BTC", "0.1"); got != tt.want {
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
		{"Test", " " + "BTC" + " - A new order type: Sell as been created! "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderCreatedMsg("BTC", "Sell"); got != tt.want {
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
		{"Test", " " + "BTC" + " - Order fulfilled!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderDoneMsg("BTC"); got != tt.want {
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
		{"Test", " " + "BTC" + " - Profit target trigged"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrdertriggerMsg("BTC"); got != tt.want {
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
		{"Test", " " + "BTC" + " - Waiting to get order fulfilled..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderWaintMsg("BTC"); got != tt.want {
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
		{"Test", " " + "BTC" + " - Profit done!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProfitMsg("BTC"); got != tt.want {
				t.Errorf("profitMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
